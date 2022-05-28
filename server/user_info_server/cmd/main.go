package main

import (
	"fmt"
	userInfoPb "github.com/PsychologicalExperiment/backEnd/api/user_info_server"
	userInfo "github.com/PsychologicalExperiment/backEnd/server/user_info_server/internal/services"
	"github.com/PsychologicalExperiment/backEnd/server/user_info_server/internal/util"
	"github.com/PsychologicalExperiment/backEnd/util/plugins"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	// 初始设置
	util.InitConfig()

	// 设置grpc log
	logger := plugins.NewLogger(&util.GConfig.LoggerConfig)
	grpclog.SetLoggerV2(logger)

	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("fail to listen: %v", err)
	}

	s := grpc.NewServer()
	userInfoPb.RegisterUserServiceServer(s, &userInfo.UserInfoServerImpl{})

	reflection.Register(s)

	fmt.Printf("go backend start...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("fail to serve: %v", err)
	}
}
