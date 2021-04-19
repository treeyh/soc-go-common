package database

import (
	"context"
	"fmt"
	"github.com/treeyh/soc-go-common/core/logger"
	"go.uber.org/zap"
	glog "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

type Logger struct {
	glog.Config
	LogLevel glog.LogLevel
	log      *logger.AppLogger
	zapTags  []zap.Field

	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

func NewLogger(log *logger.AppLogger, config glog.Config, zapFields ...zap.Field) glog.Interface {
	var (
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)

	return &Logger{
		log:          log,
		Config:       config,
		LogLevel:     config.LogLevel,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
		zapTags:      zapFields,
	}
}

// LogMode log mode
func (l *Logger) LogMode(level glog.LogLevel) glog.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

func (l Logger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= glog.Info {
		m := fmt.Sprintf(l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
		l.log.InfoCtx(ctx, m, l.zapTags...)
	}
}

func (l Logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= glog.Info {
		m := fmt.Sprintf(l.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
		l.log.WarnCtx(ctx, m, l.zapTags...)
	}
}

func (l Logger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= glog.Info {
		m := fmt.Sprintf(l.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
		l.log.ErrorCtx(ctx, m, l.zapTags...)
	}
}

func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel > glog.Silent {
		elapsed := time.Since(begin)
		var m string
		switch {
		case err != nil && l.LogLevel >= glog.Error:
			sql, rows := fc()
			if rows == -1 {
				m = fmt.Sprintf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				m = fmt.Sprintf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
			l.log.ErrorCtx(ctx, m, l.zapTags...)
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= glog.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
			if rows == -1 {
				m = fmt.Sprintf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				m = fmt.Sprintf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
			l.log.WarnCtx(ctx, m, l.zapTags...)
		case l.LogLevel == glog.Info:
			sql, rows := fc()
			if rows == -1 {
				m = fmt.Sprintf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				m = fmt.Sprintf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
			l.log.InfoCtx(ctx, m, l.zapTags...)
		}
	}
}
