package mysql

import (
	"context"
	"time"

	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/internal/entity"
	"github.com/PsychologicalExperiment/backEnd/util/plugins/log"
	"github.com/gofrs/uuid"
)

type ExperimentDaoImpl struct {
}

func (e *ExperimentDaoImpl) SaveExperiment(
	ctx context.Context,
	exps *entity.ExperimentEntity,
) (id string, err error) {
	tx, err := MasterClient()
	if err != nil {
		return "", err
	}
	uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	exps.ExperimentId = uid.String()
	if err := tx.Table(experimentInfoTableName).Debug().Save(exps).Error; err != nil {
		log.Errorf("SaveExperiment error: %+v", err)
		return "", err
	}
	return exps.ExperimentId, nil
}

func (e *ExperimentDaoImpl) SaveSubjectRecord(
	ctx context.Context,
	rcd *entity.SubjectRecordEntity,
) (id string, err error) {
	tx, err := MasterClient()
	if err != nil {
		return "", err
	}
	uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	rcd.SubjectRecordId = uid.String()
	if err := tx.Table(subjectRecordTableName).Debug().Save(rcd).Error; err != nil {
		log.Errorf("SaveSubjectRecord error: %+v", err)
		return "", err
	}
	return rcd.SubjectRecordId, nil
}

func (e *ExperimentDaoImpl) UpdateExperiment(
	ctx context.Context,
	exp *entity.ExperimentEntity,
) (err error) {
	tx, err := MasterClient()
	if err != nil {
		return err
	}
	if err := tx.Table(experimentInfoTableName).Debug().
		Model(&entity.ExperimentEntity{}).
		Where("researcher_id", exp.ResearcherId).
		Omit("created_at", "researcher_id").
		Updates(exp).Error; err != nil {
		log.Errorf("UpdateExperiment error: %+v", err)
		return err
	}
	return nil
}

func (e *ExperimentDaoImpl) UpdateSubjectRecord(
	ctx context.Context,
	rcd *entity.SubjectRecordEntity,
) (err error) {
	tx, err := MasterClient()
	if err != nil {
		return err
	}
	if err := tx.Table(subjectRecordTableName).
		Debug().
		Model(&entity.SubjectRecordEntity{}).
		Where("participant_id", rcd.ParticipantId).
		Omit("created_at", "researcher_id").
		Update("state", rcd.State).Error; err != nil {
		log.Errorf("UpdateSubjectRecord error: %+v", err)
		return err
	}
	return nil
}

func (e *ExperimentDaoImpl) FindExperiment(
	ctx context.Context,
	exp_id string,
) (*entity.ExperimentEntity, error) {
	tx, err := SlaveClient()
	if err != nil {
		return nil, err
	}
	var res *entity.ExperimentEntity
	if err := tx.Table(experimentInfoTableName).
		Debug().
		Where("experiment_id = ?", exp_id).
		Find(&res).Error; err != nil {
		log.Errorf("FindExperiment error: %+v", err)
		return nil, err
	}
	return res, nil
}

func (e *ExperimentDaoImpl) FindExperiments(
	ctx context.Context,
	qry QueryExperimentReq,
) ([]*entity.ExperimentEntity, int64, error) {
	tx, err := SlaveClient()
	if err != nil {
		return nil, 0, err
	}
	tx = tx.Table(experimentInfoTableName).Debug()
	if qry.MinPrice != 0 {
		tx = tx.Where("price > ?", qry.MinPrice)
	}
	if qry.ResearcherId != 0 {
		tx = tx.Where("researcher_id = ?", qry.ResearcherId)
	}
	if qry.EndTime != 0 {
		tx = tx.Where("end_time > ?", time.Unix(qry.EndTime, 0))
	}
	// TODO: 新增条件在这里加
	var cnt int64
	if err := tx.Count(&cnt).Error; err != nil {
		log.Errorf("FindExperiments get count error: %+v", err)
		return nil, 0, err
	}
	if qry.Limit == 0 {
		tx = tx.Limit(20)
	} else {
		tx = tx.Limit(qry.Limit)
	}
	if qry.Offset != 0 {
		tx = tx.Offset(qry.Offset)
	}
	var res []*entity.ExperimentEntity
	if err := tx.Find(&res).Error; err != nil {
		log.Errorf("FindExperiments error: %+v", err)
		return nil, 0, err
	}
	return res, cnt, nil
}

func (e *ExperimentDaoImpl) FindSubjectRecord(
	ctx context.Context,
	id string,
) (*entity.SubjectRecordEntity, error) {
	tx, err := SlaveClient()
	if err != nil {
		return nil, err
	}
	var res *entity.SubjectRecordEntity
	if err := tx.Table(subjectRecordTableName).
		Debug().
		Where("subject_record_id = ?", id).
		Find(res).Error; err != nil {
		log.Errorf("FindSubjectRecord error: %+v", err)
		return nil, err
	}
	return res, nil
}

func (e *ExperimentDaoImpl) FindSubjectRecords(
	ctx context.Context,
	qry QuerySubjectRecordReq,
) ([]*entity.SubjectRecordEntity, int64, error) {

	tx, err := SlaveClient()
	if err != nil {
		return nil, 0, err
	}
	tx = tx.Table(subjectRecordTableName).Debug()
	// TODO: 新增条件在这里加
	var cnt int64
	counter := tx
	if err := counter.Count(&cnt).Error; err != nil {
		log.Errorf("FindSubjectRecords get count error: %+v", err)
		return nil, 0, err
	}
	if qry.Limit == 0 {
		tx = tx.Limit(20).Offset(qry.Offset)
	} else {
		tx = tx.Limit(qry.Limit).Offset(qry.Offset)
	}
	var res []*entity.SubjectRecordEntity
	if err := tx.Find(&res).Error; err != nil {
		log.Errorf("FindSubjectRecords error: %+v", err)
		return nil, 0, err
	}
	return res, cnt, nil
}
