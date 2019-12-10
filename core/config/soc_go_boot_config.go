package config

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

type SocBoot struct {
	App *AppConfig
}
