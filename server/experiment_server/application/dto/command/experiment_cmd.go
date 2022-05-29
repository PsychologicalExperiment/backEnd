package command

import (
	"gopkg.in/validator.v2"
)

//  主试创建一个实验
type AddExperimentCmd struct {
	RequestId      string `validate:""`
	Title          string `validate:"min=1,max=128"`
	Description    string `validate:"min=0"`
	ResearcherId   string `validate:"min=1,max=32"`
	ExperimentTime int32  `validate:"min=0"`
	ParticipantNum int32  `validate:"min=0"`
}

func (a *AddExperimentCmd) CheckParam() error {

	if err := validator.Validate(a); err != nil {
		return err
	}
	return nil
}

//  更新一个实验
type UpdateExperimentCmd struct {
	RequestId      string `validate:""`
	ExperimentId   string `validate:"min=36,max=36"`
	Title          string `validate:"min=1,max=128"`
	Description    string `validate:"min=0"`
	ResearcherId   string `validate:"min=1,max=32"`
	ExperimentTime int32  `validate:"min=0"`
	ParticipantNum int32  `validate:"min=0"`
	State          int32  `validate:"min=0,max=3"`
}

func (u *UpdateExperimentCmd) CheckParam() error {
	if err := validator.Validate(u); err != nil {
		return err
	}
	return nil
}

//  增加一条被试记录
type AddSubjectRecordCmd struct {
	RequestId     string `validate:""`
	ExperimentId  string `validate:"min=0,max=36"`
	ParticipantId string `validate:"min=0,max=36"`
}

func (a *AddSubjectRecordCmd) CheckParam() error {

	if err := validator.Validate(a); err != nil {
		return err
	}
	return nil
}

//  更新一条被试记录
type UpdateSubjectRecordCmd struct {
	RequestId       string `validate:""`
	SubjectRecordId string `validate:"min=0,max=36"`
	UserId          string `validate:"min=0,max=36"`
	State           int32  `validate:"min=0"`
}

func (u *UpdateSubjectRecordCmd) CheckParam() error {

	if err := validator.Validate(u); err != nil {
		return err
	}
	return nil
}
