package main

import (
	"fmt"
	"net"

	pb "github.com/PsychologicalExperiment/backEnd/api/experiment_server"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/internal/impl"
	"github.com/PsychologicalExperiment/backEnd/util/plugins/config"
	"github.com/PsychologicalExperiment/backEnd/util/plugins/log"
	_ "github.com/PsychologicalExperiment/backEnd/util/plugins/mon"
	_ "github.com/PsychologicalExperiment/backEnd/util/plugins/naming"
	_ "github.com/PsychologicalExperiment/backEnd/util/plugins/recovery"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer(
		grpc.StreamInterceptor(middleware.ChainStreamServer(
			validator.StreamServerInterceptor(),
			recovery.StreamServerInterceptor(),
			prometheus.StreamServerInterceptor,
		)),
		grpc.UnaryInterceptor(middleware.ChainUnaryServer(
			validator.UnaryServerInterceptor(),
			prometheus.UnaryServerInterceptor,
		)),
	)
	conn, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.Config.Server.Port))
	if err != nil {
		log.Fatalf("tcp error: %+v", err)
	}
	expImpl := &impl.ExperimentServerImpl{}
	pb.RegisterExperimentServiceServer(s, expImpl)
	prometheus.DefaultServerMetrics.InitializeMetrics(s)
	if err := s.Serve(conn); err != nil {
		log.Fatalf("start server error %+v", err)
	}
}
