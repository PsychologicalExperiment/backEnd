package assembler

import (
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/application/dto"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/domain/entity"
)

func AssembleExperimentDTO(experimentEntity *entity.Experiment, experimentDTO *dto.ExperimentDTO) {
	experimentDTO.Description = experimentEntity.Description()
	experimentDTO.ExperimentId = experimentEntity.ExperimentId()
	experimentDTO.ExperimentTime = experimentEntity.ExperimentTime()
	experimentDTO.ParticipantNum = experimentEntity.ParticipantNum()
	experimentDTO.Title = experimentEntity.Title()
	experimentDTO.ParticipantNum = experimentEntity.ParticipantNum()
	for _, subjectRecordEntity := range experimentEntity.SubjectRecords() {
		subjectRecordDTO := &dto.SubjectRecordDTO{}
		AssembleSubjectRecordDTO(subjectRecordEntity, subjectRecordDTO)
		experimentDTO.SubjectRecords = append(experimentDTO.SubjectRecords, subjectRecordDTO)
	}
}

func AssembleSubjectRecordDTO(subjectRecordEntity *entity.SubjectRecord, subjectRecordDTO *dto.SubjectRecordDTO) {
	println("subject record: ", subjectRecordEntity.SubjectRecordId())
	subjectRecordDTO.SubjectRecordId = subjectRecordEntity.SubjectRecordId()
	subjectRecordDTO.ExperimentId = subjectRecordEntity.ExperimentId()
	subjectRecordDTO.ParticipantId = subjectRecordEntity.ParticipantId()
	subjectRecordDTO.State = subjectRecordEntity.State()
}
