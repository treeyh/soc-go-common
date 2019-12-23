package consts

import (
	"os"
	"time"
)

//默认空时间
const BlankTime = "1970-01-01 00:00:00"
const BlankDate = "1970-01-01"

//全局日期格式
const AppDateFormat = "2006-01-02"
const AppTimeFormat = "2006-01-02 15:04:05"
const AppSystemTimeFormat = "2006-01-02T15:04:05Z"
const AppSystemTimeFormat8 = "2006-01-02T15:04:05+08:00"

var BlankTimeObject, _ = time.Parse(AppTimeFormat, BlankTime)

//默认空字符串
const BlankString = ""

// linux操作系统
const (
	// GOOSLinux linux 系统标识
	GOOSLinux = "linux"
	// GOOSMac mac 系统标识
	GOOSMac = "darwin"
	// GOOSAndroid android 系统标识
	GOOSAndroid = "android"
	// GOOSFreebsd freebsd 系统标识
	GOOSFreebsd = "freebsd"
	// GOOSOpenbsd openbsd 系统标识
	GOOSOpenbsd = "openbsd"
	// GOOSSolaris solaris 系统标识
	GOOSSolaris = "solaris"
	// GOOSWindows windows 系统标识
	GOOSWindows = "windows"
)

const (

	// DBTypeMysql 数据库类型
	DBTypeMysql = "mysql"
)

//const (
//	ContextTraceKey = "Tracer"
//	ContextContextKey = "ParentContext"
//)

const (

	// EvnRunName 环境变量名
	EvnSocBootRunName = "SOC_BOOT_RUN_ENV"

	// EnvLocal 本地环境
	EnvLocal = "local"
	// EnvDev 开发环境
	EnvDev = "dev"
	// EnvTest 测试环境
	EnvTest = "test"
	// EnvStag 预发布环境
	EnvStag = "stag"
	// EnvProd 生产环境
	EnvProd = "prod"
)

// GetCurrentEnv 获取当前环境值
func GetCurrentEnv() string {
	env := os.Getenv(EvnSocBootRunName)
	if "" == env {
		env = EnvLocal
	}
	return env
}
