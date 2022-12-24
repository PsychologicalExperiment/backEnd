package main

import (
	userInfoPb "github.com/PsychologicalExperiment/backEnd/api/user_info_server"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	userInfoPb.RegisterUserServiceServer()
}
