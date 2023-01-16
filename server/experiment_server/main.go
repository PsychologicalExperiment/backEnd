package main

import (
	"fmt"
	"net"

	pb "github.com/PsychologicalExperiment/backEnd/api/experiment_server"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/internal/impl"
	"github.com/PsychologicalExperiment/backEnd/util/plugins/config"
	"github.com/PsychologicalExperiment/backEnd/util/plugins/log"
	_ "github.com/PsychologicalExperiment/backEnd/util/plugins/naming"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
			grpcprometheus.StreamServerInterceptor,
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			grpcprometheus.UnaryServerInterceptor,
		)),
	)
	conn, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.Config.Server.Port))
	if err != nil {
		log.Fatalf("tcp error: %+v", err)
	}
	impl := &impl.ExperimentServerImpl{}
	pb.RegisterExperimentServiceServer(s, impl)
	grpcprometheus.DefaultServerMetrics.InitializeMetrics(s)
	if err := s.Serve(conn); err != nil {
		log.Fatalf("start server error %+v", err)
	}
}
