package services

import (
	"context"

	userInfoPb "github.com/PsychologicalExperiment/backEnd/api/user_info_server"
	"github.com/PsychologicalExperiment/backEnd/server/user_info_server/internal/services/serverErr"
	"google.golang.org/grpc/grpclog"
)

const (
	searchKeyEmail string = "email"
)

func (u *UserInfoServerImpl) Register(
	ctx context.Context,
	req *userInfoPb.RegisterReq,
) (*userInfoPb.RegisterRsp, error) {
	err := registerParamCheck(req)
	if err != nil {
		grpclog.Errorf("param check failed, req: %+v", req)
		return serverErr.RegisterErrRsp(err), nil
	}
	isUsed, err := u.isUinqueKeyUsed(req.UserInfo.Email, searchKeyEmail)
	if err != nil {
		grpclog.Errorf("check email unique failed, req: %+v", req)
		return serverErr.RegisterErrRsp(err), nil
	}
	if isUsed {
		grpclog.Errorf("email is already used, req: %+v", req)
		return serverErr.RegisterErrRsp(serverErr.New(serverErr.ErrEmailAlreadyUsed)), nil
	}
	return nil, nil
}

func registerParamCheck(
	req *userInfoPb.RegisterReq,
) error {
	if req.UserInfo == nil {
		grpclog.Errorf("userInfo is not provided, req: %+v", req)
		return serverErr.New(serverErr.ErrUserInfoNotProvided)
	}
	return nil
}
