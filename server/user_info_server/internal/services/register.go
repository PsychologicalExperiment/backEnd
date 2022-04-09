package register

import (
	"context"

	commonPb "github.com/PsychologicalExperiment/backEnd/api/api_common"
	userInfoPb "github.com/PsychologicalExperiment/backEnd/api/user_info_server"
)

func Register(
	ctx context.Context,
	req *userInfoPb.RegisterReq,
) (*userInfoPb.RegisterRsp, error) {
	commonPb.CommonHead
	return nil, nil
}
