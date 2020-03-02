package database

import (
	"sync"
	"time"

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

	_MasterConfigName = "master"
)

// InitDataSource 初始化db
func InitDataSource(dbConfigs map[string]config.DBConfig) {

	poolMutex.Lock()
	defer poolMutex.Unlock()

	for k, v := range dbConfigs {
		initDataSourcePool(k, v)
	}
}

func initDataSourcePool(name string, config config.DBConfig) errors.AppError {

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
		panic(" db init fail. name:" + name + ". err:" + err.Error())
		return errors.NewAppErrorByExistError(errors.DbInitConnFail, err)
	}
	db.LogMode(config.LogMode)

	db.DB().SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.DB().SetMaxIdleConns(maxIdle)
	db.DB().SetMaxOpenConns(maxOpenConns)

	dbPools[name] = db

	return nil
}

// GetDb 获取默认数据库操作对象
func GetDb() *gorm.DB {
	return GetDbByName(_MasterConfigName)
}

// GetDbByName 获取数据库操作对象
func GetDbByName(name string) *gorm.DB {
	if dbPools == nil {
		panic(errors.NewAppError(errors.DbInitConnFail))
	}
	if v, ok := dbPools[name]; ok {
		return v
	}
	panic(errors.NewAppError(errors.DbInitConnFail))
}
