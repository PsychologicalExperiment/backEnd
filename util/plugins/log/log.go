package log

import (
	"os"

	"github.com/PsychologicalExperiment/backEnd/util/plugins/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/grpclog"

	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
)

func init() {
	enCfg := zap.NewProductionEncoderConfig()
	enCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	enCfg.EncodeLevel = zapcore.CapitalLevelEncoder

	encoder := zapcore.NewJSONEncoder(enCfg)
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	zapWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename: config.Config.Server.Log,
	})

	writerCore := zapcore.NewCore(encoder, zapWriter, zap.NewAtomicLevelAt(zap.DebugLevel))
	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel)
	core := zapcore.NewTee(
		writerCore,
		consoleCore,
	)

	opts := []zap.Option{
		zap.ErrorOutput(zapWriter),
		zap.AddCaller(),
		zap.AddCallerSkip(2),
	}

	logger := zap.New(core, opts...)
	grpczap.ReplaceGrpcLoggerV2(logger)
}

func Info(args ...interface{}) {
	grpclog.Info(args...)
}

func Infof(format string, args ...interface{}) {
	grpclog.Infof(format, args...)
}

func Error(args ...interface{}) {
	grpclog.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	grpclog.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	grpclog.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	grpclog.Fatalf(format, args...)
}
