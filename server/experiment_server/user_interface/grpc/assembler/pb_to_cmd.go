package assembler

import (
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/application/dto/command"
	pb "github.com/PsychologicalExperiment/backEnd/api/experiment_server"
)

func AssembleAddExperimentCmd(
	req *pb.CreateExperimentReq, 
	cmd *command.AddExperimentCmd,
) {
	cmd.RequestId = req.RequestId
	cmd.Title = req.Title
	cmd.Description = req.Description
	cmd.ResearcherId = req.ResearcherId
	cmd.ExperimentTime = req.ExperimentTime
	cmd.ParticipantNum = req.ParticipantNum
}

func AssembleUpdateExperimentCmd(
	req *pb.UpdateExperimentReq,
	cmd *command.UpdateExperimentCmd,
) {
	cmd.RequestId = req.RequestId
	cmd.ExperimentId = req.ExperimentId
	cmd.Title = req.Title
	cmd.Description = req.Description
	cmd.ResearcherId = req.ResearcherId
	cmd.ExperimentTime = req.ExperimentTime
	cmd.ParticipantNum = req.ParticipantNum
	cmd.State = int32(req.State)
}

func AssembleAddSubjectRecordCmd(
	req *pb.CreateSubjectRecordReq,
	cmd *command.AddSubjectRecordCmd,
) {
	cmd.RequestId = req.RequestId
	cmd.ExperimentId = req.ExperimentId
	cmd.ParticipantId = req.ParticipantId
}


func AssembleUpdateSubjectRecordCmd(
	req *pb.UpdateSubjectRecordReq,
	cmd *command.UpdateSubjectRecordCmd,
) {
	cmd.RequestId = req.RequestId
	cmd.SubjectRecordId = req.SubjectRecordId
	cmd.UserId = req.UserId
	cmd.State = int32(req.State)
}