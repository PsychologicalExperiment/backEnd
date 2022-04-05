package plugins

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//LoggerConfig config of logger
type LoggerConfig struct {
	Level      string `yaml:"level"`       //debug  info  warn  error
	Encoding   string `yaml:"encoding"`    //json or console
	CallFull   bool   `yaml:"call_full"`   //whether full call path or short path, default is short
	Filename   string `yaml:"file_name"`   //log file name
	MaxSize    int    `yaml:"max_size"`    //max size of log.(MB)
	MaxAge     int    `yaml:"max_age"`     //time to keep, (day)
	MaxBackups int    `yaml:"max_backups"` //max file numbers
	LocalTime  bool   `yaml:"local_time"`  //(default UTC)
	Compress   bool   `yaml:"compress"`    //default false
}

//NewLogger create logger by config
func NewLogger(lconf *LoggerConfig) *ZapLogger {
	if lconf.Filename == "" {
		logger, _ := zap.NewProduction(zap.AddCallerSkip(2))
		return NewZapLogger(logger)
	}
	enCfg := zap.NewProductionEncoderConfig()
	enCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	enCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	if lconf.CallFull {
		enCfg.EncodeCaller = zapcore.FullCallerEncoder
	}
	encoder := zapcore.NewJSONEncoder(enCfg)
	if lconf.Encoding == "console" {
		zapcore.NewConsoleEncoder(enCfg)
	}

	//zapWriter := zapcore.
	zapWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   lconf.Filename,
		MaxSize:    lconf.MaxSize,
		MaxAge:     lconf.MaxAge,
		MaxBackups: lconf.MaxBackups,
		LocalTime:  lconf.LocalTime,
	})

	newCore := zapcore.NewCore(encoder, zapWriter, zap.NewAtomicLevelAt(convertLogLevel(lconf.Level)))
	opts := []zap.Option{zap.ErrorOutput(zapWriter)}
	opts = append(opts, zap.AddCaller(), zap.AddCallerSkip(2))
	logger := zap.New(newCore, opts...)
	return NewZapLogger(logger)
}

//NewDefaultLoggerConfig create a default config
func NewDefaultLoggerConfig() *LoggerConfig {
	return &LoggerConfig{
		Level:      "debug",
		Filename:   "./logs",
		MaxSize:    1,
		MaxAge:     1,
		MaxBackups: 10,
	}
}

func convertLogLevel(levelStr string) (level zapcore.Level) {
	switch levelStr {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	}
	return
}

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger(logger *zap.Logger) *ZapLogger {
	return &ZapLogger{
		logger: logger,
	}
}

func (zl *ZapLogger) Info(args ...interface{}) {
	zl.logger.Sugar().Info(args...)
}

func (zl *ZapLogger) Infoln(args ...interface{}) {
	zl.logger.Sugar().Info(args...)
}
func (zl *ZapLogger) Infof(format string, args ...interface{}) {
	zl.logger.Sugar().Infof(format, args...)
}

func (zl *ZapLogger) Warning(args ...interface{}) {
	zl.logger.Sugar().Warn(args...)
}

func (zl *ZapLogger) Warningln(args ...interface{}) {
	zl.logger.Sugar().Warn(args...)
}

func (zl *ZapLogger) Warningf(format string, args ...interface{}) {
	zl.logger.Sugar().Warnf(format, args...)
}

func (zl *ZapLogger) Error(args ...interface{}) {
	zl.logger.Sugar().Error(args...)
}

func (zl *ZapLogger) Errorln(args ...interface{}) {
	zl.logger.Sugar().Error(args...)
}

func (zl *ZapLogger) Errorf(format string, args ...interface{}) {
	zl.logger.Sugar().Errorf(format, args...)
}

func (zl *ZapLogger) Fatal(args ...interface{}) {
	zl.logger.Sugar().Fatal(args...)
}

func (zl *ZapLogger) Fatalln(args ...interface{}) {
	zl.logger.Sugar().Fatal(args...)
}

// Fatalf logs to fatal level
func (zl *ZapLogger) Fatalf(format string, args ...interface{}) {
	zl.logger.Sugar().Fatalf(format, args...)
}

// V reports whether verbosity level l is at least the requested verbose level.
func (zl *ZapLogger) V(v int) bool {

	return false
}
