package impl

import (
	"context"

	pb "github.com/PsychologicalExperiment/backEnd/api/experiment_server"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/internal/entity"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/internal/errorcode"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/internal/mysql"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/internal/rpc"
	"github.com/PsychologicalExperiment/backEnd/util/plugins/log"
)

type ExperimentServerImpl struct {
}

func (s *ExperimentServerImpl) CreateExperiment(
	ctx context.Context,
	req *pb.CreateExperimentReq,
) (resp *pb.CreateExperimentResp, err error) {
	log.Infof("CreateExperiment request: %v", req)
	defer func() {
		if err != nil {
			myerr, ok := err.(errorcode.ErrorImpl)
			if !ok {
				myerr = errorcode.New(errorcode.ErrParamsTypeErrorInServer)
			}
			resp = &pb.CreateExperimentResp{
				CommonRsp: &pb.CommonRsp{
					Code: myerr.ErrorCode,
					Msg:  myerr.ErrorMsg,
				},
			}
		}
	}()
	e := &entity.ExperimentEntity{
		Title:          req.Title,
		Description:    req.Description,
		ResearcherId:   req.ResearcherId,
		ExperimentTime: req.ExperimentTime,
		ParticipantNum: req.ParticipantNum,
		State:          int32(pb.ExperimentState_EXP_PUBLISHED),
		Price:          req.Price,
		CurType:        req.CurType,
	}
	log.Infof("e: %+v", e)
	// TODO: 判断用户是否存在

	dao := &mysql.ExperimentDaoImpl{}
	id, err := dao.SaveExperiment(ctx, e)
	if err != nil {
		return nil, err
	}
	resp = &pb.CreateExperimentResp{
		CommonRsp: &pb.CommonRsp{
			Code: errorcode.OKCode,
			Msg:  "success",
		},
		ExperimentId: id,
	}
	return resp, nil
}

func (s *ExperimentServerImpl) QueryExperiment(
	ctx context.Context,
	req *pb.QueryExperimentReq,
) (resp *pb.QueryExperimentResp, err error) {
	log.Infof("QueryExperiment req: %+v", req)
	defer func() {
		if err != nil {
			myerr, ok := err.(errorcode.ErrorImpl)
			if !ok {
				myerr = errorcode.New(errorcode.ErrParamsTypeErrorInServer)
			}
			resp = &pb.QueryExperimentResp{
				CommonRsp: &pb.CommonRsp{
					Code: myerr.ErrorCode,
					Msg:  myerr.ErrorMsg,
				},
			}
		}
	}()

	// DB查數據
	dao := &mysql.ExperimentDaoImpl{}
	res, err := dao.FindExperiment(ctx, req.ExperimentId)
	if err != nil {
		return nil, err
	}
	resp = &pb.QueryExperimentResp{
		CommonRsp: &pb.CommonRsp{
			Code: errorcode.OKCode,
			Msg:  "success",
		},
		ExpInfo: &pb.ExperimentInfo{
			ExperimentId:   res.ExperimentId,
			Title:          res.Title,
			Description:    res.Description,
			ResearcherId:   res.ResearcherId,
			ExperimentTime: res.ExperimentTime,
			ParticipantNum: res.ParticipantNum,
			State:          pb.ExperimentState(res.State),
			CreateTime:     res.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateTime:     res.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	}
	return resp, err
}

func (s *ExperimentServerImpl) QueryExperimentList(
	ctx context.Context,
	req *pb.QueryExperimentListReq,
) (resp *pb.QueryExperimentListResp, err error) {
	log.Infof("QueryExperimentList req: %+v", req)
	defer func() {
		if err != nil {
			myerr, ok := err.(errorcode.ErrorImpl)
			if !ok {
				myerr = errorcode.New(errorcode.ErrParamsTypeErrorInServer)
			}
			resp = &pb.QueryExperimentListResp{
				CommonRsp: &pb.CommonRsp{
					Code: myerr.ErrorCode,
					Msg:  myerr.ErrorMsg,
				},
			}
		}
	}()

	qry := mysql.QueryExperimentReq{
		ResearcherId:  req.ResearcherId,
		Offset:        int(req.PageIndex),
		Limit:         int(req.PageSize),
		MinPrice:      req.MinPrice,
		OnlySeeMyself: req.OnlySeeMyself,
	}
	// DB查數據
	dao := &mysql.ExperimentDaoImpl{}
	res, num, err := dao.FindExperiments(ctx, qry)
	if err != nil {
		return nil, err
	}

	// 構造回包數據
	var exps []*pb.ExperimentInfo
	for _, v := range res {
		t := &pb.ExperimentInfo{
			ExperimentId:   v.ExperimentId,
			Title:          v.Title,
			Description:    v.Description,
			ResearcherId:   v.ResearcherId,
			ExperimentTime: v.ExperimentTime,
			ParticipantNum: v.ParticipantNum,
			CreateTime:     v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateTime:     v.UpdatedAt.Format("2006-01-02 15:04:05"),
			Price:          v.Price,
			CurType:        v.CurType,
		}
		exps = append(exps, t)
	}
	resp = &pb.QueryExperimentListResp{
		CommonRsp: &pb.CommonRsp{
			Code: errorcode.OKCode,
			Msg:  "success",
		},
		TotalNum:    int32(num),
		ExpInfoList: exps,
	}
	return resp, err
}

func (s *ExperimentServerImpl) UpdateExperiment(
	ctx context.Context,
	req *pb.UpdateExperimentReq,
) (resp *pb.UpdateExperimentResp, err error) {
	log.Info("UpdateExperiment: ", req)
	defer func() {
		if err != nil {
			myerr, ok := err.(errorcode.ErrorImpl)
			if !ok {
				myerr = errorcode.New(errorcode.ErrParamsTypeErrorInServer)
			}
			resp = &pb.UpdateExperimentResp{
				CommonRsp: &pb.CommonRsp{
					Code: myerr.ErrorCode,
					Msg:  myerr.ErrorMsg,
				},
			}
		}
	}()
	dao := &mysql.ExperimentDaoImpl{}
	e := &entity.ExperimentEntity{
		ExperimentId:   req.ExperimentId,
		Title:          req.Title,
		Description:    req.Description,
		ResearcherId:   req.ResearcherId,
		ExperimentTime: req.ExperimentTime,
		ParticipantNum: req.ParticipantNum,
		State:          int32(pb.ExperimentState_EXP_PUBLISHED),
		CurType:        req.CurType,
		Price:          req.Price,
		// CurType: req.CurType,
	}
	if err := dao.UpdateExperiment(ctx, e); err != nil {
		return nil, err
	}
	resp = &pb.UpdateExperimentResp{
		CommonRsp: &pb.CommonRsp{
			Code: errorcode.OKCode,
			Msg:  "success",
		},
	}
	return resp, nil
}

func (s *ExperimentServerImpl) CreateSubjectRecord(
	ctx context.Context,
	req *pb.CreateSubjectRecordReq,
) (resp *pb.CreateSubjectRecordResp, err error) {
	log.Info("CreateSubjectRecord: ", req)
	defer func() {
		if err != nil {
			myerr, ok := err.(errorcode.ErrorImpl)
			if !ok {
				myerr = errorcode.New(errorcode.ErrParamsTypeErrorInServer)
			}
			resp = &pb.CreateSubjectRecordResp{
				CommonRsp: &pb.CommonRsp{
					Code: myerr.ErrorCode,
					Msg:  myerr.ErrorMsg,
				},
			}
		}
	}()
	rcd := &entity.SubjectRecordEntity{
		ExperimentId:  req.ExperimentId,
		ParticipantId: req.ParticipantId,
	}
	dao := &mysql.ExperimentDaoImpl{}
	res, err := dao.SaveSubjectRecord(ctx, rcd)
	if err != nil {
		return nil, err
	}
	resp = &pb.CreateSubjectRecordResp{
		CommonRsp: &pb.CommonRsp{
			Code: errorcode.OKCode,
			Msg:  "success",
		},
		SubjectRecordId: res,
	}
	return resp, nil
}

func (s *ExperimentServerImpl) UpdateSubjectRecord(
	ctx context.Context,
	req *pb.UpdateSubjectRecordReq,
) (resp *pb.UpdateSubjectRecordResp, err error) {
	log.Info("UpdateExperiment: ", req)
	defer func() {
		if err != nil {
			myerr, ok := err.(errorcode.ErrorImpl)
			if !ok {
				myerr = errorcode.New(errorcode.ErrParamsTypeErrorInServer)
			}
			resp = &pb.UpdateSubjectRecordResp{
				CommonRsp: &pb.CommonRsp{
					Code: myerr.ErrorCode,
					Msg:  myerr.ErrorMsg,
				},
			}
		}
	}()
	rcd := &entity.SubjectRecordEntity{
		State: int32(req.State),
		// TODO: 完成時間
	}
	dao := &mysql.ExperimentDaoImpl{}
	if err := dao.UpdateSubjectRecord(ctx, rcd); err != nil {
		return nil, err
	}
	resp = &pb.UpdateSubjectRecordResp{
		CommonRsp: &pb.CommonRsp{
			Code: errorcode.OKCode,
			Msg:  "success",
		},
	}
	return resp, nil
}

func (s *ExperimentServerImpl) QuerySubjectRecord(
	ctx context.Context,
	req *pb.QuerySubjectRecordReq,
) (resp *pb.QuerySubjectRecordResp, err error) {
	log.Info("QuerySubjectRecord: ", req)
	defer func() {
		if err != nil {
			myerr, ok := err.(errorcode.ErrorImpl)
			if !ok {
				myerr = errorcode.New(errorcode.ErrParamsTypeErrorInServer)
			}
			resp = &pb.QuerySubjectRecordResp{
				CommonRsp: &pb.CommonRsp{
					Code: myerr.ErrorCode,
					Msg:  myerr.ErrorMsg,
				},
			}
		}
	}()
	// DB查數據
	dao := &mysql.ExperimentDaoImpl{}
	res, err := dao.FindSubjectRecord(ctx, req.SubjectRecordId)
	if err != nil {
		return nil, err
	}
	resp = &pb.QuerySubjectRecordResp{
		CommonRsp: &pb.CommonRsp{
			Code: errorcode.OKCode,
			Msg:  "success",
		},
		SubjectRecord: &pb.SubjectRecordInfo{
			SubjectRecordId: res.SubjectRecordId,
			ExperimentId:    res.ExperimentId,
			ParticipantId:   res.ParticipantId,
			// TODO: TimeTaken需要修改
			TimeTaken: res.FinishTime.Unix() - res.CreatedAt.Unix(),
			State:     pb.SubjectRecordState(res.State),
		},
	}
	return resp, err
}

func (s *ExperimentServerImpl) QuerySubjectRecordList(
	ctx context.Context,
	req *pb.QuerySubjectRecordListReq,
) (resp *pb.QuerySubjectRecordListResp, err error) {
	log.Infof("QuerySubjectRecordList req: %+v", req)
	defer func() {
		if err != nil {
			myerr, ok := err.(errorcode.ErrorImpl)
			if !ok {
				myerr = errorcode.New(errorcode.ErrParamsTypeErrorInServer)
			}
			resp = &pb.QuerySubjectRecordListResp{
				CommonRsp: &pb.CommonRsp{
					Code: myerr.ErrorCode,
					Msg:  myerr.ErrorMsg,
				},
			}
		}
	}()
	qry := mysql.QuerySubjectRecordReq{
		ExperimentId: req.ExperimentId,
		Offset:       int(req.PageIndex),
		Limit:        int(req.PageSize),
	}
	// DB查數據
	dao := &mysql.ExperimentDaoImpl{}
	res, num, err := dao.FindSubjectRecords(ctx, qry)
	if err != nil {
		return nil, err
	}
	var ids []int64
	for _, v := range res {
		ids = append(ids, v.ParticipantId)
	}
	ucli, err := rpc.NewUserInfoServerClient()
	if err != nil {
		return nil, err
	}
	users, err := ucli.BatchGetUserInfo(ctx, ids)
	if err != nil {
		return nil, err
	}
	// 構造回包數據
	var rcds []*pb.SubjectRecordInfo
	for _, v := range res {
		var user *pb.UserInfo
		for _, u := range users {
			if v.ParticipantId == u.UserId {
				user = &pb.UserInfo{
					Email:       u.Email,
					PhoneNumber: u.PhoneNumber,
					UserName:    u.UserName,
					Gender:      pb.GenderType(u.Gender),
					UserType:    pb.UserType(u.UserType),
					Extra:       u.Extra,
					Uid:         u.UserId,
				}
				break
			}
		}
		t := &pb.SubjectRecordInfo{
			ExperimentId:    v.ExperimentId,
			SubjectRecordId: v.SubjectRecordId,
			ParticipantId:   v.ParticipantId,
			TimeTaken:       v.FinishTime.Unix() - v.CreatedAt.Unix(),
			State:           pb.SubjectRecordState(v.State),
			UserInfo:        user,
		}
		rcds = append(rcds, t)
	}
	resp = &pb.QuerySubjectRecordListResp{
		CommonRsp: &pb.CommonRsp{
			Code: errorcode.OKCode,
			Msg:  "success",
		},
		TotalNum:          int32(num),
		SubjectRecordList: rcds,
	}
	log.Infof("QuerySubjectRecordList resp: %+v", resp)
	return resp, err
}
