package rpc

import (
	"context"
	"fmt"

	pb "github.com/PsychologicalExperiment/backEnd/api/user_info_server"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/internal/entity"
	"github.com/PsychologicalExperiment/backEnd/util/plugins/config"
	"github.com/PsychologicalExperiment/backEnd/util/plugins/log"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	cli pb.UserServiceClient
}

func NewUserInfoServerClient() (*client, error) {
	etcdcli, err := clientv3.NewFromURL(fmt.Sprintf("http://%s:%d",
		config.Config.NamingServer.IP, config.Config.NamingServer.Port))
	if err != nil {
		return nil, err
	}
	r, err := resolver.NewBuilder(etcdcli)
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial("etcd:///psychology/user_info_server", grpc.WithResolvers(r),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	t := pb.NewUserServiceClient(conn)
	cli := &client{
		cli: t,
	}
	return cli, nil
}

func (s *client) GetUserInfoById(
	ctx context.Context,
	id int64,
) (*entity.UserInfoEntity, error) {
	req := &pb.GetUserInfoBySearchKeyReq{
		UserId: id,
	}
	resp, err := s.cli.GetUserInfoBySearchKey(ctx, req)
	if err != nil {
		log.Errorf("GetUserInfoById err: %+v", err)
		return nil, err
	}
	if resp.CommonRsp.Code != 0 {
		log.Errorf("GetUserInfoBySearchKey resp: %+v", resp)
		return nil, err
	}
	res := &entity.UserInfoEntity{
		UserId:      resp.UserInfo.Uid,
		Email:       resp.UserInfo.Email,
		PhoneNumber: resp.UserInfo.PhoneNumber,
		UserName:    resp.UserInfo.UserName,
		Gender:      int32(resp.UserInfo.Gender),
		UserType:    int32(resp.UserInfo.UserType),
		Extra:       resp.UserInfo.Extra,
	}
	// TODO
	return res, nil
}

func (s *client) BatchGetUserInfo(
	ctx context.Context,
	ids []int64,
) ([]entity.UserInfoEntity, error) {
	req := &pb.BatchGetUserInfoReq{
		UserId: ids,
	}
	resp, err := s.cli.BatchGetUserInfos(ctx, req)
	if err != nil {
		log.Errorf("BatchGetUserInfo err: %+v", err)
		return nil, err
	}
	log.Infof("BatchGetUserInfo resp: %+v", resp)
	var res []entity.UserInfoEntity
	for _, v := range resp.UserInfo {
		userInfo := entity.UserInfoEntity{
			UserId:      v.Uid,
			Email:       v.Email,
			PhoneNumber: v.PhoneNumber,
			UserName:    v.UserName,
			Gender:      int32(v.Gender),
			UserType:    int32(v.UserType),
			Extra:       v.Extra,
		}
		res = append(res, userInfo)
	}
	return res, nil
}
