package assembler

import (
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/domain/entity"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/infrastructure/persistence/mysql/po"
	"time"
)

func AssembleExperimentPO(experimentEntity *entity.Experiment, experimentPO *po.ExperimentPO) {
	experimentPO.ID = experimentEntity.ID()
	experimentPO.ExperimentId = experimentEntity.ExperimentId()
	experimentPO.Title = experimentEntity.Title()
	experimentPO.Description = experimentEntity.Description()
	experimentPO.ResearcherId = experimentEntity.ResearcherId()
	experimentPO.ExperimentTime = experimentEntity.ExperimentTime()
	experimentPO.ParticipantNum = experimentEntity.ParticipantNum()
	experimentPO.State = experimentEntity.State()
}

func AssembleSubjectRecordPO(subjectRecordEntity *entity.SubjectRecord, subjectRecordPO *po.SubjectRecordPO) {
	subjectRecordPO.ID = subjectRecordEntity.ID()
	subjectRecordPO.SubjectRecordId = subjectRecordEntity.SubjectRecordId()
	subjectRecordPO.ExperimentId = subjectRecordEntity.ExperimentId()
	subjectRecordPO.ParticipantId = subjectRecordEntity.ParticipantId()
	subjectRecordPO.State = subjectRecordEntity.State()
	//t, _ := time.ParseInLocation("2006-01-02", subjectRecordEntity.TimeTaken(), time.Local)
	subjectRecordPO.FinishTime = time.Now()
}

// func AssembleExperimentEntity(experimentPO *po.ExperimentPO, subjectRecordPOList []*po.SubjectRecordPO) *entity.Experiment {
// 	var subjectRecords []*entity.SubejctRecord
// 	for _, v := range subjectRecordPOList {
// 		subjectRecordBuilder := &entity.SubejctRecordBuilder{}
// 		subjectRecordBuilder.ExperimentID(v.ExperimentID).
// 			UserID(v.UserID).
// 			SubjectRecordID(v.SubjectRecordID).
// 			State(v.State)
// 		subjectRecords = append(subjectRecords, subjectRecordBuilder.Build())
// 	}
// 	experimentBuilder := &entity.ExperimentBuilder{}
// 	experimentBuilder.Description(experimentPO.Description).
// 		ExperimentID(experimentPO.ExperimentID).
// 		Title(experimentPO.Title).
// 		InternalName(experimentPO.InternalName).
// 		Description(experimentPO.Description).
// 		UserID(experimentPO.UserID).
// 		ExperimentTime(experimentPO.ExperimentTime).
// 		ParticipantNum(experimentPO.ParticipantNum)
// 	return experimentBuilder.Build()
// }

// func AssembleSubjectRecordEntity(subjectRecordPO *po.SubjectRecordPO) *entity.SubejctRecord {
// 	subjectRecordBuilder := &entity.SubejctRecordBuilder{}
// 	subjectRecordBuilder.ExperimentID(subjectRecordPO.ExperimentID).
// 		UserID(subjectRecordPO.UserID).
// 		SubjectRecordID(subjectRecordPO.SubjectRecordID).
// 		State(subjectRecordPO.State)
// 	return subjectRecordBuilder.Build()
// }
