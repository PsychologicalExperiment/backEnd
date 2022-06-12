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

func (e *Experiment) SaveExperiment(
	ctx context.Context,
	experimentEntity *entity.Experiment,
) (err error) {

	experimentPO := &po.ExperimentPO{}
	assembler.AssembleExperimentPO(experimentEntity, experimentPO)

	experimentDAO := new(dao.ExperimentDao)
	// _ = new(dao.ExperimentDao)
	// experimentDAO.Client()
	experimentDAO.Client().Debug().Save(experimentPO)

	return nil
}

func (e *Experiment) SaveSubjectRecord(
	ctx context.Context,
	record *entity.SubjectRecord,
) (err error) {

	subjectRecordPO := &po.SubjectRecordPO{}
	assembler.AssembleSubjectRecordPO(record, subjectRecordPO)

	subjectRecordDAO := new(dao.SubjectRecordDAO)
	subjectRecordDAO.Client().Debug().Save(subjectRecordPO)

	return nil
}	

func (e *Experiment) UpdateExperiment(
	ctx context.Context,
	experimentEntity *entity.Experiment,
) (err error) {
	
	experimentPO := &po.ExperimentPO{}
	assembler.AssembleExperimentPO(experimentEntity, experimentPO)

	experimentDAO := new(dao.ExperimentDao)
	// _ = new(dao.ExperimentDao)
	// experimentDAO.Client()
	experimentDAO.Client().Debug().Save(experimentPO)
	experimentDAO.Client().Debug().Model(experimentPO).Where("researcher_id", experimentPO.ResearcherId).
				Omit("created_at", "researcher_id").Updates(experimentPO)
					

	return nil
}

func (e *Experiment) UpdateSubjectRecord(
	ctx context.Context,
	record *entity.SubjectRecord,
) (err error) {

	subjectRecordPO := &po.SubjectRecordPO{}
	assembler.AssembleSubjectRecordPO(record, subjectRecordPO)

	subjectRecordDAO := new(dao.SubjectRecordDAO)
	// _ = new(dao.ExperimentDao)
	// experimentDAO.Client()
	subjectRecordDAO.Client().Debug().Save(subjectRecordPO)
	subjectRecordDAO.Client().Debug().Model(subjectRecordPO).Where("participant_id", subjectRecordPO.ParticipantId).
				Omit("created_at", "researcher_id").Update("state", subjectRecordPO.State)
					

	return nil
}

func (e *Experiment) FindExperiment(
	ctx context.Context,
	exp_id string,
) (*entity.Experiment, error) {

	experimentPO := &po.ExperimentPO{}
	experimentDAO := new(dao.ExperimentDao)
	experimentDAO.Client().Debug().Where("experiment_id = ?", exp_id).Take(experimentPO)

	subjectRecordPOList := []*po.SubjectRecordPO{}
	// var subjectRecordPOList []po.SubjectRecordPO
	// subjectRecordPOList := []&po.SubjectRecordPO{}
	subjectRecordDAO := new(dao.SubjectRecordDAO)
	subjectRecordDAO.Client().Debug().Where("experiment_id = ?", exp_id).Find(&subjectRecordPOList)

	experimentEntity := assembler.AssembleExperimentEntity(experimentPO, subjectRecordPOList)

	return experimentEntity, nil
	// return &entity.Experiment{}, nil
}

func (e *Experiment) FindExperimentsByResearcherID(
	ctx context.Context,
	id string,
	page int32,
	size int32,
) ([]*entity.Experiment, int32, error) {

	experimentPOList := []*po.ExperimentPO{}
	experimentDAO := new(dao.ExperimentDao)
	experimentDAO.Client().Debug().Offset(int(page*size)).Limit(int(size)).Where("researcher_id = ?", id).Find(&experimentPOList)

	var experimentEntityList []*entity.Experiment
	for _, v := range experimentPOList {
		temp := assembler.AssembleExperimentEntity(v, nil)
		experimentEntityList = append(experimentEntityList, temp)
	}

	return experimentEntityList, int32(len(experimentPOList)), nil
}


func (e *Experiment) FindSubjectRecord(
	ctx context.Context,
	id string,
) (*entity.SubjectRecord, error) {

	subjectRecordPO := &po.SubjectRecordPO{}
	subjectRecordDAO := new(dao.SubjectRecordDAO)
	subjectRecordDAO.Client().Debug().Where("subject_record_id = ?", id).Take(subjectRecordPO)


	subjectRecordEntity := assembler.AssembleSubjectRecordEntity(subjectRecordPO)

	return subjectRecordEntity, nil

}

func (e *Experiment) FindSubjectRecordsByExpID(
	ctx context.Context, 
	id string, 
	page int32, 
	size int32,
) ([]*entity.SubjectRecord, int32, error) {

	subjectRecordPOList := []*po.SubjectRecordPO{}
	subjectRecordDAO := new(dao.SubjectRecordDAO)
	subjectRecordDAO.Client().Debug().Offset(int(page*size)).Limit(int(size)).Where("experiment_id = ?", id).Find(&subjectRecordPOList)

	var subjectRecordEntityList []*entity.SubjectRecord
	for _, v := range subjectRecordPOList {
		temp := assembler.AssembleSubjectRecordEntity(v)
		subjectRecordEntityList = append(subjectRecordEntityList, temp)
	}

	return subjectRecordEntityList, int32(len(subjectRecordEntityList)), nil
}