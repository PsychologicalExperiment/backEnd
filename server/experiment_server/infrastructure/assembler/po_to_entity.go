package assembler

import (
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/domain/entity"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/infrastructure/persistence/mysql/po"
)

func AssembleExperimentEntity(
	experimentPO *po.ExperimentPO, 
	subjectRecordPOList []*po.SubjectRecordPO,
) *entity.Experiment {
	var subjectRecords []*entity.SubjectRecord
	for _, v := range subjectRecordPOList {
		subjectRecordBuilder := &entity.SubjectRecordBuilder{}
		subjectRecordBuilder.ExperimentID(v.ExperimentId).
			ParticipantId(v.ParticipantId).
			SubjectRecordID(v.SubjectRecordId).
			State(v.State)
		subjectRecords = append(subjectRecords, subjectRecordBuilder.Build())
	}
	experimentBuilder := &entity.ExperimentBuilder{}
	experimentBuilder.Description(experimentPO.Description).
		ExperimentID(experimentPO.ExperimentId).
		Title(experimentPO.Title).
		Description(experimentPO.Description).
		ResearcherId(experimentPO.ResearcherId).
		ExperimentTime(experimentPO.ExperimentTime).
		ParticipantNum(experimentPO.ParticipantNum)
	return experimentBuilder.Build()
}

func AssembleSubjectRecordEntity(
	subjectRecordPO *po.SubjectRecordPO,
) *entity.SubjectRecord {
	subjectRecordBuilder := &entity.SubjectRecordBuilder{}
	subjectRecordBuilder.ExperimentID(subjectRecordPO.ExperimentId).
		ParticipantId(subjectRecordPO.ParticipantId).
		SubjectRecordID(subjectRecordPO.SubjectRecordId).
		State(subjectRecordPO.State)
	return subjectRecordBuilder.Build()
}
