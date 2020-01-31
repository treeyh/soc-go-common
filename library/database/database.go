package database

import (
	"database/sql"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/treeyh/soc-go-common/core/config"
	"github.com/treeyh/soc-go-common/core/errors"
)

var (
	dbPools   = make(map[string]*gorm.DB)
	poolMutex sync.Mutex
)

// InitDataSource 初始化db
func InitDataSource(dbConfigs *map[string]config.DBConfig) {

	poolMutex.Lock()
	defer poolMutex.Unlock()

	for k, v := range *dbConfigs {
		initDataSourcePool(k, v)
	}
}

func initDataSourcePool(name string, config config.DBConfig) {

	maxIdle := 20
	if config.MaxIdleConns > 0 {
		maxIdle = config.MaxIdleConns
	}

	maxOpenConns := 50
	if config.MaxOpenConns > 0 {
		maxOpenConns = config.MaxOpenConns
	}

	connMaxLifetime := time.Duration(3600)
	if config.ConnMaxLifetime > 0 {
		connMaxLifetime = time.Duration(config.ConnMaxLifetime)
	}

	db, err := gorm.Open(config.Type, config.DBUrl)
	if err != nil {
		return errors.NewAppErrorByExistError(errors.DbInitConnFail, err)
	}
	db.LogMode(config.LogMode)

	db.DB().SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.DB().SetMaxIdleConns(maxIdle)
	db.DB().SetMaxOpenConns(maxOpenConns)

	dbPools[name] = db

	return nil
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
