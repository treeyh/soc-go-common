package database

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/treeyh/soc-go-common/core/config"
	"github.com/treeyh/soc-go-common/core/errors"
	"github.com/treeyh/soc-go-common/core/logger"
	"github.com/treeyh/soc-go-common/core/utils/file"
	"github.com/treeyh/soc-go-common/core/utils/json"
	"github.com/treeyh/soc-go-common/tests"
	"path"
	"testing"
	"time"
)

type ObjectIdPo struct {

	// Id 主键
	Id int64 `gorm:"type:bigint(20);column:id" json:"id"`

	// OrgId 组织id
	OrgId int64 `gorm:"type:bigint(20);column:org_id" json:"orgId"`

	// SysCode 系统编号
	SysCode string `gorm:"type:varchar(32);column:sys_code" json:"sysCode"`

	// Code 对象编号
	Code string `gorm:"type:varchar(64);column:code" json:"code"`

	// MaxId 当前最大编号
	MaxId int64 `gorm:"type:bigint(20);column:max_id" json:"maxId"`

	// Step 步长
	Step int `gorm:"type:int(11);column:step" json:"step"`

	// Creator 创建人
	Creator int64 `gorm:"type:bigint(20);column:creator" json:"creator"`

	// CreateTime 创建时间
	CreateTime time.Time `gorm:"type:datetime;column:create_time" json:"createTime"`

	// Updator 更新人
	Updator int64 `gorm:"type:bigint(20);column:updator" json:"updator"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"type:datetime;column:update_time" json:"updateTime"`

	// Version 乐观锁
	Version int `gorm:"type:int(11);column:version" json:"version"`

	// DelFlag 是否删除,1是,2否
	DelFlag int `gorm:"type:tinyint(4);column:del_flag" json:"delFlag"`
}

func (*ObjectIdPo) TableName() string {
	return "sys_object_id"
}

func initDb() {
	dbJson, _ := file.ReadSmallFile(path.Join(file.GetCurrentPath(), "..", "..", "tests", "database.log"))

	dbConfig := &config.DBConfig{}
	err := json.FromJson(*dbJson, dbConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	wecmap := make(map[string]config.DBConfig)
	wecmap[_MasterConfigName] = *dbConfig
	InitDataSource(&wecmap)
}

func TestGetDB(t *testing.T) {
	convey.Convey("log test", t, tests.TestStartUp(func() {

		objectId := ObjectIdPo{
			Id:         1,
			OrgId:      1,
			SysCode:    "syscode",
			Code:       "user",
			MaxId:      1000,
			Step:       200,
			Creator:    23,
			CreateTime: time.Now(),
			Updator:    45,
			UpdateTime: time.Now(),
			Version:    1,
			DelFlag:    2,
		}

		err := GetDb().Create(&objectId).Error
		//t := GetDb().NewRecord(&objectId)
		//fmt.Println(t)
		fmt.Println(err)
		logger.Logger().Info(errors.NewAppErrorByExistError(errors.DbOperationFail, err))

		fmt.Println(json.ToJsonIgnoreError(objectId))

	}, initDb))
}

func TestGetDB2(t *testing.T) {
	convey.Convey("log test", t, tests.TestStartUp(func() {

		var objId ObjectIdPo

		row := GetDb().Where("org_id = ? AND code = ?", 1, "us1er").First(&objId).RowsAffected

		fmt.Println(row)
		fmt.Println(json.ToJsonIgnoreError(objId))

	}, initDb))
}
