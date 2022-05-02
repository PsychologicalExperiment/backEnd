package assembler

import (
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/application/dto/command"
	pb "github.com/PsychologicalExperiment/backEnd/api/experiment_server"
)

func AssembleNewExperimentCmd(req *pb.NewExperimentReq, cmd *command.AddExperimentCmd) {
	cmd.Title = req.ExpInfo.Title
	cmd.Description = req.ExpInfo.Description
	cmd.UserID = req.ExpInfo.ResearcherId
	cmd.ExperimentTime = req.ExpInfo.ExperimentTime
	cmd.ParticipantNum = req.ExpInfo.ParticipantNum
}