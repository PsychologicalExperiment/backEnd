package assembler

import (
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/application/dto/query"
	pb "github.com/PsychologicalExperiment/backEnd/api/experiment_server"
)

func AssembleGetExperimentQry(
	req *pb.QueryExperimentReq, 
	qry *query.GetExperimentQry,
) {
	qry.RequestId = req.RequestId 
	qry.ExperimentId = req.ExperimentId
}

func AssembleGetExperimentListQry(
	req *pb.QueryExperimentListReq,
	qry *query.GetExperimentListQry,
) {
	qry.RequestId = req.RequestId 
	qry.ResearcherId = req.ResearcherId
	qry.PageIndex = req.PageIndex
	qry.PageSize = req.PageSize
}

func AssembleGetSubjectRecordQry(
	req *pb.QuerySubjectRecordReq,
	qry *query.GetSubjectRecordQry,
) {
	qry.RequestId = req.RequestId	
	qry.SubjectRecordId = req.SubjectRecordId
}

func AssembleGetSubjectRecordListQry(
	req *pb.QuerySubjectRecordListReq,
	qry *query.GetSubjectRecordListQry,
) {
	qry.RequestId = req.RequestId	
	qry.ExperimentId = req.ExperimentId
	qry.PageSize = req.PageSize
	qry.PageIndex = req.PageIndex
}