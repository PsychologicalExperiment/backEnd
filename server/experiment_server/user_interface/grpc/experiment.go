package grpc

import (
	"context"
	"fmt"

	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/application/dto/query"

	pb "github.com/PsychologicalExperiment/backEnd/api/experiment_server"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/application/dto/command"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/application/service"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/user_interface/grpc/assembler"
)

type ExperimentServiceImpl struct {
	ApplicationService *service.ApplicationService
	pb.UnimplementedExperimentServiceServer
}

//  用户接口层用于将pb转换成DTO，然后调用application的service
func (e *ExperimentServiceImpl) CreateExperiment(
	ctx context.Context,
	req *pb.CreateExperimentReq,
) (resp *pb.CreateExperimentResp, err error) {
	//  初始化dto
	cmd := &command.AddExperimentCmd{}
	//  转换数据
	assembler.AssembleAddExperimentCmd(req, cmd)

	resp = &pb.CreateExperimentResp{}
	// experimentDTO, err := e.ApplicationService.NewExperiment(ctx, cmd)
	experimentDTO, err := e.ApplicationService.CreateExperiment(ctx, cmd)
	//  调用service
	if err != nil {
		assembler.AssembleCreateExperimentErrResp(err, resp)
	}
	//  dto转换为pb
	assembler.AssembleCreateExperimentResp(experimentDTO, resp)
	return
}

func (e *ExperimentServiceImpl) QueryExperiment(
	ctx context.Context,
	req *pb.QueryExperimentReq,
) (resp *pb.QueryExperimentResp, err error) {

	query := &query.GetExperimentQry{}
	assembler.AssembleGetExperimentQry(req, query)

	experimentDTO, err := e.ApplicationService.QueryExperiment(ctx, query)

	resp = &pb.QueryExperimentResp{}
	if err != nil {
		assembler.AssembleQueryExperimentErrResp(err, resp)
	}

	fmt.Println("experimentDTO: ", experimentDTO)
	
	assembler.AssembleQueryExperimentResp(experimentDTO, resp)
	fmt.Println(resp)
	return resp, err
}

func (e *ExperimentServiceImpl) QueryExperimentList(
	ctx context.Context,
	req *pb.QueryExperimentListReq,
) (resp *pb.QueryExperimentListResp, err error) {

	query := &query.GetExperimentListQry{}
	assembler.AssembleGetExperimentListQry(req, query)

	experimentDTOList, err := e.ApplicationService.QueryExperimentList(ctx, query)

	resp = &pb.QueryExperimentListResp{}
	assembler.AssembleQueryExperimentListResp(experimentDTOList, resp)
	return
}

func (e *ExperimentServiceImpl) UpdateExperiment(
	ctx context.Context,
	req *pb.UpdateExperimentReq,
) (resp *pb.UpdateExperimentResp, err error) {

	cmd := &command.UpdateExperimentCmd{}
	assembler.AssembleUpdateExperimentCmd(req, cmd)

	experimentDTO, err := e.ApplicationService.UpdateExperiment(ctx, cmd)

	resp = &pb.UpdateExperimentResp{}
	assembler.AssembleUpdateExperimentResp(experimentDTO, resp)

	return 
}

func (e *ExperimentServiceImpl) CreateSubjectRecord(
	ctx context.Context,
	req *pb.CreateSubjectRecordReq,
) (resp *pb.CreateSubjectRecordResp, err error) {

	cmd := &command.AddSubjectRecordCmd{}
	assembler.AssembleAddSubjectRecordCmd(req, cmd)

	subjectRecordDTO, err := e.ApplicationService.CreateSubjectRecord(ctx, cmd)

	resp = &pb.CreateSubjectRecordResp{}
	println("DTO: %v", subjectRecordDTO)
	assembler.AssembleCreateSubjectRecordResp(subjectRecordDTO, resp)

	return
}

func (e *ExperimentServiceImpl) UpdateSubjectRecord(
	ctx context.Context,
	req *pb.UpdateSubjectRecordReq,
) (resp *pb.UpdateSubjectRecordResp, err error) {
	
	cmd := &command.UpdateSubjectRecordCmd{}
	assembler.AssembleUpdateSubjectRecordCmd(req, cmd)

	subjectRecordDTO, err := e.ApplicationService.UpdateSubjectRecord(ctx, cmd)

	resp = &pb.UpdateSubjectRecordResp{}
	assembler.AssembleUpdateSubjectRecordResp(subjectRecordDTO, resp)

	return 
}

func (e *ExperimentServiceImpl) QuerySubjectRecord(
	ctx context.Context,
	req *pb.QuerySubjectRecordReq,
) (resp *pb.QuerySubjectRecordResp, err error) {
	
	query := &query.GetSubjectRecordQry{}
	assembler.AssembleGetSubjectRecordQry(req, query)

	subjectRecordDTO, err := e.ApplicationService.QuerySubjectRecord(ctx, query)

	resp = &pb.QuerySubjectRecordResp{}
	assembler.AssembleQuerySubjectRecordResp(subjectRecordDTO, resp)
	
	return
}

func (e *ExperimentServiceImpl) QuerySubjectRecordList(
	ctx context.Context,
	req *pb.QuerySubjectRecordListReq,
) (resp *pb.QuerySubjectRecordListResp, err error) {

	query := &query.GetSubjectRecordListQry{}
	assembler.AssembleGetSubjectRecordListQry(req, query)

	subjectRecordDTOList, err := e.ApplicationService.QuerySubjectRecordList(ctx, query)

	resp = &pb.QuerySubjectRecordListResp{}
	assembler.AssembleQuerySubjectRecordListResp(subjectRecordDTOList, resp)

	return 
}