package dto

type ExperimentDTO struct {
	ExperimentId     string
	Title            string
	Description      string
	ResearcherId     string
	ExperimentTime   int32
	ParticipantNum   int32
	State            int32
	CreateTime       string
	UpdateTime       string
	SubjectRecordNum int32
	SubjectRecords   []*SubjectRecordDTO
}

type SubjectRecordDTO struct {
	SubjectRecordId string
	ExperimentId    string
	ParticipantId   string
	TimeTaken       string
	State           int32
}
