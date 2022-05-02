package adapter

import (
	"context"

	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/infrastructure/assembler"

	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/domain/entity"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/infrastructure/persistence/mysql/dao"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/infrastructure/persistence/mysql/po"
)

// 实现domain防腐层接口

type Experiment struct{}

func (e *Experiment) Save(ctx context.Context, experimentEntity *entity.Experiment) (err error) {
	// if experimentEntity.ID() > 0 {
	// 	err = e.updateExperiment(ctx, experimentEntity)
	// 	// TODO log
	// 	for _, subjectRecordEntity := range experimentEntity.SubjectRecords() {
	// 		if subjectRecordEntity.ID() > 0 {
	// 			_ = e.updateSubjectRecord(ctx, subjectRecordEntity)
	// 		} else {
	// 			_ = e.insertSubjectRecord(ctx, subjectRecordEntity)
	// 		}
	// 	}
	// } else {
	// 	err = e.insertExperiment(ctx, experimentEntity)
	// }
	experimentPO := &po.ExperimentPO{}
	assembler.AssembleExperimentPO(experimentEntity, experimentPO)

	experimentDAO := new(dao.ExperimentDao)
	// _ = new(dao.ExperimentDao)
	// experimentDAO.Client()
	experimentDAO.Client().Debug().Save(experimentPO)

	return nil
}

func (e *Experiment) Find(ctx context.Context, exp_id string) (*entity.Experiment, error) {
	return &entity.Experiment{}, nil
}

//  更新
// func (e *Experiment) updateExperiment(ctx context.Context, experimentEntity *entity.Experiment) (err error) {

// 	experimentPO := &po.ExperimentPO{}
// 	assembler.AssembleExperimentPO(experimentEntity, experimentPO)

// 	experimentDAO := new(dao.ExperimentDao)

// 	experimentDAO.Client().Save(experimentPO)

// }

// func (e *Experiment) insertExperiment(ctx context.Context, experimentEntity *entity.Experiment) (err error) {

// }

// func (e *Experiment) updateSubjectRecord(ctx context.Context, subjectRecordEntity *entity.SubjectRecord) (err error) {

// }

// func (e *Experiment) insertSubjectRecord(ctx context.Context, subjectRecordEntity *entity.SubjectRecord) (err error) {

// }
