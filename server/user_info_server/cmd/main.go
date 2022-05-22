package main

import (
	userInfoPb "github.com/PsychologicalExperiment/backEnd/api/user_info_server"
	userInfo "github.com/PsychologicalExperiment/backEnd/server/user_info_server/internal/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("fail to listen: %v", err)
	}

	s := grpc.NewServer()
	userInfoPb.RegisterUser(s, &userInfo.UserInfoServerImpl{})
}
