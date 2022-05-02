package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/PsychologicalExperiment/backEnd/api/experiment_server"
	applicationservice "github.com/PsychologicalExperiment/backEnd/server/experiment_server/application/service"
	domainservice "github.com/PsychologicalExperiment/backEnd/server/experiment_server/domain/service"
	infrastructureadapter "github.com/PsychologicalExperiment/backEnd/server/experiment_server/infrastructure/adapter"
	grpcinterface "github.com/PsychologicalExperiment/backEnd/server/experiment_server/user_interface/grpc"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

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
