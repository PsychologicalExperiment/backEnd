package service

import (
	"context"

	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/domain/entity"
	domainport "github.com/PsychologicalExperiment/backEnd/server/experiment_server/domain/port"
)

//  领域服务
type ExperimentDomainService struct {
	//  基础层的实现
	experimentPort domainport.ExperimentPort
}

//  领域服务初始化
func NewExperimentDomainService(port domainport.ExperimentPort) *ExperimentDomainService {
	return &ExperimentDomainService{
		experimentPort: port,
	}
}

//  实现业务逻辑
func (e *ExperimentDomainService) CreateNewExperiment(ctx context.Context, experimentEntity *entity.Experiment) (err error) {

	//  生成一个新实验ID
	experimentEntity.GenExperimentID()

	//  插入数据库
	e.experimentPort.Save(ctx, experimentEntity)

	return nil
}
