package service

import (
	"fmt"
	"context"

	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/domain/entity"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/domain/service"

	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/application/assembler"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/application/dto"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/application/dto/command"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/application/dto/query"
	errCode "github.com/PsychologicalExperiment/backEnd/server/experiment_server/common/errorcode"
)

//  application依赖domain的Port
type ApplicationService struct {
	ExperimentDomainSvr *service.ExperimentDomainService
}

//  创建新实验，实现业务逻辑
func (a *ApplicationService) CreateExperiment(
	ctx context.Context,
	cmd *command.AddExperimentCmd,
) (*dto.ExperimentDTO, error) {
	//  前置参数校验
	if err := cmd.CheckParam(); err != nil {
		return &dto.ExperimentDTO{}, errCode.New(errCode.ErrParamsInvalid)
	}
	// TODO 输出日志
	builder := entity.ExperimentBuilder{}
	experimentEntity := builder.Description(cmd.Description).
		Title(cmd.Title).
		ExperimentTime(cmd.ExperimentTime).
		ResearcherId(cmd.ResearcherId).
		ParticipantNum(cmd.ParticipantNum).Build()

	//  保存
	a.ExperimentDomainSvr.CreateNewExperiment(ctx, experimentEntity)

	//  返回DTO对象
	experimentDTO := &dto.ExperimentDTO{}
	println("%v", experimentDTO)
	assembler.AssembleExperimentDTO(experimentEntity, experimentDTO)

	return experimentDTO, nil
}

func (a *ApplicationService) QueryExperiment(
	ctx context.Context,
	query *query.GetExperimentQry,
) (*dto.ExperimentDTO, error) {

	if err := query.CheckParam(); err != nil {
		return &dto.ExperimentDTO{}, errCode.New(errCode.ErrParamsInvalid)
	}

	experimentEntity, err := a.ExperimentDomainSvr.QueryExperiment(ctx, query.ExperimentId)
	if err != nil {
		return &dto.ExperimentDTO{}, errCode.New(errCode.ErrQueryRecordNotFound)
	}
	experimentDTO := &dto.ExperimentDTO{}
	assembler.AssembleExperimentDTO(experimentEntity, experimentDTO)
	return experimentDTO, nil
}

func (a *ApplicationService) QueryExperimentList(
	ctx context.Context,
	query *query.GetExperimentListQry,
) ([]*dto.ExperimentDTO, error) {

	if err := query.CheckParam(); err != nil {
		return nil, errCode.New(errCode.ErrParamsInvalid)
	}

	experimentEntityList, _, err := a.ExperimentDomainSvr.QueryExperimentList(ctx, query.ResearcherId,
						query.PageIndex, query.PageSize)
	if err != nil {
		return nil, errCode.New(errCode.ErrQueryRecordNotFound)
	}
	var experimentDTOList []*dto.ExperimentDTO
	for _, v := range experimentEntityList {
		experimentDTO := &dto.ExperimentDTO{}
		assembler.AssembleExperimentDTO(v, experimentDTO)
		experimentDTOList = append(experimentDTOList, experimentDTO)
	}
	return experimentDTOList, nil
}

func (a *ApplicationService) UpdateExperiment(
	ctx context.Context,
	cmd *command.UpdateExperimentCmd,
) (*dto.ExperimentDTO, error) {

	if err := cmd.CheckParam(); err != nil {
		return &dto.ExperimentDTO{}, errCode.New(errCode.ErrParamsInvalid)
	}

	builder := entity.ExperimentBuilder{}
	experimentEntity := builder.Description(cmd.Description).
		ExperimentID(cmd.ExperimentId).
		Title(cmd.Title).
		ExperimentTime(cmd.ExperimentTime).
		ResearcherId(cmd.ResearcherId).
		State(cmd.State).
		ParticipantNum(cmd.ParticipantNum).Build()

	a.ExperimentDomainSvr.UpdateExperiment(ctx, experimentEntity)

	experimentDTO := &dto.ExperimentDTO{}
	println("%v", experimentDTO)
	assembler.AssembleExperimentDTO(experimentEntity, experimentDTO)
	fmt.Println("experimentDTO:%v", experimentDTO)
	return experimentDTO, nil
}

func (a *ApplicationService) CreateSubjectRecord(
	ctx context.Context,
	cmd *command.AddSubjectRecordCmd,
) (*dto.SubjectRecordDTO, error) {

	if err := cmd.CheckParam(); err != nil {
		return &dto.SubjectRecordDTO{}, errCode.New(errCode.ErrParamsInvalid)
	}

	builder := entity.SubjectRecordBuilder{}
	subjectRecordEntity := builder.ExperimentID(cmd.ExperimentId).
						ParticipantId(cmd.ParticipantId).
						Build()

	a.ExperimentDomainSvr.CreateNewSubjectRecord(ctx, subjectRecordEntity)

	subjectRecordDTO := &dto.SubjectRecordDTO{}
	println("subjectRecordEntity:%v", subjectRecordEntity)
	assembler.AssembleSubjectRecordDTO(subjectRecordEntity, subjectRecordDTO)		
	println("experimentDTO:%v", subjectRecordDTO)
	return subjectRecordDTO, nil
}

func (a *ApplicationService) UpdateSubjectRecord(
	ctx context.Context,
	cmd *command.UpdateSubjectRecordCmd,
) (*dto.SubjectRecordDTO, error) {
	if err := cmd.CheckParam(); err != nil {
		return nil, errCode.New(errCode.ErrParamsInvalid)
	}
	//   判断更新状态逻辑

	builder := entity.SubjectRecordBuilder{}
	subjectRecordEntity := builder.SubjectRecordID(cmd.SubjectRecordId).
							State(cmd.State).Build()
	
	a.ExperimentDomainSvr.UpdateSubjectRecord(ctx, subjectRecordEntity)

	subjectRecordDTO := &dto.SubjectRecordDTO{}
	assembler.AssembleSubjectRecordDTO(subjectRecordEntity, subjectRecordDTO)

	return subjectRecordDTO, nil
}

func (a *ApplicationService) QuerySubjectRecord(
	ctx context.Context,
	query *query.GetSubjectRecordQry,
) (*dto.SubjectRecordDTO, error) {

	if err := query.CheckParam(); err != nil {
		return nil, errCode.New(errCode.ErrParamsInvalid)
	}

	subjectRecordEntity, _ := a.ExperimentDomainSvr.QuerySubjectRecord(ctx, query.SubjectRecordId)
	subjectRecordDTO := &dto.SubjectRecordDTO{}
	assembler.AssembleSubjectRecordDTO(subjectRecordEntity, subjectRecordDTO)

	return subjectRecordDTO, nil
}

//  
func (a *ApplicationService) QuerySubjectRecordList(
	ctx context.Context,
	query *query.GetSubjectRecordListQry,
) ([]*dto.SubjectRecordDTO, error) {

	if err := query.CheckParam(); err != nil {
		return nil, errCode.New(errCode.ErrParamsInvalid)
	}

	var subjectRecordDTOList []*dto.SubjectRecordDTO
	subjectRecordEntityList, _ := a.ExperimentDomainSvr.QuerySubjectRecordList(ctx, query.ExperimentId, query.PageIndex, query.PageSize)

	for _, v := range subjectRecordEntityList {
		temp := &dto.SubjectRecordDTO{}
		assembler.AssembleSubjectRecordDTO(v, temp)
		subjectRecordDTOList = append(subjectRecordDTOList, temp)
	}
	return subjectRecordDTOList, nil
}
