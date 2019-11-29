package config

type LogConfig struct {
	LogPath      string
	Level        string
	FileSize     int64
	FileNum      int
	IsConsoleOut bool
	Tag          string
}
