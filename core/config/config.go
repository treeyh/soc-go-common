package config

type LogConfig struct {
	LogPath      string
	Level        string
	FileSize     int64
	FileNum      int
	IsConsoleOut bool
	Tag          string
}

type AppConfig struct {
	RunMode   int
	CacheMode string
	Name      string
	AppCode   string
	AppKey    string
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

//mq配置
type MysqlConfig struct {
	Host     string
	Port     int
	Usr      string
	Pwd      string
	Database string
}
