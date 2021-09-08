package redis

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/treeyh/soc-go-common/core/config"
	"github.com/treeyh/soc-go-common/core/utils/json"
	"strconv"
	"testing"
	"time"
)

const (
	_master = "master"
	_node1  = "node1"
)

func getTestRedisConfigMap() map[string]config.RedisConfig {
	redisConfigs := map[string]config.RedisConfig{
		_master: config.RedisConfig{
			Host:     "192.168.1.101",
			Port:     6379,
			Password: "",
			Database: 0,
		},
		_node1: config.RedisConfig{
			Host:     "192.168.1.101",
			Port:     6379,
			Password: "",
			Database: 1,
		},
	}

	return redisConfigs
}

func TestGetProxy(t *testing.T) {

	InitRedisPool(getTestRedisConfigMap())

	key := "test:cache"
	val := "2222"
	err := GetProxy().SetEx(context.Background(), key, val, 60)
	assert.NoError(t, err, "redis SetEx error not nil %+v", err)
	v, e := GetProxy().Get(context.Background(), key)
	fmt.Println(v, e)
	assert.Equal(t, v, val, "redis Get not equal %s  %s.", v, val)

	val = "333"
	err = GetProxyByName(_node1).SetEx(context.Background(), key, val, 60)

	assert.NoError(t, err, "redis SetEx error not nil")

	v, e = GetProxyByName(_node1).Get(context.Background(), key)
	fmt.Println(v, e)

	assert.Equal(t, v, val, "redis Get not equal %s  %s.", v, val)
}

func TestRedisProxy_SetEx(t *testing.T) {

	InitRedisPool(getTestRedisConfigMap())

	key := "test:SetEx"
	val := "a"
	sleep := 2

	err := GetProxy().SetEx(context.Background(), key, val, sleep)
	assert.NoError(t, err)

	v, err := GetProxy().Get(context.Background(), key)
	assert.NoError(t, err)
	assert.Equal(t, val, v)

	time.Sleep(time.Duration(sleep+1) * time.Second)

	v2, err := GetProxy().Get(context.Background(), key)
	assert.NoError(t, err)
	assert.Equal(t, v2, "")

	err = GetProxy().SetEx(context.Background(), key, val, sleep)
	assert.NoError(t, err)

	v1, err := GetProxy().Get(context.Background(), key)
	assert.NoError(t, err)
	assert.Equal(t, val, v1)

	GetProxy().Expire(context.Background(), key, 3+sleep)

	time.Sleep(time.Duration(sleep+1) * time.Second)

	v4, err := GetProxy().Get(context.Background(), key)
	assert.NoError(t, err)
	assert.Equal(t, v4, val)
}

func TestRedisProxy_Del(t *testing.T) {

	InitRedisPool(getTestRedisConfigMap())

	for i := 0; i < 10; i++ {

		key := fmt.Sprintf("test:TestRedisProxy_Del_%d", i)
		val := "a"

		err := GetProxy().Set(context.Background(), key, val)
		assert.NoError(t, err)

		v, err := GetProxy().Get(context.Background(), key)
		assert.NoError(t, err)
		assert.Equal(t, val, v)

		count, err := GetProxy().Del(context.Background(), key)
		assert.NoError(t, err)
		assert.Equal(t, count, 1)

		v2, err := GetProxy().Get(context.Background(), key)
		assert.NoError(t, err)
		assert.Equal(t, v2, "")
	}
}

func TestRedisProxy_Incrby(t *testing.T) {

	InitRedisPool(getTestRedisConfigMap())

	key := "test:TestRedisProxy_Incrby"

	count := int64(100)

	c1, err := GetProxy().Incrby(context.Background(), key, count)
	assert.NoError(t, err)
	assert.Equal(t, c1, count)

	c2, err := GetProxy().Incrby(context.Background(), key, count)
	assert.NoError(t, err)
	assert.Equal(t, c2, count+c1)

	c3, err := GetProxy().Decrby(context.Background(), key, count)
	assert.NoError(t, err)
	assert.Equal(t, c3, c2-count)

	c4, err := GetProxy().Get(context.Background(), key)
	assert.NoError(t, err)
	assert.Equal(t, c4, strconv.FormatInt(c3, 10))

	c5, err := GetProxy().Exist(context.Background(), key)
	assert.NoError(t, err)
	assert.True(t, c5)

	c, err := GetProxy().Del(context.Background(), key)
	assert.NoError(t, err)
	assert.Equal(t, c, 1)
}

func TestRedisProxy_Scan(t *testing.T) {

	InitRedisPool(getTestRedisConfigMap())

	for i := 0; i < 53; i++ {

		key := fmt.Sprintf("test:TestRedisProxy_Scan_%d", i)
		val := "a"

		err := GetProxy().Set(context.Background(), key, val)
		assert.NoError(t, err)
	}

	index := int64(0)
	for i := 0; i < 100; i++ {
		ii, ss, err := GetProxy().Scan(context.Background(), index, "", 10)
		assert.NoError(t, err)
		t.Log("index:", ii, ";ss:", ss)
		index = ii
		if ii == 0 {
			break
		}
	}

	for i := 0; i < 53; i++ {

		key := fmt.Sprintf("test:TestRedisProxy_Scan_%d", i)
		val := "a"

		err := GetProxy().Set(context.Background(), key, val)
		assert.NoError(t, err)
	}
}

func TestRedisProxy_TryGetDistributedLock(t *testing.T) {

	InitRedisPool(getTestRedisConfigMap())

	key := "test:TestRedisProxy_TryGetDistributedLock"
	val := "a"
	rs, err := GetProxy().TryGetDistributedLock(context.Background(), key, val)
	assert.NoError(t, err)
	assert.True(t, rs)

	rs2, err := GetProxy().ReleaseDistributedLock(context.Background(), key, val)
	assert.NoError(t, err)
	assert.True(t, rs2)

	rs3, err := GetProxy().TryGetDistributedLock(context.Background(), key, val)
	assert.NoError(t, err)
	assert.True(t, rs3)

	err = GetProxy().SetEx(context.Background(), key, "b", 10)
	assert.NoError(t, err)

	rs4, err := GetProxy().ReleaseDistributedLock(context.Background(), key, val)
	assert.NoError(t, err)
	assert.False(t, rs4)

}

func TestRedisProxy_HMGet(t *testing.T) {
	InitRedisPool(getTestRedisConfigMap())

	key := "test:TestRedisProxy_HMGet"

	sm := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}

	err := GetProxy().HMSet(context.Background(), key, sm)
	assert.NoError(t, err)

	sm2, err := GetProxy().HMGet(context.Background(), key, "a", "b", "c")
	assert.NoError(t, err)
	assert.Equal(t, sm["a"], sm2["a"])
	assert.Equal(t, sm["b"], sm2["b"])
	assert.Equal(t, sm["c"], sm2["c"])

	GetProxy().Del(context.Background(), key)
}

func TestRedisProxy_SetBit(t *testing.T) {

	InitRedisPool(getTestRedisConfigMap())

	key := "test:TestRedisProxy_SetBit"

	err := GetProxy().SetBit(context.Background(), key, 0, 1, 10)
	assert.NoError(t, err)

	c, err := GetProxy().GetBit(context.Background(), key, 0)
	assert.NoError(t, err)
	assert.Equal(t, c, 1)

	c2, err := GetProxy().GetBit(context.Background(), key, 4)
	assert.NoError(t, err)
	assert.Equal(t, c2, 0)

	c3, err := GetProxy().BitCount(context.Background(), key)
	assert.NoError(t, err)
	assert.Equal(t, c3, 1)

	c4, err := GetProxy().BitFieldGetU(context.Background(), key, 1, 0)
	assert.NoError(t, err)
	assert.Equal(t, c4, int64(1))

	GetProxy().Del(context.Background(), key)
}

func TestMGet(t *testing.T) {

	InitRedisPool(getTestRedisConfigMap())

	kvs := map[string]string{
		"test:aaa": "aaa",
		"test:bbb": "bbb",
		"test:ccc": "ccc",
	}
	err := GetProxy().MSet(context.Background(), kvs)
	fmt.Println(err)

	v, e := GetProxy().Get(context.Background(), "test:bbb")
	fmt.Println(v, e)

	assert.Equal(t, v, "bbb", "redis Get not equal %s  %s.", v, "bbb")
	v, e = GetProxy().Get(context.Background(), "test:aaa")
	fmt.Println(v, e)
	assert.Equal(t, v, "aaa", "redis Get not equal %s  %s.", v, "aaa")

	keys := make([]string, 4)
	keys[0] = "test:aaa"
	keys[1] = "test:ddd"
	keys[2] = "test:bbb"
	keys[3] = "test:ccc"

	vv, e := GetProxy().MGet(context.Background(), keys...)
	fmt.Println(json.ToJsonIgnoreError(vv), e)
	assert.Equal(t, len(vv), len(keys), "redis Get not equal %d  %d.", len(vv), len(keys))

	GetProxy().Del(context.Background(), keys...)
}
