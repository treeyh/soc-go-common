package logger

import (
	"fmt"
	"github.com/treeyh/soc-go-common/core/config"
	"github.com/treeyh/soc-go-common/core/errors"
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

	err := errors.NewAppError(errors.ParamError, "body")
	log.Error2(err, "test error")

}
