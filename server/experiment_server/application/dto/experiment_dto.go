package dto

type ExperimentDTO struct {
	ExperimentID   string
	Title          string
	InternalName   string
	Description    string
	UserID         string
	ExperimentTime int32
	ParticipantNum int32
	SubjectRecords []*SujectRecordDTO
}

type SujectRecordDTO struct {
	SubjectRecordID string
	ExperimentID    string
	UserID          string
	State           int32
}