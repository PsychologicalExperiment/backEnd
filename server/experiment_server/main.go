package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"

	pb "github.com/PsychologicalExperiment/backEnd/api/experiment_server"
	applicationservice "github.com/PsychologicalExperiment/backEnd/server/experiment_server/application/service"
	domainservice "github.com/PsychologicalExperiment/backEnd/server/experiment_server/domain/service"
	infrastructureadapter "github.com/PsychologicalExperiment/backEnd/server/experiment_server/infrastructure/adapter"
	grpcinterface "github.com/PsychologicalExperiment/backEnd/server/experiment_server/user_interface/grpc"
	"github.com/PsychologicalExperiment/backEnd/util/etcd"
	"github.com/natefinch/lumberjack"
	"google.golang.org/grpc"
	log "google.golang.org/grpc/grpclog"
)

var (
	port = flag.Int("port", 51000, "The server port")
)

func main() {
	flag.Parse()
	enCfg := zap.NewProductionEncoderConfig()
	enCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	enCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewJSONEncoder(enCfg)
	zapcore.NewConsoleEncoder(enCfg)
	zapWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename: "/data/log/experiment_server.log",
	})
	newCore := zapcore.NewCore(encoder, zapWriter, zap.NewAtomicLevelAt(zap.DebugLevel))
	opts := []zap.Option{zap.ErrorOutput(zapWriter)}
	opts = append(opts, zap.AddCaller(), zap.AddCallerSkip(2))
	logger := zap.New(newCore, opts...)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpczap.ReplaceGrpcLoggerV2(logger)
	// 设置监控
	httpServer := &http.Server{
		Handler: promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{}), Addr: fmt.Sprintf("0.0.0.0:%d", 9092),
	}
	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_recovery.StreamServerInterceptor(),
			grpczap.StreamServerInterceptor(logger),
			grpcprometheus.StreamServerInterceptor,
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
			grpczap.UnaryServerInterceptor(logger),
			grpcprometheus.UnaryServerInterceptor,
		)),
	)
	log.Infof("server start")
	nconn, err := grpc.Dial("159.75.15.177:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer func() {
		if err := nconn.Close(); err != nil {
			log.Error(err)
		}
	}()
	if err != nil {
		log.Error("naming-service error: ", err)
	}
	//namingcli := namingserver.NewEasonNamingServiceClient(nconn)
	ip, err := etcd.GetLocalIP()
	if err != nil {
		log.Fatal(err)
	}
	if err := etcd.EtcdRegisterServer(context.Background(),
			"experiment_server", fmt.Sprintf("%s:%d", ip, port), 10); err != nil {
		log.Errorf("register server error: %+v", err)
		log.Fatal(err)
	}
	//req := &namingserver.RegisterServerReq{
	//	Namespace: "etcd",
	//	SvrName:   "experiment_server",
	//	Addr:      fmt.Sprintf("%s:%d", ip, port),
	//}
	//resp, err := namingcli.RegisterServer(context.Background(), req)
	//if err != nil {
	//	log.Error("register server error: ", err)
	//}
	//if resp.Code != 0 {
	//	log.Error("register server error: ", resp.Msg)
	//}
	appService := &applicationservice.ApplicationService{
		ExperimentDomainSvr: domainservice.NewExperimentDomainService(&infrastructureadapter.Experiment{}),
	}

	grpcService := &grpcinterface.ExperimentServiceImpl{
		ApplicationService: appService,
	}
	pb.RegisterExperimentServiceServer(s, grpcService)
	grpcprometheus.DefaultServerMetrics.InitializeMetrics(s)
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
