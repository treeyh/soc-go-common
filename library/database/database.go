package database

import (
	"github.com/treeyh/soc-go-common/core/config"
	"github.com/treeyh/soc-go-common/core/consts"
	"github.com/treeyh/soc-go-common/core/errors"
	"github.com/treeyh/soc-go-common/core/logger"
	"github.com/treeyh/soc-go-common/library/tracing"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
	"strings"
	"sync"
	"time"
)

var (
	dbPools   = make(map[string]*gorm.DB)
	poolMutex sync.Mutex

	_MasterConfigName = "master"

	log = logger.Logger()
)

// InitDataSource 初始化db
func InitDataSource(dbConfigs map[string]config.DBConfig) {

	poolMutex.Lock()
	defer poolMutex.Unlock()

	for k, v := range dbConfigs {
		initDataSourcePool(k, v)
	}
}

// getLogLevel 获取日志记录级别
func getLogLevel(level string) glogger.LogLevel {

	if level == "info" {
		return glogger.Info
	} else if level == "warn" {
		return glogger.Warn
	} else if level == "error" {
		return glogger.Error
	} else if level == "silent" {
		return glogger.Silent
	} else {
		return glogger.Warn
	}
}

func initDataSourcePool(name string, config config.DBConfig) errors.AppError {

	maxIdle := 15
	if config.MaxIdleConns > 0 {
		maxIdle = config.MaxIdleConns
	}

	maxOpenConns := 30
	if config.MaxOpenConns > 0 {
		maxOpenConns = config.MaxOpenConns
	}

	connMaxLifetime := time.Duration(3600) * time.Second
	if config.ConnMaxLifetime > 0 {
		connMaxLifetime = time.Duration(config.ConnMaxLifetime) * time.Second
	}

	slowThreshold := 1000
	if config.SlowThreshold > 0 {
		slowThreshold = config.SlowThreshold
	}

	var glog glogger.Interface
	if config.LogMode {
		glog = NewLogger(log, glogger.Config{
			SlowThreshold: time.Duration(slowThreshold) * time.Millisecond,
			Colorful:      false,
			LogLevel:      getLogLevel(config.LogLevel),
		}, zap.String("socLog", "gorm"))
	} else {
		glog = nil
	}

	gconfig := &gorm.Config{
		SkipDefaultTransaction:                   false,
		NamingStrategy:                           nil,
		FullSaveAssociations:                     false,
		Logger:                                   glog,
		NowFunc:                                  nil,
		DryRun:                                   false,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
		DisableNestedTransaction:                 false,
		AllowGlobalUpdate:                        false,
		QueryFields:                              false,
		CreateBatchSize:                          0,
		ClauseBuilders:                           nil,
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  nil,
	}

	// mysql dsn := "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	// postgresql dsn := "user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var db *gorm.DB
	var err error
	if consts.DBTypePostgresql == config.Type {
		dbUrl := "postgres://" + strings.TrimSpace(config.DbUrl)

		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN: dbUrl,
		}), gconfig)
	} else {
		db, err = gorm.Open(mysql.New(mysql.Config{
			DSN: config.DbUrl,
		}), gconfig)
	}
	if err != nil {
		panic(" db init fail. name:" + name + ". err:" + err.Error())
		return errors.NewAppErrorByExistError(errors.DbInitConnFail, err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(" db get db fail. name:" + name + ". " + err.Error())
		return errors.NewAppErrorByExistError(errors.DbInitConnFail, err)
	}

	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxIdleTime(connMaxLifetime)

	// 截取 db url : user:password@tcp(host:3306)/dbName?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai
	addr := ""
	addrs := strings.Split(config.DbUrl, "@")
	if len(addrs) > 1 {
		addrs = strings.Split(addrs[1], "?")
		if len(addrs) > 1 {
			addr = strings.ReplaceAll(strings.ReplaceAll(addrs[0], "tcp(", ""), ")", "")
		}
	}

	db.Use(NewSkyWalkingPlugin(WithTracer(tracing.GetTracer()), WithUrl(addr)))

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
