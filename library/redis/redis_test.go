package redis

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/treeyh/soc-go-common/core/config"
	"testing"
)

func getTestRedisConfigMap() *map[string]config.RedisConfig {
	redisConfigs := &map[string]config.RedisConfig{
		"Master": config.RedisConfig{
			Host:     "192.168.1.148",
			Port:     6379,
			Password: "",
			Database: 7,
		},
		//"node1" : config.RedisConfig{
		//	Host:           "192.168.1.148",
		//	Port:           6379,
		//	Password:       "",
		//	Database:       8,
		//},
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

		convey.ShouldEqual(err, nil)
		v, e := GetProxy().Get("test:bbb")
		fmt.Println(v, e)
		convey.ShouldEqual(v, "bbb")
		v, e = GetProxy().Get("test:aaa")
		fmt.Println(v, e)
		convey.ShouldEqual(v, "aaa")

		keys := make([]interface{}, len(kvs))
		keys[0] = "test:aaa"
		keys[1] = "test:bbb"
		keys[2] = "test:ccc"

		vv, e := GetProxy().MGet(keys)
		fmt.Println(vv, e)
		convey.ShouldEqual(len(vv), 3)
	})
}
