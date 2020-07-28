package config

type LogConfig struct {
	LogPath          string `json:"logPath"`
	Level            string `json:"level"`
	FileSize         int64  `json:"fileSize"`
	FileNum          int    `json:"fileNum"`
	IsSyncConsoleOut bool   `json:"isSyncConsoleOut"`
	Tag              string `json:"tag"`
}

//redis配置
type RedisConfig struct {
	Host           string `json:"host"`
	Port           int    `json:"port"`
	Password       string `json:"password"`
	Database       int    `json:"database"`
	MaxIdle        int    `json:"maxIdle"`
	MaxIdleTimeout int    `json:"maxIdleTimeout"`
	ConnectTimeout int    `json:"connectTimeout"`
	ReadTimeout    int    `json:"readTimeout"`
	WriteTimeout   int    `json:"writeTimeout"`
}

//数据库配置
type DBConfig struct {
	Type  string `json:"type"`
	DBUrl string `json:"dbUrl"`
	// MaxIdleConns 连接池中的最大闲置连接数
	MaxIdleConns int `json:"maxIdleConns"`
	// MaxOpenConns 数据库的最大连接数量
	MaxOpenConns int `json:"maxOpenConns"`
	// ConnMaxLifetime 连接的最大可复用时间, 秒
	ConnMaxLifetime int `json:"connMaxLifetime"`
	// LogMode 是否记录日志
	LogMode bool `json:"logMode"`
}

// TraceConfig trace配置
type TraceConfig struct {
	// Enable 是否开启
	Enable bool `json:"enable"`
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
	// Type 类型，weapp：微信小程序
	Type string `json:"type"`

	// Token 由开发者可以任意填写，用作生成签名
	Token string `json:"token"`
	// EncodingAESKey 消息加密密钥由 43 位字符组成，可随机修改，字符范围为 A-Z，a-z，0-9。
	EncodingAESKey string `json:"encodingAesKey"`
	// 	MessageEncodingType 消息加解密方式 1 明文模式, 2 安全模式 ,3 兼容模式
	MessageEncodingType int `json:"messageEncodingType"`
}
