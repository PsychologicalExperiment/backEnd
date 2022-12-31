package main

import (
	"fmt"
	userInfoPb "github.com/PsychologicalExperiment/backEnd/api/user_info_server"
	userInfo "github.com/PsychologicalExperiment/backEnd/server/user_info_server/internal/services"
	"github.com/PsychologicalExperiment/backEnd/server/user_info_server/internal/util"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/natefinch/lumberjack"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	log "google.golang.org/grpc/grpclog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net"
	"net/http"
)

func main() {
	// 初始设置
	util.InitConfig()

	fmt.Printf("go backend start111...\n")

	// 设置grpc log
	enCfg := zap.NewProductionEncoderConfig()
	enCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	enCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewJSONEncoder(enCfg)
	zapWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename: "/data/log/user_info.log",
	})
	newCore := zapcore.NewCore(encoder, zapWriter, zap.NewAtomicLevelAt(zap.DebugLevel))
	opts := []zap.Option{zap.ErrorOutput(zapWriter)}
	opts = append(opts, zap.AddCaller(), zap.AddCallerSkip(2))
	logger := zap.New(newCore, opts...)
	grpc_zap.ReplaceGrpcLoggerV2(logger)

	// 设置监控
	httpServer := &http.Server{
		Handler: promhttp.HandlerFor(
			prometheus.DefaultGatherer,
			promhttp.HandlerOpts{}),
		Addr: fmt.Sprintf("0.0.0.0:%d", 9093),
	}
	// 设置数据库连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		util.GConfig.SqlConfig.User, util.GConfig.SqlConfig.Password, util.GConfig.SqlConfig.Ip, util.GConfig.SqlConfig.Port, util.GConfig.SqlConfig.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("fail to listen: %v", err)
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

	// server服务注册
	userInfoPb.RegisterUserServiceServer(s, userInfo.NewUserInfoServerImpl(db))
	grpc_prometheus.DefaultServerMetrics.InitializeMetrics(s)
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal("start prometheus server error")
		}
		log.Info("start prometheus server success")
	}()
	//reflection.Register(s)

	fmt.Printf("go backend start...\n")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("fail to serve: %v", err)
	}
}
