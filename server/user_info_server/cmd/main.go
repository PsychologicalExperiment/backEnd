package main

import (
	"fmt"
	userInfoPb "github.com/PsychologicalExperiment/backEnd/api/user_info_server"
	userInfo "github.com/PsychologicalExperiment/backEnd/server/user_info_server/internal/services"
	"github.com/PsychologicalExperiment/backEnd/server/user_info_server/internal/util"
	"github.com/PsychologicalExperiment/backEnd/util/plugins"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net"
)

func main() {
	// 初始设置
	util.InitConfig()

	fmt.Printf("go backend start111...\n")

	// 设置grpc log
	logger := plugins.NewLogger(&util.GConfig.LoggerConfig)
	grpclog.SetLoggerV2(logger)

	// 设置数据库连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		util.GConfig.SqlConfig.User, util.GConfig.SqlConfig.Password, util.GConfig.SqlConfig.Ip, util.GConfig.SqlConfig.Port, util.GConfig.SqlConfig.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("fail to listen: %v", err)
	}
	s := grpc.NewServer()

	// server服务注册
	userInfoPb.RegisterUserServiceServer(s, userInfo.NewUserInfoServerImpl(db))

	//reflection.Register(s)

	fmt.Printf("go backend start...\n")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("fail to serve: %v", err)
	}
}
