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

func (u *UserInfoServerImpl) Register(
	ctx context.Context,
	req *userInfoPb.RegisterReq,
) (*userInfoPb.RegisterRsp, error) {
	grpclog.Infof("req: %+v", req)
	err := u.registerParamCheck(req)
	if err != nil {
		grpclog.Errorf("param check failed, req: %+v", req)
		return serverErr.RegisterErrRsp(err), nil
	}
	token, err := pkg.GenerateUserToken(req.UserInfo.Email, util.GConfig.TokenSecretKey, time.Duration(time.Hour*time.Duration(util.GConfig.TokenExpireHour)))
	if err != nil {
		grpclog.Errorf("generate token failed, error: %+v, req: %+v", err, req)
		return serverErr.RegisterErrRsp(serverErr.New(serverErr.ErrGenerateTokenFailed)), nil
	}
	err = u.setTokenForUser(ctx, req.UserInfo.Email, token)
	if err != nil {
		grpclog.Errorf("set token failed, req: %+v", req)
		return serverErr.RegisterErrRsp(err), nil
	}
	err = u.insertUserInfo(&userInfo{
		Email:       req.UserInfo.Email,
		PhoneNumber: req.UserInfo.PhoneNumber,
		UserName:    req.UserInfo.UserName,
		Gender:      req.UserInfo.Gender,
		Password:    req.UserInfo.Password,
		UserType:    req.UserInfo.UserType,
	})
	if err != nil {
		grpclog.Errorf("inser user info failed, req: %+v", req)
		return serverErr.RegisterErrRsp(err), nil
	}
	rsp := &userInfoPb.RegisterRsp{
		CommonRsp: &commonPb.CommonRsp{
			Code: uint32(serverErr.OKCode),
			Msg:  "ok",
		},
		Token: token,
	}
	return rsp, nil
}

func (u *UserInfoServerImpl) registerParamCheck(
	req *userInfoPb.RegisterReq,
) error {
	if req.UserInfo == nil {
		grpclog.Errorf("userInfo is not provided, req: %+v", req)
		return serverErr.New(serverErr.ErrUserInfoNotProvided)
	}
	if req.UserInfo.Email == "" {
		grpclog.Errorf("email is not provided, req: %+v", req)
		return serverErr.New(serverErr.ErrEmailNotProvided)
	}
	if req.UserInfo.Password == "" {
		grpclog.Errorf("password is not provided, req: %+v", req)
		return serverErr.New(serverErr.ErrPasswordNotProvided)
	}
	isUsed, err := u.isUinqueKeyUsed(req.UserInfo.Email, searchKeyEmail)
	if err != nil {
		grpclog.Errorf("check email unique failed, req: %+v", req)
		return err
	}
	if isUsed {
		grpclog.Errorf("email is already used, req: %+v", req)
		return serverErr.New(serverErr.ErrEmailAlreadyUsed)
	}
	if req.UserInfo.PhoneNumber != "" {
		isUsed, err := u.isUinqueKeyUsed(req.UserInfo.PhoneNumber, searchKeyPhoneNumber)
		if err != nil {
			grpclog.Errorf("check phone unique failed, req: %+v", req)
			return err
		}
		if isUsed {
			grpclog.Errorf("phone number is already used, req: %+v", req)
			return serverErr.New(serverErr.ErrPhoneNumberAlreadyUsed)
		}
	}
	if req.UserInfo.Gender == uint32(userInfoPb.GenderType_GENDER_TYPE_INVALID) {
		grpclog.Errorf("gender invalid, req: %+v", req)
		return serverErr.New(serverErr.ErrGenderInvalid)
	}
	return nil
}
