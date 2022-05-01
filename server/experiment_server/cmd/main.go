package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/PsychologicalExperiment/backEnd/api/experiment_server"

)


type server struct {
	pb.UnimplementedExperimentServiceServer
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func (s *server) QueryExperiment (ctx context.Context, in *pb.QuerySubjectRecordReq) (*pb.QuerySubjectRecordResp, error) {
	
	return &pb.QuerySubjectRecordResp{}, nil
}

func main() {

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("fialed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterExperimentServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}