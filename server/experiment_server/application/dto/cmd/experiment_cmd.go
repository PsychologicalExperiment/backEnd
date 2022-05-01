package cmd

//  主试创建一个实验
type AddExperimentCmd struct {
	Title          string
	InternalName   string
	Description    string
	UserID         string
	ExperimentTime int32
	ParticipantNum int32
}

//  被试参与实验
type PariticipateExperimentCmd struct {
	UserID       string
	ExperimentID string
}

//  被试完成实验
type FinishExperimentCmd struct {
	SubjectRecordID string
	State           string
}
