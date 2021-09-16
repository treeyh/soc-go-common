package consts

import (
	"os"
	"time"
)

// 默认空时间

const BlankTime = "1970-01-01 00:00:00"
const BlankDate = "1970-01-01"

//全局日期格式

const AppDateFormat = "2006-01-02"
const AppDateFormat2 = "20060102"
const AppMonthFormat = "200601"
const AppTimeFormat = "2006-01-02 15:04:05"
const AppTimeFormatMillisecond = "2006-01-02 15:04:05.999"
const AppTimeFormat2 = "20060102150405"
const AppSystemTimeFormat = "2006-01-02T15:04:05Z07:00"
const AppSystemTimeFormat8 = "2006-01-02T15:04:05.999Z07:00"

var BlankTimeObject, _ = time.Parse(AppTimeFormat, BlankTime)

// BlankString 默认空字符串
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

const (

	// EvnSocRunName 环境变量名
	EvnSocRunName = "SOC_RUN_ENV"

	// EnvLocal 本地环境
	EnvLocal = "local"
	// EnvDev 开发环境
	EnvDev = "dev"
	// EnvUnitTest 单元测试环境
	EnvUnitTest = "utest"
	// EnvTest 测试环境
	EnvTest = "test"
	// EnvStag 预发布环境
	EnvStag = "stag"
	// EnvProd 生产环境
	EnvProd = "prod"
)

var _env = ""

// GetCurrentEnv 获取当前环境值
func GetCurrentEnv() string {
	if _env == "" {
		_env = os.Getenv(EvnSocRunName)
		if "" == _env {
			_env = EnvLocal
		}
		return _env
	}
	return _env
}

const (
	// HeaderTraceIdKey 用于http header
	HeaderTraceIdKey = "soc-trace-id"

	// HeaderAppCodeKey 用于http header
	HeaderAppCodeKey = "soc-app-code"

	// HeaderPartnerCodeKey 合作方id的http header
	HeaderPartnerCodeKey = "soc-partner-code"

	// HeaderSignKey 请求签名 http header
	HeaderSignKey = "soc-sign"

	// HeaderTimestampKey 请求时间戳 http header
	HeaderTimestampKey = "soc-timestamp"

	// HeaderSignPolicyKey 请求签名策略 http header
	HeaderSignPolicyKey = "soc-sign-policy"

	// ContextTracerKey traceId 的 context key
	ContextTracerKey = "soc_trace_id"

	// ContextParentSpanKey parent span 的 context key
	ContextParentSpanKey = "soc_parent_span"
)

// LineSep 换行符
const (
	LineSep = "\n"

	// EmptyStr 空字符串
	EmptyStr = ""
)
