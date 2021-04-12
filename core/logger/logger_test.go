package logger

import (
	"fmt"
	"github.com/treeyh/soc-go-common/core/config"
	"testing"
)

func TestInitLogger(t *testing.T) {

	log := Logger()

	fmt.Println(log.logConfig)
	log.Info("abc")

	logconfig := config.LogConfig{
		LogPath:          "./logs/log.log",
		Level:            "info",
		FileSize:         1024,
		FileNum:          20,
		IsSyncConsoleOut: true,
	}

	InitLogger(_logDefaultName, &logconfig, true)

	log.Info("123")

}
