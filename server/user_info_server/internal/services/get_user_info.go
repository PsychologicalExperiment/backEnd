package services

import (
	"context"

	userInfoPb "github.com/PsychologicalExperiment/backEnd/api/user_info_server"
	"github.com/PsychologicalExperiment/backEnd/server/user_info_server/internal/services/serverErr"
	"google.golang.org/grpc/grpclog"
)

func (u *UserInfoServerImpl) GetUserInfoBySearchKey(
	ctx context.Context,
	req *userInfoPb.GetUserInfoBySearchKeyReq,
) (*userInfoPb.GetUserInfoBySearchKeyRsp, error) {
	grpclog.Infof("req: %+v", req)
	var user []userInfo
	var err error
	if req.Email != "" {
		user, err = u.getUserInfosByKey(searchKeyEmail, req.Email)
	} else if req.PhoneNumber != "" {
		user, err = u.getUserInfosByKey(searchKeyPhoneNumber, req.PhoneNumber)
	} else if req.UserId != 0 {
		user, err = u.getUserInfosByKey(searchKeyUserId, req.UserId)
	} else {
		return &userInfoPb.GetUserInfoBySearchKeyRsp{
			CommonRsp: serverErr.CommonRsp(serverErr.New(serverErr.ErrUserInfoNotProvided)),
		}, nil
	}

	// 用户未找到
	if err != nil || len(user) == 0 {
		grpclog.Errorf("get user info failed, req: %+v", req)
		resp := &userInfoPb.GetUserInfoBySearchKeyRsp{
			CommonRsp: serverErr.CommonRsp(serverErr.New(serverErr.ErrorUserNotFound)),
		}
		return resp, nil
	}

	// 验证成功
	resp := &userInfoPb.GetUserInfoBySearchKeyRsp{
		CommonRsp: serverErr.CommonRsp(serverErr.New(serverErr.OKCode)),
		UserInfo: &userInfoPb.UserInfo{
			UserType:    userInfoPb.UserType(user[0].UserType),
			Email:       user[0].Email,
			PhoneNumber: user[0].PhoneNumber,
			UserName:    user[0].UserName,
			Gender:      userInfoPb.GenderType(user[0].Gender),
			Extra:       user[0].Extra,
			Uid:         int64(user[0].ID),
		},
	}

	return resp, nil
}
