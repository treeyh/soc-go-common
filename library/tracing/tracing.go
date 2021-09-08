package tracing

import (
	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/treeyh/soc-go-common/core/config"
	"github.com/treeyh/soc-go-common/core/errors"
	"github.com/treeyh/soc-go-common/core/logger"
)

const (
	HttpComponent int32 = 2

	GormComponent  int32 = 700001
	RadixComponent int32 = 700002
)

var (
	log    = logger.Logger()
	tracer *go2sky.Tracer
	report go2sky.Reporter
)

// InitTracing 初始化sky walking , 其他中间件依赖，需要优先初始化
func InitTracing(traceConfig config.TraceConfig, appName string) {

	if !traceConfig.Enable {
		return
	}

	if traceConfig.Server != "" {
		_report, err := reporter.NewGRPCReporter(traceConfig.Server)
		if err != nil {
			log.Error("SkyWalking init fail." + err.Error())
			panic(errors.NewAppErrorByExistError(errors.SkyWalkingNotInit, err))
		}
		report = _report
	} else {
		_report, err := reporter.NewLogReporter()
		if err != nil {
			log.Error("SkyWalking init fail." + err.Error())
			panic(errors.NewAppErrorByExistError(errors.SkyWalkingNotInit, err))
		}
		report = _report
	}

	_tracer, err := go2sky.NewTracer(appName, go2sky.WithReporter(report))
	if err != nil {
		log.Error("SkyWalking init tracer fail." + err.Error())
		panic(errors.NewAppErrorByExistError(errors.SkyWalkingNotInit, err))
	}
	tracer = _tracer
}

// CloseReport 关闭report
func CloseReport() {
	if report != nil {
		report.Close()
	}
}

// GetReporter 获取Reporter对象
func GetReporter() go2sky.Reporter {
	return report
}

// GetTracer 获取tracer对象
func GetTracer() *go2sky.Tracer {
	return tracer
}
