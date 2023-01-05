package impl

import (
	"context"
	"fmt"
	pb "github.com/PsychologicalExperiment/backEnd/api/eason"
	"github.com/PsychologicalExperiment/backEnd/server/eason/internal/domain"
)

type EasonNamingServiceImpl struct {
	pb.UnimplementedEasonNamingServiceServer
}

// RegisterServer 服务注册
func (s *EasonNamingServiceImpl) RegisterServer(
	ctx context.Context,
	req *pb.RegisterServerReq,
) (resp *pb.RegisterServerResp, err error) {
	r := &domain.Register{}
	uri := fmt.Sprintf("/%s/%s", req.Namespace, req.SvrName)
	if err := r.EtcdRegister(ctx, uri, req.Addr, 10); err != nil {
		resp := &pb.RegisterServerResp{
			Code: 1001,
			Msg: err.Error(),
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
	uri := fmt.Sprintf("/%s/%s", req.Namespace, req.SvrName)
	r := &domain.Register{}
	if err := r.EtcdUnRegister(ctx, uri, req.Addr); err != nil {
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
	d := &domain.Discovery{}
	addr, err := d.DiscoverServer(req.Namespace, req.SvrName)
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
		Ip:   addr,
	}
	return resp, nil
}
