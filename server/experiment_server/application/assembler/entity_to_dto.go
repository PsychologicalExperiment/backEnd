package assembler

import (
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/application/dto"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/domain/entity"
)

func AssembleExperimentDTO(experimentEntity *entity.Experiment, experimentDTO *dto.ExperimentDTO) {
	experimentDTO.Description = experimentEntity.Description()
	experimentDTO.ExperimentID = experimentEntity.ExperimentID()
	experimentDTO.ExperimentTime = experimentEntity.ExperimentTime()
	experimentDTO.InternalName = experimentEntity.InternalName()
	experimentDTO.ParticipantNum = experimentEntity.ParticipantNum()
	experimentDTO.Title = experimentEntity.Title()
	experimentDTO.UserID = experimentEntity.UserID()
	for _, subjectRecordEntity := range experimentEntity.SubjectRecords() {
		subjectRecordDTO := &dto.SujectRecordDTO{}
		AssembleSubjectRecordDTO(subjectRecordEntity, subjectRecordDTO)
		experimentDTO.SubjectRecords = append(experimentDTO.SubjectRecords, subjectRecordDTO)
	}
}

func AssembleSubjectRecordDTO(subjectRecordEntity *entity.SubjectRecord, subjectRecordDTO *dto.SubjectRecordDTO) {
	subjectRecordDTO.SubjectRecordID = subjectRecordEntity.SubjectRecordID()
	subjectRecordDTO.ExperimentID = subjectRecordEntity.ExperimentID()
	subjectRecordDTO.UserID = subjectRecordEntity.UserID()
	subjectRecordDTO.State = subjectRecordEntity.State()
}
