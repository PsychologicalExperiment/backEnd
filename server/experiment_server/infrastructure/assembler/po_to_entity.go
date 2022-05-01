package assembler

import (
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/infrastructure/persistence/mysql/po"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/domain/entity"

)

func AssembleExperimentPO(experimentEntity *entity.Experiment, experimentPO *po.ExperimentPO){
	return &po.ExperimentPO{
		ExperimentID: experimentEntity
	}
}


func AssembleSubjectRecord(subjectRecord *entity.SubjectRecord, subjectRecordPO *po.SubjectRecordPO) {

}
