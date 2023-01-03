package impl

import (
	"context"
	"fmt"
	pb "github.com/PsychologicalExperiment/backEnd/api/eason"
	"github.com/PsychologicalExperiment/backEnd/server/eason/etcd"
	"google.golang.org/grpc"
)

type EasonNamingServiceImpl struct {
	pb.UnimplementedEasonNamingServiceServer
}

func (s *EasonNamingServiceImpl) RegisterServer(
	ctx context.Context,
	req *pb.RegisterServerReq,
) (resp *pb.RegisterServerResp, err error) {
	if err := etcd.Register(req.Namespace, req.SvrName, req.Addr); err != nil {
		resp := &pb.RegisterServerResp{
			Code: 1001,
			Msg:  err.Error(),
		}
		return resp, nil
	}
	resp = &pb.RegisterServerResp{
		Code: 0,
		Msg:  "success",
	}
	return resp, nil
}

func (s *EasonNamingServiceImpl) UnRegisterServer(
	ctx context.Context,
	req *pb.UnRegisterServerReq,
) (resp *pb.UnRegisterServerResp, err error) {
	etcdKey := fmt.Sprintf("/%s/%s/%s", req.Namespace, req.SvrName, req.Addr)
	if err := etcd.UnRegister(etcdKey); err != nil {
		resp := &pb.UnRegisterServerResp{
			Code: 1002,
			Msg:  err.Error(),
		}
		return resp, nil
	}
	resp = &pb.UnRegisterServerResp{
		Code: 0,
		Msg:  "success",
	}
	return resp, nil
}

func (s *EasonNamingServiceImpl) DiscoverServer(
	ctx context.Context,
	req *pb.DiscoverServerReq,
) (resp *pb.DiscoverServerResp, err error) {
	r := etcd.NewResolver()
	conn, err := grpc.Dial(fmt.Sprintf("etcd:///%s/%s", req.Namespace, req.SvrName), grpc.WithResolvers(r))
	if err != nil {
		resp := &pb.DiscoverServerResp{
			Code: 1003,
			Msg:  err.Error(),
		}
		return resp, nil
	}
	resp = &pb.DiscoverServerResp{
		Code: 0,
		Msg:  "success",
		Ip:   conn.Target(),
	}
	return resp, nil
}
