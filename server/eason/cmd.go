package main

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net"
	"net/http"

	pb "github.com/PsychologicalExperiment/backEnd/api/eason"
	"github.com/PsychologicalExperiment/backEnd/server/eason/impl"
	"github.com/natefinch/lumberjack"
	"google.golang.org/grpc"
	log "google.golang.org/grpc/grpclog"
)

func main() {
	enCfg := zap.NewProductionEncoderConfig()
	enCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	enCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewJSONEncoder(enCfg)
	zapcore.NewConsoleEncoder(enCfg)
	zapWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename: "/data/log/eason.log",
	})
	newCore := zapcore.NewCore(encoder, zapWriter, zap.NewAtomicLevelAt(zap.DebugLevel))
	opts := []zap.Option{zap.ErrorOutput(zapWriter)}
	opts = append(opts, zap.AddCaller(), zap.AddCallerSkip(2))
	logger := zap.New(newCore, opts...)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8001))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpc_zap.ReplaceGrpcLoggerV2(logger)
	// 设置监控
	httpServer := &http.Server{
		Handler: promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{}), Addr: fmt.Sprintf("0.0.0.0:%d", 9094),
	}
	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_recovery.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(logger),
			grpc_prometheus.StreamServerInterceptor,
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_prometheus.UnaryServerInterceptor,
		)),
	)
	log.Infof("naming_server start")
	pb.RegisterEasonNamingServiceServer(s, &impl.EasonNamingServiceImpl{})
	grpc_prometheus.DefaultServerMetrics.InitializeMetrics(s)
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal("start prometheus server error")
		}
		log.Info("start prometheus server success")
	}()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
