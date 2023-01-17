package main

import (
	"fmt"
	"net"

	userInfoPb "github.com/PsychologicalExperiment/backEnd/api/user_info_server"
	userInfo "github.com/PsychologicalExperiment/backEnd/server/user_info_server/internal/services"
	"github.com/PsychologicalExperiment/backEnd/util/plugins/config"
	_ "github.com/PsychologicalExperiment/backEnd/util/plugins/naming"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	log "google.golang.org/grpc/grpclog"
)

func main() {
	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
		)),
	)
	conn, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.Config.Server.Port))
	if err != nil {
		log.Fatalf("tcp error: %+v", err)
	}
	// server服务注册
	userInfoPb.RegisterUserServiceServer(s, userInfo.NewUserInfoServerImpl())
	grpc_prometheus.DefaultServerMetrics.InitializeMetrics(s)
	if err := s.Serve(conn); err != nil {
		log.Fatalf("fail to serve: %v", err)
	}
}
