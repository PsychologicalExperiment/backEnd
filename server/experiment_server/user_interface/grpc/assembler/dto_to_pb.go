package assembler

import (
	pb "github.com/PsychologicalExperiment/backEnd/api/experiment_server"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/application/dto"
)

func AssembleNewExperimentResp(experimentDTO *dto.ExperimentDTO, newExperimentResp *pb.NewExperimentResp) {
	newExperimentResp.ExperimentId = experimentDTO.ExperimentID
}
