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
}

// AppConfig 应用配置
type AppConfig struct {
	Name    string
	Server  *ServerConfig
	AppCode string
	AppKey  string
}

// ServerConfig 服务配置
type ServerConfig struct {
	Port        int
	ContextPath string
}
