package service

import (
	"context"

	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/domain/entity"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/domain/service"

	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/application/assembler"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/application/dto"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/application/dto/command"
)

//  application依赖domain的Port
type ApplicationService struct {
	ExperimentDomainSvr *service.ExperimentDomainService
}

//  创建新实验，实现业务逻辑
func (a *ApplicationService) NewExperiment(ctx context.Context, cmd *command.AddExperimentCmd) (*dto.ExperimentDTO, error) {
	//  前置参数校验
	if err := cmd.CheckParam() ; err != nil {
		return &dto.ExperimentDTO{}, err
	}
	// TODO 输出日志
	builder := entity.ExperimentBuilder{}
	experimentEntity := builder.Description(cmd.Description).
		Title(cmd.Title).
		ExperimentTime(cmd.ExperimentTime).
		UserID(cmd.UserID).
		ParticipantNum(cmd.ParticipantNum).Build()

	//  保存
	a.ExperimentDomainSvr.CreateNewExperiment(ctx, experimentEntity)

	//  返回DTO对象
	experimentDTO := &dto.ExperimentDTO{}
	println("%v", experimentDTO)
	assembler.AssembleExperimentDTO(experimentEntity, experimentDTO)

	return experimentDTO, nil
}
