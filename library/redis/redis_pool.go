package redis

import (
	"fmt"
	"github.com/SkyAPM/go2sky"
	"github.com/mediocregopher/radix/v3"
	"github.com/treeyh/soc-go-common/core/errors"
	"github.com/treeyh/soc-go-common/library/tracing"
	"strconv"
	"sync"
	"time"

	"github.com/treeyh/soc-go-common/core/config"
)

var (
	redisPools = make(map[string]*redisPool)
	poolMutex  sync.Mutex
)

type redisPool struct {
	name   string
	pool   *radix.Pool
	tracer *go2sky.Tracer
	addr   string
}

// InitRedisPool 初始化redis
func InitRedisPool(redisConfigs map[string]config.RedisConfig) {

	poolMutex.Lock()
	defer poolMutex.Unlock()

	for k, v := range redisConfigs {
		initRedisPool(k, v)
	}
}

func initRedisPool(name string, config config.RedisConfig) {

	poolSize := 15
	if config.PoolSize > 0 {
		poolSize = config.PoolSize
	}

	maxIdleTimeout := 10 * time.Second
	if config.MaxIdleTimeout > 0 {
		maxIdleTimeout = time.Duration(config.MaxIdleTimeout) * time.Second
	}

	connectTimeout := 3 * time.Second
	if config.ConnectTimeout > 0 {
		connectTimeout = time.Duration(config.ConnectTimeout) * time.Second
	}

	readTimeout := 3 * time.Second
	if config.ReadTimeout > 0 {
		readTimeout = time.Duration(config.ReadTimeout) * time.Second
	}

	writeTimeout := 3 * time.Second
	if config.WriteTimeout > 0 {
		writeTimeout = time.Duration(config.WriteTimeout) * time.Second
	}

	opts := make([]radix.DialOpt, 0)
	if config.User != "" && config.Password != "" {
		opts = append(opts, radix.DialAuthUser(config.User, config.Password))
	} else if config.Password != "" {
		opts = append(opts, radix.DialAuthPass(config.Password))
	}
	opts = append(opts, radix.DialSelectDB(config.Database), radix.DialConnectTimeout(connectTimeout),
		radix.DialReadTimeout(readTimeout), radix.DialWriteTimeout(writeTimeout), radix.DialTimeout(maxIdleTimeout))

	// DefaultConnFunc is a ConnFunc which will return a Conn for a redis instance
	// using sane defaults.
	connFunc := func(network, addr string) (radix.Conn, error) {
		return radix.Dial(network, addr, opts...)
	}

	poolOpts := make([]radix.PoolOpt, 0)
	poolOpts = append(poolOpts, radix.PoolConnFunc(connFunc))

	// 建立连接池
	addr := config.Host + ":" + strconv.Itoa(config.Port)
	_redisPool, err1 := radix.NewPool("tcp", addr, poolSize, poolOpts...)
	if err1 != nil {
		panic(fmt.Sprintf("init redis pool fail. %+v", err1))
	}

	redisPools[name] = &redisPool{
		name:   name,
		pool:   _redisPool,
		tracer: tracing.GetTracer(),
		addr:   addr,
	}
}

func GetRedisPool(name string) *redisPool {
	if redisPools == nil {
		panic(errors.NewAppError(errors.RedisNotInit))
	}
	if v, ok := redisPools[name]; ok {
		return v
	}
	panic(fmt.Sprintf("init redis poll fail. %+v", errors.NewAppError(errors.RedisNotInit)))
}
