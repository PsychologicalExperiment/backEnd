package assembler

import (
	"fmt"
	commonPb "github.com/PsychologicalExperiment/backEnd/api/api_common"
	pb "github.com/PsychologicalExperiment/backEnd/api/experiment_server"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/application/dto"
	errCode "github.com/PsychologicalExperiment/backEnd/server/experiment_server/common/errorcode"
	log "google.golang.org/grpc/grpclog"
)

func AssembleCreateExperimentResp(
	experimentDTO *dto.ExperimentDTO,
	resp *pb.CreateExperimentResp,
) {
	// newExperimentResp.ExperimentId = experimentDTO.ExperimentID
	*resp = pb.CreateExperimentResp{
		CommonRsp: &commonPb.CommonRsp{
			Code: uint32(errCode.OKCode),
			Msg:  "ok",
		},
		ExperimentId: experimentDTO.ExperimentId,
	}
}

func AssembleCreateExperimentErrResp(
	err error,
	resp *pb.CreateExperimentResp,
) {
	myerr, ok := err.(errCode.ErrorImpl)
	if !ok {
		myerr = errCode.New(errCode.ErrParamsTypeErrorInServer)
	}
	*resp = pb.CreateExperimentResp{
		CommonRsp: &commonPb.CommonRsp{
			Code: uint32(myerr.ErrorCode),
			Msg:  myerr.ErrorMsg,
		},
	}
}

func AssembleUpdateExperimentResp(
	experimentDTO *dto.ExperimentDTO,
	resp *pb.UpdateExperimentResp,
) {
	*resp = pb.UpdateExperimentResp{
		CommonRsp: &commonPb.CommonRsp{
			Code: uint32(errCode.OKCode),
			Msg:  "ok",
		},
		ExperimentId: experimentDTO.ExperimentId,
	}
}

func AssembleUpdateExperimentErrResp(
	err error,
	resp *pb.UpdateExperimentResp,
) {
	myerr, ok := err.(errCode.ErrorImpl)
	if !ok {
		myerr = errCode.New(errCode.ErrParamsTypeErrorInServer)
	}
	*resp = pb.UpdateExperimentResp{
		CommonRsp: &commonPb.CommonRsp{
			Code: uint32(myerr.ErrorCode),
			Msg:  myerr.ErrorMsg,
		},
	}
}

func AssembleQueryExperimentResp(
	experimentDTO *dto.ExperimentDTO,
	resp *pb.QueryExperimentResp,
) {
	var subjectRecords []*pb.SubjectRecordInfo
	for _, v := range experimentDTO.SubjectRecords {
		tmp := &pb.SubjectRecordInfo{
			SubjectRecordId: v.SubjectRecordId,
			ExperimentId:    v.ExperimentId,
			ParticipantId:   v.ParticipantId,
			TimeTaken:       v.TimeTaken,
			State:           pb.SubjectRecordState(v.State),
		}
		subjectRecords = append(subjectRecords, tmp)
	}
	*resp = pb.QueryExperimentResp{
		CommonRsp: &commonPb.CommonRsp{
			Code: uint32(errCode.OKCode),
			Msg:  "ok",
		},
		ExpInfo: &pb.ExperimentInfo{
			ExperimentId:      experimentDTO.ExperimentId,
			Title:             experimentDTO.Title,
			Description:       experimentDTO.Description,
			ResearcherId:      experimentDTO.ResearcherId,
			ExperimentTime:    experimentDTO.ExperimentTime,
			ParticipantNum:    experimentDTO.ParticipantNum,
			State:             pb.ExperimentState(experimentDTO.State),
			CreateTime:        experimentDTO.CreateTime,
			UpdateTime:        experimentDTO.UpdateTime,
			SubjectRecordsNum: experimentDTO.SubjectRecordNum,
		},
		SubjectRecords: subjectRecords,
	}
	fmt.Println("resp: ", resp)
}

func AssembleQueryExperimentErrResp(
	err error,
	resp *pb.QueryExperimentResp,
) {
	myerr, ok := err.(errCode.ErrorImpl)
	if !ok {
		myerr = errCode.New(errCode.ErrParamsTypeErrorInServer)
	}
	*resp = pb.QueryExperimentResp{
		CommonRsp: &commonPb.CommonRsp{
			Code: uint32(myerr.ErrorCode),
			Msg:  myerr.ErrorMsg,
		},
	}
}

func AssembleQueryExperimentListResp(
	experimentDTOList []*dto.ExperimentDTO,
	resp *pb.QueryExperimentListResp,
) {
	var experiments []*pb.ExperimentInfo
	for _, v := range experimentDTOList {
		tmp := &pb.ExperimentInfo{
			ExperimentId:      v.ExperimentId,
			Title:             v.Title,
			Description:       v.Description,
			ResearcherId:      v.ResearcherId,
			ExperimentTime:    v.ExperimentTime,
			ParticipantNum:    v.ParticipantNum,
			State:             pb.ExperimentState(v.State),
			CreateTime:        v.CreateTime,
			UpdateTime:        v.UpdateTime,
			SubjectRecordsNum: v.SubjectRecordNum,
		}
		experiments = append(experiments, tmp)
	}
	*resp = pb.QueryExperimentListResp{
		CommonRsp: &commonPb.CommonRsp{
			Code: uint32(errCode.OKCode),
			Msg:  "ok",
		},
		ExpInfoList: experiments,
	}
	log.Info("QueryExperimentList resp: %v", resp)
}

func AssembleQueryExperimentListErrResp(
	err error,
	resp *pb.QueryExperimentListResp,
) {
	myerr, ok := err.(errCode.ErrorImpl)
	if !ok {
		myerr = errCode.New(errCode.ErrParamsTypeErrorInServer)
	}
	*resp = pb.QueryExperimentListResp{
		CommonRsp: &commonPb.CommonRsp{
			Code: uint32(myerr.ErrorCode),
			Msg:  myerr.ErrorMsg,
		},
	}
}

func AssembleCreateSubjectRecordResp(
	subjectRecordDTO *dto.SubjectRecordDTO,
	resp *pb.CreateSubjectRecordResp,
) {
	*resp = pb.CreateSubjectRecordResp{
		CommonRsp: &commonPb.CommonRsp{
			Code: uint32(errCode.OKCode),
			Msg:  "ok",
		},
		SubjectRecordId: subjectRecordDTO.SubjectRecordId,
	}
}

func AssembleCreateSUbjectRecordErrResp(
	err error,
	resp *pb.CreateSubjectRecordResp,
) {
	myerr, ok := err.(errCode.ErrorImpl)
	if !ok {
		myerr = errCode.New(errCode.ErrParamsTypeErrorInServer)
	}
	*resp = pb.CreateSubjectRecordResp{
		CommonRsp: &commonPb.CommonRsp{
			Code: uint32(myerr.ErrorCode),
			Msg:  myerr.ErrorMsg,
		},
	}
}

func AssembleUpdateSubjectRecordResp(
	subjectRecordDTO *dto.SubjectRecordDTO,
	resp *pb.UpdateSubjectRecordResp,
) {
	*resp = pb.UpdateSubjectRecordResp{
		CommonRsp: &commonPb.CommonRsp{
			Code: uint32(errCode.OKCode),
			Msg:  "ok",
		},
		SubjectRecordId: subjectRecordDTO.SubjectRecordId,
	}
}

func AssembleUpdateSubjectRecordErrResp(
	err error,
	resp *pb.UpdateSubjectRecordResp,
) {
	myerr, ok := err.(errCode.ErrorImpl)
	if !ok {
		myerr = errCode.New(errCode.ErrParamsTypeErrorInServer)
	}
	*resp = pb.UpdateSubjectRecordResp{
		CommonRsp: &commonPb.CommonRsp{
			Code: uint32(myerr.ErrorCode),
			Msg:  myerr.ErrorMsg,
		},
	}
}

func AssembleQuerySubjectRecordResp(
	subjectRecordDTO *dto.SubjectRecordDTO,
	resp *pb.QuerySubjectRecordResp,
) {
	*resp = pb.QuerySubjectRecordResp{
		CommonRsp: &commonPb.CommonRsp{
			Code: uint32(errCode.OKCode),
			Msg:  "ok",
		},
		SubjectRecord: &pb.SubjectRecordInfo{
			SubjectRecordId: subjectRecordDTO.SubjectRecordId,
			ExperimentId:    subjectRecordDTO.ExperimentId,
			ParticipantId:   subjectRecordDTO.ParticipantId,
			TimeTaken:       subjectRecordDTO.TimeTaken,
			State:           pb.SubjectRecordState(subjectRecordDTO.State),
		},
	}
	log.Info("QuerySubjectRecord resp: %v", resp)
}

func AssembleQuerySubjectRecordErrResp(
	err error,
	resp *pb.QuerySubjectRecordResp,
) {
	myerr, ok := err.(errCode.ErrorImpl)
	if !ok {
		myerr = errCode.New(errCode.ErrParamsTypeErrorInServer)
	}
	*resp = pb.QuerySubjectRecordResp{
		CommonRsp: &commonPb.CommonRsp{
			Code: uint32(myerr.ErrorCode),
			Msg:  myerr.ErrorMsg,
		},
	}
}

func AssembleQuerySubjectRecordListResp(
	subjectRecordDTOList []*dto.SubjectRecordDTO,
	resp *pb.QuerySubjectRecordListResp,
) {
	var subjectRecordList []*pb.SubjectRecordInfo
	for _, v := range subjectRecordDTOList {
		tmp := &pb.SubjectRecordInfo{
			SubjectRecordId: v.SubjectRecordId,
			ExperimentId:    v.ExperimentId,
			ParticipantId:   v.ParticipantId,
			TimeTaken:       v.TimeTaken,
			State:           pb.SubjectRecordState(v.State),
		}
		subjectRecordList = append(subjectRecordList, tmp)
	}
	*resp = pb.QuerySubjectRecordListResp{
		CommonRsp: &commonPb.CommonRsp{
			Code: uint32(errCode.OKCode),
			Msg:  "ok",
		},
		SubjectRecordList: subjectRecordList,
	}
}

func AssembleQuerySubjectRecordListErrResp(
	err error,
	resp *pb.QuerySubjectRecordListResp,
) {
	myerr, ok := err.(errCode.ErrorImpl)
	if !ok {
		myerr = errCode.New(errCode.ErrParamsTypeErrorInServer)
	}
	*resp = pb.QuerySubjectRecordListResp{
		CommonRsp: &commonPb.CommonRsp{
			Code: uint32(myerr.ErrorCode),
			Msg:  myerr.ErrorMsg,
		},
	}
}
