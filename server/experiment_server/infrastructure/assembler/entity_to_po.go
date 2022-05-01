package assembler

import (
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/domain/entity"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/infrastructure/persistence/mysql/po"
)

func AssembleExperimentEntity(experimentPO *po.ExperimentPO, subjectRecordPOList []*po.SubjectRecordPO) *entity.Experiment {
	var subjectRecords []*entity.SubejctRecord
	for _, v := range subjectRecordPOList {
		subjectRecordBuilder := &entity.SubejctRecordBuilder{}
		subjectRecordBuilder.ExperimentID(v.ExperimentID).
			UserID(v.UserID).
			SubjectRecordID(v.SubjectRecordID).
			State(v.State)
		subjectRecords = append(subjectRecords, subjectRecordBuilder.Build())
	}
	experimentBuilder := &entity.ExperimentBuilder{}
	experimentBuilder.Description(experimentPO.Description).
		ExperimentID(experimentPO.ExperimentID).
		Title(experimentPO.Title).
		InternalName(experimentPO.InternalName).
		Description(experimentPO.Description).
		UserID(experimentPO.UserID).
		ExperimentTime(experimentPO.ExperimentTime).
		ParticipantNum(experimentPO.ParticipantNum)
	return experimentBuilder.Build()
}

func AssembleSubjectRecordEntity(subjectRecordPO *po.SubjectRecordPO) *entity.SubejctRecord {
	subjectRecordBuilder := &entity.SubejctRecordBuilder{}
	subjectRecordBuilder.ExperimentID(subjectRecordPO.ExperimentID).
		UserID(subjectRecordPO.UserID).
		SubjectRecordID(subjectRecordPO.SubjectRecordID).
		State(subjectRecordPO.State)
	return subjectRecordBuilder.Build()
}
