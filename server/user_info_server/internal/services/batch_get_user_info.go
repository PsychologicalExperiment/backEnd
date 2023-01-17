package services

import (
	"context"

	pb "github.com/PsychologicalExperiment/backEnd/api/user_info_server"
	"github.com/PsychologicalExperiment/backEnd/server/user_info_server/internal/services/serverErr"
	"github.com/PsychologicalExperiment/backEnd/util/plugins/log"
)

func (u *UserInfoServerImpl) BatchGetUserInfos(
	ctx context.Context,
	req *pb.BatchGetUserInfoReq,
) (*pb.BatchGetUserInfoRsp, error) {
	log.Infof("GetBatchUserInfo req: %+v", req)
	userInfos, err := u.batchGetUserInfo(ctx, req.UserId)
	if err != nil {
		log.Errorf("batch get user info error: %+v", err)
		resp := &pb.BatchGetUserInfoRsp{
			CommonRsp: serverErr.CommonRsp(err),
		}
		return resp, nil
	}
	var pbUserInfo []*pb.UserInfo
	for _, info := range userInfos {
		i := &pb.UserInfo{
			Email:       info.Email,
			PhoneNumber: info.PhoneNumber,
			UserName:    info.UserName,
			Gender:      pb.GenderType(info.Gender),
			UserType:    pb.UserType(info.UserType),
			Extra:       info.Extra,
			Uid:         int64(info.ID),
		}
		pbUserInfo = append(pbUserInfo, i)
	}
	resp := &pb.BatchGetUserInfoRsp{
		CommonRsp: serverErr.CommonRsp(serverErr.New(serverErr.OKCode)),
		UserInfo:  pbUserInfo,
	}
	return resp, nil
}
