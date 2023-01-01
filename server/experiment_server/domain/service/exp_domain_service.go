package service

import (
	"context"

	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/domain/entity"
	domainport "github.com/PsychologicalExperiment/backEnd/server/experiment_server/domain/port"
	log "google.golang.org/grpc/grpclog"
)

// 领域服务
type ExperimentDomainService struct {
	//  基础层的实现
	experimentPort domainport.ExperimentPort
}

// 领域服务初始化
func NewExperimentDomainService(port domainport.ExperimentPort) *ExperimentDomainService {
	return &ExperimentDomainService{
		experimentPort: port,
	}
}

// 实现业务逻辑
func (e *ExperimentDomainService) CreateNewExperiment(
	ctx context.Context,
	experimentEntity *entity.Experiment,
) (err error) {

	//  生成一个新实验ID
	experimentEntity.GenExperimentID()

	//  插入数据库
	e.experimentPort.SaveExperiment(ctx, experimentEntity)

	return nil
}

func (e *ExperimentDomainService) UpdateExperiment(
	ctx context.Context,
	experimentEntity *entity.Experiment,
) (err error) {

	e.experimentPort.UpdateExperiment(ctx, experimentEntity)

	return nil

}

func (e *ExperimentDomainService) QueryExperiment(
	ctx context.Context,
	id string,
) (*entity.Experiment, error) {
	log.Info("queryExperiment, experiment_id: %s", id)
	experimentEntity, err := e.experimentPort.FindExperiment(ctx, id)
	if err != nil {
		return &entity.Experiment{}, err
	}

	return experimentEntity, nil
}

func (e *ExperimentDomainService) QueryExperimentList(
	ctx context.Context,
	id string,
	page int32,
	size int32,
) ([]*entity.Experiment, int32, error) {

	experimentEntityList, count, err := e.experimentPort.FindExperimentsByResearcherID(ctx, id, page, size)
	if err != nil {
		return nil, 0, err
	}

	return experimentEntityList, count, nil
}

func (e *ExperimentDomainService) CreateNewSubjectRecord(
	ctx context.Context,
	subjectRecordEntity *entity.SubjectRecord,
) (err error) {

	subjectRecordEntity.GenSubjectRecordId()

	e.experimentPort.SaveSubjectRecord(ctx, subjectRecordEntity)

	return nil
}

func (e *ExperimentDomainService) UpdateSubjectRecord(
	ctx context.Context,
	subjectRecordEntity *entity.SubjectRecord,
) (err error) {

	e.experimentPort.UpdateSubjectRecord(ctx, subjectRecordEntity)
	return nil
}

func (e *ExperimentDomainService) QuerySubjectRecord(
	ctx context.Context,
	id string,
) (subjectRecordEntity *entity.SubjectRecord, err error) {

	subjectRecordEntity, err = e.experimentPort.FindSubjectRecord(ctx, id)

	return subjectRecordEntity, nil
}

func (e *ExperimentDomainService) QuerySubjectRecordList(
	ctx context.Context,
	id string,
	page int32,
	size int32,
) (subjectRecordEntityList []*entity.SubjectRecord, err error) {

	subjectRecordEntityList, _, err = e.experimentPort.FindSubjectRecordsByExpID(ctx, id, page, size)

	return subjectRecordEntityList, nil
}
