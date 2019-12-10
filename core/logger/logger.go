package logger

import (
	"fmt"
	"github.com/treeyh/soc-go-common/core/config"
	"github.com/treeyh/soc-go-common/core/utils/strs"
	"os"
	"strconv"
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	_lock           sync.Mutex
	_logger         = map[string]*AppLogger{}
	_logTagKey      = "tag"
	_logDefaultName = "default"
)

var _defaultLogConfig = config.LogConfig{
	LogPath:          "",
	Level:            "info",
	FileSize:         0,
	FileNum:          0,
	IsSyncConsoleOut: true,
	Tag:              "default",
}

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

type AppLogger struct {
	logConfig *config.LogConfig

	log              *zap.Logger
	name             string
	path             string
	level            zapcore.Level
	isSyncConsoleOut bool
}

func obj2String(msg interface{}) string {
	str, ok := msg.(string)
	if !ok {
		str = strs.ObjectToString(msg)
	}
	return str
}

func (s *AppLogger) Info(msg interface{}, fields ...zap.Field) {
	s.log.Info(obj2String(msg), fields...)
}

func (s *AppLogger) Infof(fmtstr string, args ...interface{}) {
	msg := fmt.Sprintf(fmtstr, args...)
	s.log.Info(msg)
}

func (s *AppLogger) Error(msg interface{}, fields ...zap.Field) {
	s.log.Error(obj2String(msg), fields...)
}

func (s *AppLogger) Errorf(fmtstr string, args ...interface{}) {
	msg := fmt.Sprintf(fmtstr, args...)
	s.log.Error(msg)
}

func (s *AppLogger) Debug(msg interface{}, fields ...zap.Field) {
	s.log.Debug(obj2String(msg), fields...)
}

func (s *AppLogger) Debugf(fmtstr string, args ...interface{}) {
	msg := fmt.Sprintf(fmtstr, args...)
	s.log.Debug(msg)
}

func (s *AppLogger) Warn(msg interface{}, fields ...zap.Field) {
	s.log.Warn(obj2String(msg), fields...)
}

func (s *AppLogger) Warnf(fmtstr string, args ...interface{}) {
	msg := fmt.Sprintf(fmtstr, args...)
	s.log.Warn(msg)
}

func (s *AppLogger) DPanic(msg interface{}, fields ...zap.Field) {
	s.log.DPanic(obj2String(msg), fields...)
}

func (s *AppLogger) DPanicf(fmtstr string, args ...interface{}) {
	msg := fmt.Sprintf(fmtstr, args...)
	s.log.DPanic(msg)
}

func (s *AppLogger) Panic(msg interface{}, fields ...zap.Field) {
	s.log.Panic(obj2String(msg), fields...)
}

func (s *AppLogger) Panicf(fmtstr string, args ...interface{}) {
	msg := fmt.Sprintf(fmtstr, args...)
	s.log.Panic(msg)
}

func (s *AppLogger) Fatal(msg interface{}, fields ...zap.Field) {
	s.log.Fatal(obj2String(msg), fields...)
}

func (s *AppLogger) Fatalf(fmtstr string, args ...interface{}) {
	msg := fmt.Sprintf(fmtstr, args...)
	s.log.Fatal(msg)
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[strings.ToLower(lvl)]; ok {
		return level
	}
	return zapcore.InfoLevel
}

// GetDefaultLogger 获得默认Logger对象
func Logger() *AppLogger {
	logger := LoggerByName(_logDefaultName)
	if logger != nil {
		return logger
	}
	return InitLogger(_logDefaultName, &_defaultLogConfig)
}

func LoggerByName(name string) *AppLogger {
	if logger, ok := _logger[name]; ok {
		return logger
	}
	return nil
}

func InitLogger(name string, logConfig *config.LogConfig) *AppLogger {
	sysLogger := &AppLogger{
		name:      name,
		logConfig: logConfig,
	}

	sysLogger.init()
	_logger[name] = sysLogger
	return sysLogger
}

// Init 初始化Logger对象
func (s *AppLogger) init() {
	if s.log == nil {
		_lock.Lock()
		defer _lock.Unlock()
		if s.log == nil {
			s.initLogger()
		}
	}
}

func (s *AppLogger) initLogger() *AppLogger {
	logConfig := s.logConfig
	if logConfig == nil {
		panic(s.name + " log config not exist!")
	}

	level := getLoggerLevel(logConfig.Level)
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder.EncodeCaller = zapcore.FullCallerEncoder

	var core zapcore.Core

	var fileWriter zapcore.WriteSyncer
	fileType := false
	if "" != logConfig.LogPath {
		fileWriter = zapcore.AddSync(&lumberjack.Logger{
			Filename:   logConfig.LogPath,                // ⽇志⽂件路径
			MaxSize:    s.getFileSizeByConfig(logConfig), // megabytes
			MaxBackups: s.getFileNumByConfig(logConfig),  //最多保留20个备份
			LocalTime:  true,
			Compress:   true, // 是否压缩 disabled by default
		})
		fileType = true
	}

	var consoleOut zapcore.WriteSyncer
	consoleType := false

	if logConfig.IsSyncConsoleOut == true {
		// High-priority output should also go to standard error, and low-priority
		// output should also go to standard out.
		consoleOut = zapcore.Lock(os.Stdout)
		consoleType = true
	}

	if consoleType && fileType {
		core = zapcore.NewCore(zapcore.NewJSONEncoder(encoder),
			zapcore.NewMultiWriteSyncer(consoleOut, fileWriter),
			zap.NewAtomicLevelAt(level))
	} else if consoleType {
		core = zapcore.NewCore(zapcore.NewJSONEncoder(encoder),
			zapcore.NewMultiWriteSyncer(consoleOut),
			zap.NewAtomicLevelAt(level))
	} else if fileType {
		core = zapcore.NewCore(zapcore.NewJSONEncoder(encoder),
			zapcore.NewMultiWriteSyncer(fileWriter),
			zap.NewAtomicLevelAt(level))
	} else {
		panic(s.name + " log config not exist!")
	}

	// 设置初始化字段
	var logger *zap.Logger
	if "" != logConfig.Tag {
		filed := zap.Fields(zap.String(_logTagKey, logConfig.Tag))
		logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), filed)
	} else {
		logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	}

	//s.log = logger.Sugar()
	s.log = logger
	s.path = logConfig.LogPath
	s.level = level
	s.isSyncConsoleOut = logConfig.IsSyncConsoleOut
	return s
}

func (s *AppLogger) getFileSizeByConfig(logConfig *config.LogConfig) int {
	if logConfig.FileSize < 1 {
		return 1024
	}

	i, _ := strconv.Atoi(strconv.FormatInt(logConfig.FileSize, 10))
	return i
}

func (s *AppLogger) getFileNumByConfig(logConfig *config.LogConfig) int {
	if logConfig.FileNum < 1 {
		return 20
	}
	return logConfig.FileNum
}
