package services

import (
	"context"
	"strconv"
	"strings"

	userInfoPb "github.com/PsychologicalExperiment/backEnd/api/user_info_server"
	"github.com/PsychologicalExperiment/backEnd/server/user_info_server/internal/services/serverErr"
	"google.golang.org/grpc/grpclog"
)

func (u *UserInfoServerImpl) Login(
	ctx context.Context,
	req *userInfoPb.LoginReq,
) (*userInfoPb.LoginRsp, error) {
	grpclog.Infof("req: %+v", req)
	var user []userInfo
	var err error
	if req.Email != "" {
		user, err = u.getUserInfosByKey(searchKeyEmail, req.Email)
	} else {
		user, err = u.getUserInfosByKey(searchKeyPhoneNumber, req.PhoneNumber)
	}
	// 用户未找到
	if err != nil || len(user) == 0 {
		grpclog.Errorf("get user info failed, req: %+v", req)
		resp := &userInfoPb.LoginRsp{
			CommonRsp: serverErr.CommonRsp(serverErr.New(serverErr.ErrorUserNotFound)),
		}
		return resp, nil
	}
	// 密码错误
	if strings.Compare(user[0].Password, req.Password) != 0 {
		grpclog.Errorf("password not right, req: %+v", req)
		resp := &userInfoPb.LoginRsp{
			CommonRsp: serverErr.CommonRsp(serverErr.New(serverErr.ErrorPasswordNotRight)),
		}
		return resp, nil
	}
	// 验证成功
	resp := &userInfoPb.LoginRsp{
		CommonRsp: serverErr.CommonRsp(serverErr.New(serverErr.OKCode)),
		UserInfo: &userInfoPb.UserInfo{UserType: userInfoPb.UserType(user[0].UserType), Email: user[0].Email, PhoneNumber: user[0].PhoneNumber,
			UserName: user[0].UserName, Gender: userInfoPb.GenderType(user[0].Gender), Extra: user[0].Extra, Uid: strconv.Itoa(int(user[0].ID))},
	}
	return resp, nil
}

func (u *UserInfoServerImpl) loginParamCheck(
	req *userInfoPb.LoginReq,
) error {
	return nil
}
