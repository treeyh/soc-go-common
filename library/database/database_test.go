package database

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/treeyh/soc-go-common/core/config"
	"github.com/treeyh/soc-go-common/core/utils/file"
	"github.com/treeyh/soc-go-common/core/utils/json"
	"path"
	"strconv"
	"testing"
	"time"
)

/**

CREATE TABLE `test_student` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
`name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '姓名',
`code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '学号',
`birthday` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '生日',
`height` DECIMAL(5,2) NOT NULL DEFAULT '0' COMMENT '身高',
`creator` bigint NOT NULL DEFAULT '0' COMMENT '创建人',
`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updator` bigint NOT NULL DEFAULT '0' COMMENT '更新人',
`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
`del_flag` tinyint NOT NULL DEFAULT '2' COMMENT '是否删除,1是,2否',
PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='学生';

*/

type StudentPo struct {

	// Id 主键
	Id int64 `gorm:"type:bigint(20);column:id" json:"id"`

	// Name 姓名
	Name string `gorm:"type:varchar(32);column:name" json:"name"`

	// Code 学号
	Code string `gorm:"type:varchar(64);column:code" json:"code"`

	// Birthday 生日
	Birthday time.Time `gorm:"type:datetime;column:birthday" json:"birthday"`

	// Height 身高
	Height float64 `gorm:"type:int(11);column:height" json:"height"`

	// Creator 创建人
	Creator int64 `gorm:"type:bigint(20);column:creator" json:"creator"`

	// CreateTime 创建时间
	CreateTime time.Time `gorm:"type:datetime;column:create_time" json:"createTime"`

	// Updator 更新人
	Updator int64 `gorm:"type:bigint(20);column:updator" json:"updator"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"type:datetime;column:update_time" json:"updateTime"`

	// DelFlag 是否删除,1是,2否
	DelFlag int `gorm:"type:tinyint(4);column:del_flag" json:"delFlag"`
}

func (*StudentPo) TableName() string {
	return "test_student"
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
	InitDataSource(wecmap)
}

func TestGetDB(t *testing.T) {
	initDb()

	now := time.Now().Unix()

	studentPo := StudentPo{
		Id:         now,
		Name:       "name_" + strconv.FormatInt(now, 10),
		Code:       "code_" + strconv.FormatInt(now, 10),
		Birthday:   time.Now(),
		Height:     1.81,
		Creator:    23,
		CreateTime: time.Now(),
		Updator:    45,
		UpdateTime: time.Now(),
		DelFlag:    2,
	}

	err := GetDb().Create(&studentPo).Error

	assert.NoError(t, err)

	fmt.Println(json.ToJsonIgnoreError(studentPo))

	var studentPo2 StudentPo
	row := GetDb().Where("id = ?", now).Where("del_flag = ?", 2).First(&studentPo2).RowsAffected

	assert.Equal(t, row, int64(1))
	assert.Equal(t, studentPo2.Id, now)

	var studentPo3 StudentPo
	row2 := GetDb().Where("id = ?", 1).Where("del_flag = ?", 2).First(&studentPo3).RowsAffected

	assert.NotEqual(t, row2, int64(1))
	assert.Equal(t, studentPo3.Id, int64(0))
}
