package config

type LogConfig struct {
	LogPath          string
	Level            string
	FileSize         int64
	FileNum          int
	IsSyncConsoleOut bool
	Tag              string
}

//redis配置
type RedisConfig struct {
	Host           string
	Port           int
	Password       string
	Database       int
	MaxIdle        int
	MaxActive      int
	MaxIdleTimeout int
	ConnectTimeout int
	ReadTimeout    int
	WriteTimeout   int
}

//数据库配置
type DBConfig struct {
	Type  string
	DBUrl string
	// MaxIdleConns 连接池中的最大闲置连接数
	MaxIdleConns int
	// MaxOpenConns 数据库的最大连接数量
	MaxOpenConns int
	// ConnMaxLifetime 连接的最大可复用时间, 秒
	ConnMaxLifetime int
	// LogMode 是否记录日志
	LogMode bool
}

// TraceConfig trace配置
type TraceConfig struct {
	// Enable 是否开启
	Enable bool
	// Server 服务地址
	Server string
}

// WeChatConfig 微信配置
type WeChatConfig struct {
	// AppId 微信应用id
	AppId string
	// AppSecret 微信秘钥
	AppSecret string
	// Host 微信接口host
	Host string
	// Type 类型，weapp：微信小程序
	Type string
}
