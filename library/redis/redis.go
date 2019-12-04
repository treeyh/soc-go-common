package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/treeyh/soc-go-common/core/errors"
	"github.com/treeyh/soc-go-common/core/utils/times"
	"math/rand"
	"reflect"
)

const (
	_LockDistributedLua = "local v;" +
		"v = redis.call('setnx',KEYS[1],ARGV[1]);" +
		"if tonumber(v) == 1 then\n" +
		"    redis.call('expire',KEYS[1],ARGV[2])\n" +
		"end\n" +
		"return v"

	_UnLockDistributedLua = "if redis.call('get',KEYS[1]) == ARGV[1]\n" +
		"then\n" +
		"    return redis.call('del',KEYS[1])\n" +
		"else\n" +
		"    return 0\n" +
		"end"

	_DistributedTimeOut = 4
	_DistributedSuccess = 1

	_MasterConfigName = "Master"
)

var (
	_LockDistributedLuaScript   = redis.NewScript(1, _LockDistributedLua)
	_UnLockDistributedLuaScript = redis.NewScript(1, _UnLockDistributedLua)

	_redisCaches = make(map[string]*RedisProxy)
)

type RedisProxy struct {
	name string
	conn redis.Conn
}

func (rp *RedisProxy) ZAdd(key string, score float64, value string) errors.AppError {
	conn := rp.Connect()
	defer rp.Close(conn)

	_, err := conn.Do("zadd", key, score, value)
	return errors.NewAppErrorExistError(errors.RedisOperationFail, err)
}

func (rp *RedisProxy) SetEx(key string, value string, ex int64) errors.AppError {
	conn := rp.Connect()
	defer rp.Close(conn)

	_, err := conn.Do("setex", key, ex, value)
	return errors.NewAppErrorExistError(errors.RedisOperationFail, err)
}

func (rp *RedisProxy) Set(key string, value string) errors.AppError {
	conn := rp.Connect()
	defer rp.Close(conn)

	_, err := conn.Do("set", key, value)
	return errors.NewAppErrorExistError(errors.RedisOperationFail, err)
}

func (rp *RedisProxy) MSet(fieldValue map[string]string) errors.AppError {
	conn := rp.Connect()
	defer rp.Close(conn)

	args := []interface{}{}
	args = append(args)
	for k, v := range fieldValue {
		args = append(append(args, k), v)
	}
	_, err := conn.Do("mset", args...)
	return errors.NewAppErrorExistError(errors.RedisOperationFail, err)
}

func (rp *RedisProxy) Get(key string) (string, errors.AppError) {
	conn := rp.Connect()
	defer rp.Close(conn)

	rs, err := conn.Do("get", key)
	if err != nil {
		return "", errors.NewAppErrorExistError(errors.RedisOperationFail, err)
	}
	if rs == nil {
		return "", nil
	}
	return string(rs.([]byte)), nil
}

func (rp *RedisProxy) MGet(keys ...interface{}) ([]string, errors.AppError) {
	conn := rp.Connect()
	defer rp.Close(conn)

	rs, err := conn.Do("mget", keys...)
	if err != nil {
		return nil, errors.NewAppErrorExistError(errors.RedisOperationFail, err)
	}

	if rs == nil {
		return nil, nil
	}
	//re, err := redis.Values(rs, err)
	//fmt.Println(re)
	list := rs.([]interface{})
	resultList := make([]string, 0)
	for _, v := range list {
		if bytes, ok := v.([]byte); ok {
			resultList = append(resultList, string(bytes))
		}
	}
	return resultList, nil
}

func (rp *RedisProxy) Del(keys ...interface{}) (int64, errors.AppError) {
	conn := rp.Connect()
	defer rp.Close(conn)

	rs, err := conn.Do("Del", keys...)
	if err != nil {
		return 0, errors.NewAppErrorExistError(errors.RedisOperationFail, err)
	}
	if rs == nil {
		return 0, errors.NewAppErrorExistError(errors.RedisOperationFail, err)
	}
	return rs.(int64), nil
}

func (rp *RedisProxy) Incrby(key string, v int64) (int64, errors.AppError) {
	conn := rp.Connect()
	defer rp.Close(conn)

	rs, err := conn.Do("INCRBY", key, v)
	if err != nil {
		return 0, errors.NewAppErrorExistError(errors.RedisOperationFail, err)
	}
	if rs == nil {
		return 0, errors.NewAppErrorExistError(errors.RedisOperationFail, err)
	}
	return rs.(int64), errors.NewAppErrorExistError(errors.RedisOperationFail, err)
}

func (rp *RedisProxy) Exist(key string) (bool, errors.AppError) {
	conn := rp.Connect()
	defer rp.Close(conn)

	rs, err := conn.Do("EXISTS", key)
	if err != nil {
		return false, errors.NewAppErrorExistError(errors.RedisOperationFail, err)
	}
	if rs == nil {
		return false, nil
	}
	return rs.(int64) == 1, nil
}

func (rp *RedisProxy) Expire(key string, expire int64) (bool, errors.AppError) {
	conn := rp.Connect()
	defer rp.Close(conn)

	rs, err := conn.Do("EXPIRE", key, expire)
	if err != nil {
		return false, errors.NewAppErrorExistError(errors.RedisOperationFail, err)
	}
	if rs == nil {
		return false, nil
	}
	return rs.(int64) == 1, nil
}

func (rp *RedisProxy) TryGetDistributedLock(key string, v string) (bool, errors.AppError) {
	conn := rp.Connect()
	defer rp.Close(conn)

	end := times.GetNowMillisecond() + _DistributedTimeOut*1000
	for times.GetNowMillisecond() <= end {
		rs, err := _LockDistributedLuaScript.Do(conn, key, v, _DistributedTimeOut)
		if err != nil {
			return false, errors.NewAppErrorExistError(errors.RedisOperationFail, err)
		}
		if rs.(int64) == _DistributedSuccess {
			return true, nil
		}
		times.SleepMillisecond(80 + int64(rand.Int31n(30)))
	}

	return false, nil
}

func (rp *RedisProxy) ReleaseDistributedLock(key string, v string) (bool, errors.AppError) {
	conn := rp.Connect()
	defer rp.Close(conn)

	rs, err := _UnLockDistributedLuaScript.Do(conn, key, v)
	if err != nil {
		return false, errors.NewAppErrorExistError(errors.RedisOperationFail, err)
	}
	if rs.(int64) == _DistributedSuccess {
		return true, nil
	}

	return false, nil
}

func (rp *RedisProxy) HGet(key, field string) (string, errors.AppError) {
	conn := rp.Connect()
	defer rp.Close(conn)

	rs, err := conn.Do("HGET", key, field)
	if err != nil {
		return "", errors.NewAppErrorExistError(errors.RedisOperationFail, err)
	}
	if rs == nil {
		return "", nil
	}
	return string(rs.([]byte)), nil
}

func (rp *RedisProxy) HSet(key, field, value string) errors.AppError {
	conn := rp.Connect()
	defer rp.Close(conn)

	_, err := conn.Do("HSET", key, field, value)
	return errors.NewAppErrorExistError(errors.RedisOperationFail, err)
}

func (rp *RedisProxy) HDel(key string, fields ...interface{}) (int64, errors.AppError) {
	conn := rp.Connect()
	defer rp.Close(conn)

	args := []interface{}{}
	args = append(append(args, key), fields...)
	rs, err := conn.Do("HDEL", args)
	if err != nil {
		return 0, errors.NewAppErrorExistError(errors.RedisOperationFail, err)
	}
	if rs == nil {
		return 0, nil
	}
	return rs.(int64), nil
}

func (rp *RedisProxy) HExists(key, field string) (bool, errors.AppError) {
	conn := rp.Connect()
	defer rp.Close(conn)

	rs, err := conn.Do("HEXISTS", key, field)
	if err != nil {
		return false, errors.NewAppErrorExistError(errors.RedisOperationFail, err)
	}
	if rs == nil {
		return false, nil
	}
	return rs.(int64) == 1, nil
}

func (rp *RedisProxy) HMGet(key string, fields ...interface{}) (map[string]*string, errors.AppError) {
	result := map[string]*string{}
	conn := rp.Connect()
	defer rp.Close(conn)

	fields = append([]interface{}{key}, fields...)
	rs, err := conn.Do("HMGET", fields...)
	if err != nil {
		return result, errors.NewAppErrorExistError(errors.RedisOperationFail, err)
	}
	if rs == nil {
		return result, nil
	}
	keys := make([]string, len(fields)-1)
	for i, k := range fields {
		if i == 0 {
			continue
		}
		keys[i-1] = k.(string)
	}

	if datas, ok := rs.([]interface{}); ok {
		for i, data := range datas {
			if data == nil {
				result[keys[i]] = nil
			} else {
				dataStr := string(data.([]byte))
				result[keys[i]] = &dataStr
			}
		}
	}
	return result, nil
}

func (rp *RedisProxy) HMSet(key string, fieldValue map[string]string) errors.AppError {
	conn := rp.Connect()
	defer rp.Close(conn)

	args := []interface{}{}
	args = append(args, key)
	for k, v := range fieldValue {
		args = append(append(args, k), v)
	}
	_, err := conn.Do("HMSET", args...)
	return errors.NewAppErrorExistError(errors.RedisOperationFail, err)
}

func (rp *RedisProxy) Connect() redis.Conn {
	return GetRedisConn(rp.name)
}

func (rp *RedisProxy) Close(conn redis.Conn) errors.AppError {
	if conn != nil && conn.Err() == nil {
		conn.Close()
		err := conn.Close()
		return errors.NewAppErrorExistError(errors.RedisOperationFail, err)
	}
	return nil
}

func (rp RedisProxy) IsEmpty() bool {
	return reflect.DeepEqual(rp, RedisProxy{})
}

// GetProxy get default redis oper proxy
func GetProxy() *RedisProxy {
	return GetProxyByName(_MasterConfigName)
}

// GetProxyByName get redis oper proxy
func GetProxyByName(name string) *RedisProxy {
	if v, ok := _redisCaches[name]; ok {
		return v
	}
	redisProxy := &RedisProxy{
		name: name,
	}
	_redisCaches[name] = redisProxy
	return redisProxy
}
