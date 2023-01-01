package assembler

import (
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/domain/entity"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/infrastructure/persistence/mysql/po"
	log "google.golang.org/grpc/grpclog"
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
		ParticipantNum(experimentPO.ParticipantNum).
		CreateTime(experimentPO.CreatedAt.Format("2006-01-02 15:04:05")).
		UpdateTime(experimentPO.UpdatedAt.Format("2006-01-02 15:04:05"))
	return experimentBuilder.Build()
}

func AssembleSubjectRecordEntity(
	subjectRecordPO *po.SubjectRecordPO,
) *entity.SubjectRecord {
	log.Infof("subjectRecordPO: %+v", subjectRecordPO)
	subjectRecordBuilder := &entity.SubjectRecordBuilder{}
	subjectRecordBuilder.ExperimentID(subjectRecordPO.ExperimentId).
		ParticipantId(subjectRecordPO.ParticipantId).
		SubjectRecordID(subjectRecordPO.SubjectRecordId).
		State(subjectRecordPO.State)
	return subjectRecordBuilder.Build()
}
