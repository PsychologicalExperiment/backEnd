package main

import (
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"net"

	pb "github.com/PsychologicalExperiment/backEnd/api/experiment_server"
	applicationservice "github.com/PsychologicalExperiment/backEnd/server/experiment_server/application/service"
	domainservice "github.com/PsychologicalExperiment/backEnd/server/experiment_server/domain/service"
	infrastructureadapter "github.com/PsychologicalExperiment/backEnd/server/experiment_server/infrastructure/adapter"
	grpcinterface "github.com/PsychologicalExperiment/backEnd/server/experiment_server/user_interface/grpc"
	"github.com/PsychologicalExperiment/backEnd/util/plugins"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 51000, "The server port")
)

func main() {

	flag.Parse()
	cfg := &plugins.LoggerConfig{
		Filename: "/data/log/experiment_server.log",
	}
	log := plugins.NewLogger(cfg)
	log.Infof("test: %v", cfg)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	appService := &applicationservice.ApplicationService{
		ExperimentDomainSvr: domainservice.NewExperimentDomainService(&infrastructureadapter.Experiment{}),
	}

	grpcService := &grpcinterface.ExperimentServiceImpl{
		ApplicationService: appService,
	}

	pb.RegisterExperimentServiceServer(s, grpcService)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
