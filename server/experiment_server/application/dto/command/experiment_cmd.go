package command

import (
	"gopkg.in/validator.v2"
)

//  主试创建一个实验
type AddExperimentCmd struct {
	Title          string `validate:"min=1,max=128"`
	Description    string `validate:"min=0"`
	UserID         string `validate:"min=1,max=32"`
	ExperimentTime int32  `validate:"min=0"`
	ParticipantNum int32  `validate:"min=0"`
}

func (a *AddExperimentCmd) CheckParam() error {
	
	if err := validator.Validate(a); err != nil {
		return err
	}
	return nil
}


// //  被试参与实验
// type PariticipateExperimentCmd struct {
// 	UserID       string
// 	ExperimentID string
// }

// //  被试完成实验
// type FinishExperimentCmd struct {
// 	SubjectRecordID string
// 	State           string
// }
