package redis

import (
	"github.com/treeyh/soc-go-common/core/errors"
	"strconv"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/treeyh/soc-go-common/core/config"
)

var (
	redisPools = make(map[string]*redis.Pool)
	poolMutex  sync.Mutex
)

// InitRedisPool 初始化redis
func InitRedisPool(redisConfigs *map[string]config.RedisConfig) {

	poolMutex.Lock()
	defer poolMutex.Unlock()

	for k, v := range *redisConfigs {
		initRedisPool(k, v)
	}
}

func initRedisPool(name string, config config.RedisConfig) {

	maxIdle := 20
	if config.MaxIdle > 0 {
		maxIdle = config.MaxIdle
	}

	maxActive := 10
	if config.MaxActive > 0 {
		maxActive = config.MaxActive
	}

	maxIdleTimeout := 15
	if config.MaxIdleTimeout > 0 {
		maxIdleTimeout = config.MaxIdleTimeout
	}

	connectTimeout := time.Duration(3)
	if config.ConnectTimeout > 0 {
		connectTimeout = time.Duration(config.ConnectTimeout)
	}

	readTimeout := time.Duration(3)
	if config.ReadTimeout > 0 {
		readTimeout = time.Duration(config.ReadTimeout)
	}

	writeTimeout := time.Duration(3)
	if config.WriteTimeout > 0 {
		writeTimeout = time.Duration(config.WriteTimeout)
	}

	// 建立连接池
	redisPools[name] = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: time.Duration(maxIdleTimeout) * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", config.Host+":"+strconv.Itoa(config.Port),
				redis.DialPassword(config.Password),
				redis.DialDatabase(config.Database),
				redis.DialConnectTimeout(connectTimeout*time.Second),
				redis.DialReadTimeout(readTimeout*time.Second),
				redis.DialWriteTimeout(writeTimeout*time.Second))
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
}

func GetRedisConn(name string) redis.Conn {
	if redisPools == nil {
		panic(errors.NewAppError(errors.RedisNotInit))
	}
	if v, ok := redisPools[name]; ok {
		return v.Get()
	}
	panic(errors.NewAppError(errors.RedisNotInit))
}
