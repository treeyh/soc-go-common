package logger

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/treeyh/soc-go-common/core/config"
	"github.com/treeyh/soc-go-common/tests"
	"testing"
)

func TestInitLogger(t *testing.T) {

	convey.Convey("log test", t, tests.TestStartUp(func() {
		log := Logger()

		fmt.Println(log.logConfig)
		log.Info("abc")

		logconfig := config.LogConfig{
			LogPath:          "./logs/log.log",
			Level:            "info",
			FileSize:         1024,
			FileNum:          20,
			IsSyncConsoleOut: true,
			Tag:              "aaa",
		}

		InitLogger(_logDefaultName, &logconfig, true)

		log.Info("123")
	}, nil))

}
