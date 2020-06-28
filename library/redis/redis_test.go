package redis

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/treeyh/soc-go-common/core/config"
	"testing"
)

func getTestRedisConfigMap() map[string]config.RedisConfig {
	redisConfigs := map[string]config.RedisConfig{
		"master": config.RedisConfig{
			//Host:     "192.168.1.148",
			Host:     "127.0.0.1",
			Port:     6379,
			Password: "",
			Database: 0,
		},
		"node1": config.RedisConfig{
			//Host:     "192.168.1.148",
			Host:     "192.168.1.101",
			Port:     6379,
			Password: "",
			Database: 1,
		},
	}

	return redisConfigs
}

func TestGetProxy(t *testing.T) {

	convey.Convey("Test TestRedisGetProxy", t, func() {
		redisConfigs := getTestRedisConfigMap()
		InitRedisPool(redisConfigs)

		key := "test:cache"
		val := "2222"
		err := GetProxy().SetEx(key, val, 60)
		fmt.Println(err)

		convey.ShouldEqual(err, nil)
		v, e := GetProxy().Get(key)
		fmt.Println(v, e)
		convey.ShouldEqual(v, val)

		val = "333"
		err = GetProxyByName("node1").SetEx(key, val, 60)
		fmt.Println(err)

		convey.ShouldEqual(err, nil)

		v, e = GetProxyByName("node1").Get(key)
		fmt.Println(v, e)

		convey.ShouldEqual(v, val)
	})
}

func TestMGet(t *testing.T) {

	convey.Convey("Test TestMGet", t, func() {
		redisConfigs := getTestRedisConfigMap()
		InitRedisPool(redisConfigs)

		kvs := map[string]string{
			"test:aaa": "aaa",
			"test:bbb": "bbb",
			"test:ccc": "ccc",
		}
		err := GetProxy().MSet(kvs)
		fmt.Println(err)

		v, e := GetProxy().Get("test:bbb")
		fmt.Println(v, e)
		convey.ShouldEqual(v, "bbb")
		v, e = GetProxy().Get("test:aaa")
		fmt.Println(v, e)
		convey.ShouldEqual(v, "aaa")

		keys := make([]interface{}, 3)
		keys[0] = "test:aaa"
		keys[1] = "test:bbb"
		keys[2] = "test:ccc"

		vv, e := GetProxy().MGet(keys...)
		fmt.Println(vv, e)
		convey.ShouldEqual(len(vv), 3)
	})
}

func TestRedisProxy_SetBit(t *testing.T) {
	convey.Convey("Test SetBit", t, func() {
		redisConfigs := getTestRedisConfigMap()
		InitRedisPool(redisConfigs)

		key := "testbit"
		err := GetProxy().SetBit(key, 3, 1, 3600)
		fmt.Println(err)

		v, e := GetProxy().BitFieldGetU(key, 31, 0)

		fmt.Println(v)
		fmt.Println(e)

	})
}

func TestRedisProxy_Scan(t *testing.T) {
	convey.Convey("Test Scan", t, func() {
		redisConfigs := getTestRedisConfigMap()
		InitRedisPool(redisConfigs)

		nextIndex, v, e := GetProxy().Scan(0, "a*", 10)

		fmt.Println(nextIndex)
		fmt.Println(v)
		fmt.Println(e)

	})
}
