package services

import (
	"context"
	"time"

	commonPb "github.com/PsychologicalExperiment/backEnd/api/api_common"
	userInfoPb "github.com/PsychologicalExperiment/backEnd/api/user_info_server"
	"github.com/PsychologicalExperiment/backEnd/server/user_info_server/internal/services/serverErr"
	"github.com/PsychologicalExperiment/backEnd/server/user_info_server/internal/util"
	"github.com/PsychologicalExperiment/backEnd/util/pkg"
	"google.golang.org/grpc/grpclog"
)

func (u *UserInfoServerImpl) Login(
	ctx context.Context,
	req *userInfoPb.LoginReq,
) (*userInfoPb.LoginRsp, error) {
	grpclog.Infof("req: %+v", req)
	var email string
	var user []userInfo
	var err error
	if req.Email != "" {
		user, err = u.getUserInfosByKey(searchKeyEmail, req.Email)
	} else {
		user, err = u.getUserInfosByKey(searchKeyPhoneNumber, req.PhoneNumber)
	}
	if err != nil {
		grpclog.Errorf("get user info failed, req: %+v", req)
	}
	token, err := pkg.GenerateUserToken(email, util.GConfig.TokenSecretKey, time.Duration(time.Hour*time.Duration(util.GConfig.TokenExpireHour)))
	rsp := &userInfoPb.LoginRsp{
		CommonRsp: &commonPb.CommonRsp{
			Code: uint32(serverErr.OKCode),
			Msg:  "ok",
		},
		Token: token,
	}
	return rsp, nil
}

func (u *UserInfoServerImpl) loginParamCheck(
	req *userInfoPb.LoginReq,
) error {
	return nil
}
