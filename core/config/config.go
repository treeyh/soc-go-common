package config

type LogConfig struct {
	LogPath          string `json:"logPath"`
	Level            string `json:"level"`
	FileSize         int64  `json:"fileSize"`
	FileNum          int    `json:"fileNum"`
	IsSyncConsoleOut bool   `json:"isSyncConsoleOut"`
}

// RedisConfig redis配置
type RedisConfig struct {
	Host           string `json:"host"`
	Port           int    `json:"port"`
	User           string `json:"user"`
	Password       string `json:"password"`
	Database       int    `json:"database"`
	PoolSize       int    `json:"poolSize"`
	MaxIdleTimeout int    `json:"maxIdleTimeout"`
	ConnectTimeout int    `json:"connectTimeout"`
	ReadTimeout    int    `json:"readTimeout"`
	WriteTimeout   int    `json:"writeTimeout"`
}

// DBConfig 数据库配置
type DBConfig struct {
	// Type 数据库类型，目前仅支持 mysql
	Type  string `json:"type"`
	DbUrl string `json:"dbUrl"`
	// MaxIdleConns 连接池中的最大闲置连接数
	MaxIdleConns int `json:"maxIdleConns"`
	// MaxOpenConns 数据库的最大连接数量
	MaxOpenConns int `json:"maxOpenConns"`
	// ConnMaxLifetime 连接的最大可复用时间, 秒
	ConnMaxLifetime int `json:"connMaxLifetime"`
	// LogMode 是否记录日志
	LogMode bool `json:"logMode"`
	// SlowThreshold 慢日志记录阈值，毫秒
	SlowThreshold int `json:"slowThreshold"`
	// LogLevel 日志级别,类型如下：silent：无日志; error：错误日志; warn：警告日志（默认）; info：info日志
	LogLevel string `json:"logLevel"`
}

// TraceConfig trace配置
type TraceConfig struct {
	// Enable 是否开启
	Enable bool `json:"enable"`
	// Type 类型，仅支持 SkyWalking
	Type string `json:"type"`
	// Namespace 命名空间，默认为空，有值的话会加到header头的key前缀传递，如：sw8变为{namespace}-sw8
	Namespace string `json:"namespace"`
	// Server 服务地址
	Server string `json:"server"`
}

// WeChatConfig 微信配置
type WeChatConfig struct {
	// AppId 微信应用id
	AppId string `json:"appId"`
	// AppSecret 微信秘钥
	AppSecret string `json:"appSecret"`
	// Host 微信接口host
	Host string `json:"host"`
	// Type 类型，weapp：微信小程序 app：应用
	Type string `json:"type"`

	// Token 由开发者可以任意填写，用作生成签名
	Token string `json:"token"`
	// EncodingAESKey 消息加密密钥由 43 位字符组成，可随机修改，字符范围为 A-Z，a-z，0-9。
	EncodingAESKey string `json:"encodingAesKey"`
	// 	MessageEncodingType 消息加解密方式 1 明文模式, 2 安全模式 ,3 兼容模式
	MessageEncodingType int `json:"messageEncodingType"`
}

// ALiYunConfig 阿里云配置
type ALiYunConfig struct {
	// AccessKey 阿里云访问key
	AccessKey string `json:"accessKey"`
	// AccessKeySecret 阿里云访问key密钥
	AccessKeySecret string `json:"accessKeySecret"`
	// RegionId 阿里云访问区域
	RegionId string `json:"regionId"`
}

// I18nConfig 国际化配置
type I18nConfig struct {
	// Enable 是否开启
	Enable bool `json:"enable"`
	// Path 国际化文件配置目录
	Path string `json:"path"`
	// DefaultLang 默认语言
	DefaultLang string `json:"defaultLang"`
}
