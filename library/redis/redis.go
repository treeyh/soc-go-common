package redis

import (
	"github.com/mediocregopher/radix/v3"
	"github.com/treeyh/soc-go-common/core/errors"
	"github.com/treeyh/soc-go-common/core/utils/times"
	"math/rand"
	"reflect"
	"strconv"
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

	_MasterConfigName = "master"
)

var (
	_LockDistributedLuaScript   = radix.NewEvalScript(1, _LockDistributedLua)
	_UnLockDistributedLuaScript = radix.NewEvalScript(1, _UnLockDistributedLua)

	_redisCaches = make(map[string]*RedisProxy)
)

type RedisProxy struct {
	name string
}

func (rp *RedisProxy) ZAdd(key string, score float64, value string) errors.AppError {
	pool := rp.Pool()
	err := pool.Do(radix.Cmd(nil,"zadd", key, strconv.FormatFloat(score, 'f', 10, 64), value))

	return errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
}

func (rp *RedisProxy) SetEx(key string, value string, ex int) errors.AppError {
	pool := rp.Pool()

	err := pool.Do(radix.Cmd(nil,"setex", key, strconv.Itoa(ex), value))

	return errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
}

func (rp *RedisProxy) Set(key string, value string) errors.AppError {
	pool := rp.Pool()

	err := pool.Do(radix.Cmd(nil,"set", key, value))
	return errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
}

func (rp *RedisProxy) MSet(fieldValue map[string]string) errors.AppError {
	pool := rp.Pool()

	args := make([]string, 0)
	args = append(args)
	for k, v := range fieldValue {
		args = append(append(args, k), v)
	}

	err := pool.Do(radix.Cmd(nil,"mset", args...))
	return errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
}

func (rp *RedisProxy) Get(key string) (string, errors.AppError) {
	pool := rp.Pool()

	var result string
	err := pool.Do(radix.Cmd(&result,"get", key))

	if err != nil {
		return "", errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
	}
	return result, nil
}

func (rp *RedisProxy) MGet(keys ...string) ([]string, errors.AppError) {
	pool := rp.Pool()

	ls := make([]string, len(keys))
	err := pool.Do(radix.Cmd(&ls,"mget", keys...))

	if err != nil {
		return nil, errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
	}
	return ls, nil
}

func (rp *RedisProxy) Del(keys ...string) (int, errors.AppError) {
	pool := rp.Pool()

	var result int
	err := pool.Do(radix.Cmd(&result,"Del", keys...))
	return result, errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
}

func (rp *RedisProxy) Incrby(key string, v int64) (int64, errors.AppError) {

	pool := rp.Pool()

	var result int64
	err := pool.Do(radix.Cmd(&result,"INCRBY", key, strconv.FormatInt(v, 10)))
	return result, errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
}

func (rp *RedisProxy) Decrby(key string, v int64) (int64, errors.AppError) {

	pool := rp.Pool()

	var result int64
	err := pool.Do(radix.Cmd(&result,"DECRBY", key, strconv.FormatInt(v, 10)))
	return result, errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
}

func (rp *RedisProxy) Exist(key string) (bool, errors.AppError) {

	pool := rp.Pool()

	var result int64
	err := pool.Do(radix.Cmd(&result,"EXISTS", key))
	return result == 1, errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
}

func (rp *RedisProxy) Scan(index int64, match string, count int) (int64, []string, errors.AppError) {
	pool := rp.Pool()

	if match == "" {
		match = "*"
	}
	if count <= 0 || count > 200 {
		count = 30
	}
	var result []interface{}
	err := pool.Do(radix.Cmd(&result,"SCAN", strconv.FormatInt(index, 10), "MATCH", match, "COUNT", strconv.Itoa(count)))

	if err != nil {
		return 0, nil, errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
	}
	if len(result) != 2 {
		return 0, nil, nil
	}

	nextIndex, err := strconv.ParseInt(string(result[0].([]byte)), 10, 64)
	if err != nil {
		return 0, nil, errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
	}

	keyInts := result[1].([]interface{})
	keys := make([]string, 0)
	for _, v := range keyInts {
		keys = append(keys, string(v.([]byte)))
	}

	return nextIndex, keys, nil
}

func (rp *RedisProxy) Expire(key string, expire int) (bool, errors.AppError) {
	pool := rp.Pool()

	var result int64
	err := pool.Do(radix.Cmd(&result,"EXPIRE", key, strconv.Itoa(expire)))
	return result == 1, errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
}

func (rp *RedisProxy) TryGetDistributedLock(key string, v string) (bool, errors.AppError) {
	pool := rp.Pool()

	end := times.GetNowMillisecond() + _DistributedTimeOut*1000
	for times.GetNowMillisecond() <= end {
		var result int64
		err := pool.Do(_LockDistributedLuaScript.Cmd(&result, key, v, strconv.Itoa(_DistributedTimeOut)))
		if err != nil {
			return false, errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
		}
		if result == _DistributedSuccess {
			return true, nil
		}
		times.SleepMillisecond(80 + int64(rand.Int31n(30)))
	}

	return false, nil
}

func (rp *RedisProxy) ReleaseDistributedLock(key string, v string) (bool, errors.AppError) {
	pool := rp.Pool()

	var result int64
	err := pool.Do(_UnLockDistributedLuaScript.Cmd(&result, key, v))

	if err != nil {
		return false, errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
	}
	return result == _DistributedSuccess, nil
}

func (rp *RedisProxy) HGet(key, field string) (string, errors.AppError) {
	pool := rp.Pool()

	var result string
	err := pool.Do(radix.FlatCmd(&result,"HGET", key, field))

	if err != nil {
		return "", errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
	}
	return result, nil
}

func (rp *RedisProxy) HSet(key, field, value string) errors.AppError {
	pool := rp.Pool()

	var result string
	err := pool.Do(radix.FlatCmd(&result,"HSET", key, field, value))
	return errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
}

func (rp *RedisProxy) HDel(key string, fields ...interface{}) (int64, errors.AppError) {
	pool := rp.Pool()

	var result int64
	err := pool.Do(radix.FlatCmd(&result,"HDEL", key, fields...))

	if err != nil {
		return 0, errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
	}

	return result, nil
}

func (rp *RedisProxy) HExists(key, field string) (bool, errors.AppError) {
	pool := rp.Pool()

	var result int64
	err := pool.Do(radix.FlatCmd(&result,"HEXISTS", key, field))

	if err != nil {
		return false, errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
	}
	return result == 1, nil
}

func (rp *RedisProxy) HMGet(key string, fields ...interface{}) (map[string]string, errors.AppError) {
	result := map[string]string{}
	pool := rp.Pool()

	var rs []string
	err := pool.Do(radix.FlatCmd(&rs,"HMGET", key, fields...))

	if err != nil {
		return result, errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
	}
	if rs == nil {
		return result, nil
	}
	keys := make([]string, len(fields))
	for i, k := range fields {
		keys[i] = k.(string)
	}

	for i, data := range rs {
		result[keys[i]] = data
	}
	return result, nil
}

func (rp *RedisProxy) HMSet(key string, fieldValue map[string]string) errors.AppError {
	pool := rp.Pool()

	args := []interface{}{}
	for k, v := range fieldValue {
		args = append(append(args, k), v)
	}

	err := pool.Do(radix.FlatCmd(nil,"HMSET", key, args...))
	return errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
}

func (rp *RedisProxy) SetBit(key string, offset int, value int, ex int) errors.AppError {
	pool := rp.Pool()

	if err := pool.Do(radix.Cmd(nil,"SETBIT", key, strconv.Itoa(offset), strconv.Itoa(value))); err != nil {
		return errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
	}

	if ex > 0 {
		if err := pool.Do(radix.Cmd(nil,"EXPIRE", key, strconv.Itoa(ex))); err != nil {
			return errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
		}
	}
	return nil
}

func (rp *RedisProxy) GetBit(key string, offset int) (int, errors.AppError) {
	pool := rp.Pool()

	var result int
	err := pool.Do(radix.Cmd(&result, "GETBIT", key, strconv.Itoa(offset)))

	return result, errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
}

func (rp *RedisProxy) BitCount(key string) (int, errors.AppError) {
	pool := rp.Pool()

	var result int
	err := pool.Do(radix.Cmd(&result, "BITCOUNT", key))

	return result, errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
}

func (rp *RedisProxy) BitFieldGetU(key string, num int, start int) (int64, errors.AppError) {
	if num > 63 {
		return 0, errors.NewAppError(errors.ParamError, "num 不能大于63")
	}

	pool := rp.Pool()

	var result []int64
	err := pool.Do(radix.FlatCmd(&result, "BITFIELD", key, "GET", "u"+strconv.Itoa(num), start))

	if err != nil {
		return 0, errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
	}
	return result[0], nil
}

func (rp *RedisProxy) Pool() *radix.Pool {
	return GetRedisPool(rp.name)
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
