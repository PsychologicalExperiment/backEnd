package services

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
	
	return &pb.QuerySubjectRecordResp{Code: 0, Msg: "ok"}, nil
}

func main() {

	flag.Parse()
	lis, err := net.Listen("tpc", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fataf("fialed to listen: %v", err)
	}

	s := grpc.newServer()
	pb.RegisterExperimentServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fataf("failed to serve: %v", err)
	}

}