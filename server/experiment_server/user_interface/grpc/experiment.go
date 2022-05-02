package grpc

import (
	"context"

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
func (e *ExperimentServiceImpl) NewExperiment(ctx context.Context, req *pb.NewExperimentReq) (resp *pb.NewExperimentResp, err error) {
	//  初始化dto
	cmd := &command.AddExperimentCmd{}
	resp = &pb.NewExperimentResp{}

	//  转换数据
	assembler.AssembleNewExperimentCmd(req, cmd)

	// experimentDTO, err := e.ApplicationService.NewExperiment(ctx, cmd)
	experimentDTO, err := e.ApplicationService.NewExperiment(ctx, cmd)
	//  调用service
	if err != nil {
		resp = &pb.NewExperimentResp{
			CommonRsp: &pb.CommonRsp{
				Msg: err.Error(),
				Code: 10001,
			},
		}
	}
	//  dto转换为pb
	assembler.AssembleNewExperimentResp(experimentDTO, resp) 
	// println("experimentID: %v", experimentDTO.ExperimentID)
	// resp.CommonRsp.Code = 0
	// resp.CommonRsp.Msg = "success"
	// resp.ExperimentId = experimentDTO.ExperimentID
	// println("experimentID2: ", experimentDTO.ExperimentID)
	return
}

func (e *ExperimentServiceImpl) QueryExperiment(ctx context.Context, req *pb.QueryExperimentReq) (resp *pb.QueryExperimentResp, err error) {

	err = nil
	return
}

func (e *ExperimentServiceImpl) QueryExperimentList(ctx context.Context, req *pb.QueryExperimentListReq) (resp *pb.QueryExperimentListResp, err error) {
	return
}

func (e *ExperimentServiceImpl) CreateSubjectRecord(ctx context.Context, req *pb.CreateSubjectRecordReq) (resp *pb.CreateSubjectRecordResp, err error) {
	return
}

func (e *ExperimentServiceImpl) UpdateSubjectRecord(ctx context.Context, req *pb.UpdateSubjectRecordReq) (resp *pb.UpdateSubjectRecordResp, err error) {
	return
}

func (e *ExperimentServiceImpl) QuerySubjectRecord(ctx context.Context, req *pb.QuerySubjectRecordReq) (resp *pb.QuerySubjectRecordResp, err error) {
	return
}

func (e *ExperimentServiceImpl) QuerySubjectRecordList(ctx context.Context, req *pb.QuerySubjectRecordListReq) (resp *pb.QuerySubjectRecordListResp, err error) {
	return
}
