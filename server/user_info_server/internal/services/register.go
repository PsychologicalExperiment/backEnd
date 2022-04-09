package services

import (
	"context"

	userInfoPb "github.com/PsychologicalExperiment/backEnd/api/user_info_server"
)

func (u *UserInfoServerImpl) Register(
	ctx context.Context,
	req *userInfoPb.RegisterReq,
) (*userInfoPb.RegisterRsp, error) {

	return nil, nil
}
