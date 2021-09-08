package redis

import (
	"context"
	"github.com/SkyAPM/go2sky"
	"github.com/mediocregopher/radix/v3"
	"github.com/treeyh/soc-go-common/core/errors"
	"github.com/treeyh/soc-go-common/core/logger"
	"github.com/treeyh/soc-go-common/core/utils/times"
	"github.com/treeyh/soc-go-common/library/tracing"
	"math/rand"
	"reflect"
	v3 "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
	"strconv"
	"strings"
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

	log = logger.Logger()
)

type RedisProxy struct {
	name string
}

func (rp *RedisProxy) ZAdd(ctx context.Context, key string, score float64, value string) errors.AppError {

	return rp.do(ctx, radix.Cmd(nil, "zadd", key, strconv.FormatFloat(score, 'f', 10, 64), value))
}

func (rp *RedisProxy) SetEx(ctx context.Context, key string, value string, ex int) errors.AppError {
	return rp.do(ctx, radix.Cmd(nil, "setex", key, strconv.Itoa(ex), value))
}

func (rp *RedisProxy) Set(ctx context.Context, key string, value string) errors.AppError {
	return rp.do(ctx, radix.Cmd(nil, "set", key, value))
}

func (rp *RedisProxy) MSet(ctx context.Context, fieldValue map[string]string) errors.AppError {
	args := make([]string, 0)
	args = append(args)
	for k, v := range fieldValue {
		args = append(append(args, k), v)
	}

	return rp.do(ctx, radix.Cmd(nil, "mset", args...))
}

func (rp *RedisProxy) Get(ctx context.Context, key string) (string, errors.AppError) {
	var result string
	err := rp.do(ctx, radix.Cmd(&result, "get", key))

	if err != nil {
		return "", err
	}
	return result, nil
}

func (rp *RedisProxy) MGet(ctx context.Context, keys ...string) ([]string, errors.AppError) {
	ls := make([]string, len(keys))
	err := rp.do(ctx, radix.Cmd(&ls, "mget", keys...))

	if err != nil {
		return nil, errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
	}
	return ls, nil
}

func (rp *RedisProxy) Del(ctx context.Context, keys ...string) (int, errors.AppError) {
	var result int
	err := rp.do(ctx, radix.Cmd(&result, "Del", keys...))
	return result, err
}

func (rp *RedisProxy) Incrby(ctx context.Context, key string, v int64) (int64, errors.AppError) {

	var result int64
	err := rp.do(ctx, radix.Cmd(&result, "INCRBY", key, strconv.FormatInt(v, 10)))
	return result, err
}

func (rp *RedisProxy) Decrby(ctx context.Context, key string, v int64) (int64, errors.AppError) {

	var result int64
	err := rp.do(ctx, radix.Cmd(&result, "DECRBY", key, strconv.FormatInt(v, 10)))
	return result, err
}

func (rp *RedisProxy) Exist(ctx context.Context, key string) (bool, errors.AppError) {

	var result int64
	err := rp.do(ctx, radix.Cmd(&result, "EXISTS", key))
	return result == 1, err
}

func (rp *RedisProxy) Scan(ctx context.Context, index int64, match string, count int) (int64, []string, errors.AppError) {

	if match == "" {
		match = "*"
	}
	if count <= 0 || count > 200 {
		count = 30
	}
	var result []interface{}
	err := rp.do(ctx, radix.Cmd(&result, "SCAN", strconv.FormatInt(index, 10), "MATCH", match, "COUNT", strconv.Itoa(count)))

	if err != nil {
		return 0, nil, err
	}
	if len(result) != 2 {
		return 0, nil, nil
	}

	nextIndex, err1 := strconv.ParseInt(string(result[0].([]byte)), 10, 64)
	if err1 != nil {
		return 0, nil, errors.NewAppErrorByExistError(errors.RedisOperationFail, err1)
	}

	keyInts := result[1].([]interface{})
	keys := make([]string, 0)
	for _, v := range keyInts {
		keys = append(keys, string(v.([]byte)))
	}

	return nextIndex, keys, nil
}

func (rp *RedisProxy) Expire(ctx context.Context, key string, expire int) (bool, errors.AppError) {
	var result int64
	err := rp.do(ctx, radix.Cmd(&result, "EXPIRE", key, strconv.Itoa(expire)))
	return result == 1, err
}

func (rp *RedisProxy) TryGetDistributedLock(ctx context.Context, key string, v string) (bool, errors.AppError) {
	end := times.GetNowMillisecond() + _DistributedTimeOut*1000
	for times.GetNowMillisecond() <= end {
		var result int64
		err := rp.do(ctx, _LockDistributedLuaScript.Cmd(&result, key, v, strconv.Itoa(_DistributedTimeOut)))
		if err != nil {
			return false, err
		}
		if result == _DistributedSuccess {
			return true, nil
		}
		times.SleepMillisecond(80 + int64(rand.Int31n(30)))
	}

	return false, nil
}

func (rp *RedisProxy) ReleaseDistributedLock(ctx context.Context, key string, v string) (bool, errors.AppError) {
	var result int64
	err := rp.do(ctx, _UnLockDistributedLuaScript.Cmd(&result, key, v))

	if err != nil {
		return false, err
	}
	return result == _DistributedSuccess, nil
}

func (rp *RedisProxy) HGet(ctx context.Context, key, field string) (string, errors.AppError) {
	var result string
	err := rp.do(ctx, radix.FlatCmd(&result, "HGET", key, field))

	if err != nil {
		return "", err
	}
	return result, nil
}

func (rp *RedisProxy) HSet(ctx context.Context, key, field, value string) errors.AppError {
	var result string
	return rp.do(ctx, radix.FlatCmd(&result, "HSET", key, field, value))
}

func (rp *RedisProxy) HDel(ctx context.Context, key string, fields ...interface{}) (int64, errors.AppError) {
	var result int64
	err := rp.do(ctx, radix.FlatCmd(&result, "HDEL", key, fields...))

	if err != nil {
		return 0, err
	}

	return result, nil
}

func (rp *RedisProxy) HExists(ctx context.Context, key, field string) (bool, errors.AppError) {
	var result int64
	err := rp.do(ctx, radix.FlatCmd(&result, "HEXISTS", key, field))

	if err != nil {
		return false, err
	}
	return result == 1, nil
}

func (rp *RedisProxy) HMGet(ctx context.Context, key string, fields ...interface{}) (map[string]string, errors.AppError) {
	result := map[string]string{}

	var rs []string
	err := rp.do(ctx, radix.FlatCmd(&rs, "HMGET", key, fields...))

	if err != nil {
		return result, err
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

func (rp *RedisProxy) HMSet(ctx context.Context, key string, fieldValue map[string]string) errors.AppError {
	args := []interface{}{}
	for k, v := range fieldValue {
		args = append(append(args, k), v)
	}

	return rp.do(ctx, radix.FlatCmd(nil, "HMSET", key, args...))
}

func (rp *RedisProxy) SetBit(ctx context.Context, key string, offset int, value int, ex int) errors.AppError {

	if err := rp.do(ctx, radix.Cmd(nil, "SETBIT", key, strconv.Itoa(offset), strconv.Itoa(value))); err != nil {
		return err
	}

	if ex > 0 {
		if err := rp.do(ctx, radix.Cmd(nil, "EXPIRE", key, strconv.Itoa(ex))); err != nil {
			return err
		}
	}
	return nil
}

func (rp *RedisProxy) GetBit(ctx context.Context, key string, offset int) (int, errors.AppError) {
	var result int
	err := rp.do(ctx, radix.Cmd(&result, "GETBIT", key, strconv.Itoa(offset)))

	return result, err
}

func (rp *RedisProxy) BitCount(ctx context.Context, key string) (int, errors.AppError) {
	var result int
	err := rp.do(ctx, radix.Cmd(&result, "BITCOUNT", key))

	return result, err
}

func (rp *RedisProxy) BitFieldGetU(ctx context.Context, key string, num int, start int) (int64, errors.AppError) {
	if num > 63 {
		return 0, errors.NewAppError(errors.ParamError, "num 不能大于63")
	}

	var result []int64
	err := rp.do(ctx, radix.FlatCmd(&result, "BITFIELD", key, "GET", "u"+strconv.Itoa(num), start))

	if err != nil {
		return 0, err
	}
	return result[0], nil
}

func (rp *RedisProxy) Pool() *redisPool {
	return GetRedisPool(rp.name)
}

func (rp RedisProxy) IsEmpty() bool {
	return reflect.DeepEqual(rp, RedisProxy{})
}

func (rp *RedisProxy) do(ctx context.Context, a radix.Action) errors.AppError {

	pool := rp.Pool()

	var redisSpan go2sky.Span
	var err error
	if pool.tracer != nil {
		// sky walking span
		peer := strings.Join(a.Keys(), " ")
		if peer == "" {
			peer = "No Peer"
		}
		redisSpan, err = pool.tracer.CreateExitSpan(ctx, pool.addr, peer, func(key, value string) error {
			return nil
		})
		if err != nil {
			log.ErrorCtx2(ctx, err, errors.SkyWalkingSpanNotInit.Error())
		}
		redisSpan.SetComponent(tracing.RadixComponent)
		redisSpan.SetSpanLayer(v3.SpanLayer_Cache)
	}

	err1 := pool.pool.Do(a)

	if redisSpan != nil {
		redisSpan.Tag("cache.type", "redis")
		redisSpan.Tag("cache.addr", pool.addr)
		defer redisSpan.End()
	}
	if err1 != nil {
		log.ErrorCtx2(ctx, err1, errors.RedisOperationFail.Error())
		return errors.NewAppErrorByExistError(errors.RedisOperationFail, err1)
	}

	return nil
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
