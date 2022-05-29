package query

import (
	"gopkg.in/validator.v2"
)

//  查看单个实验详情
type GetExperimentQry struct {
	RequestId    string `validate:"min=0"`
	ExperimentId string `validate:"min=36,max=36"`
}

//  查看实验列表
type GetExperimentListQry struct {
	RequestId    string `validate:"min=0"`
	ResearcherId string `validate:"min=1,max=32"`
	PageIndex    int32  `validate:"min=0"`
	PageSize     int32  `validate:"min=0"`
}

//  获取被试记录
type GetSubjectRecordQry struct {
	RequestId       string `validate:"min=0"`
	SubjectRecordId string `validate:"min=36,max=36"`
}

type GetSubjectRecordListQry struct {
	RequestId    string `validate:"min=0"`
	ExperimentId string `validate:"min=36,max=36"`
	PageIndex    int32  `validate:"min=0"`
	PageSize     int32  `validate:"min=0"`
}

func (g *GetExperimentQry) CheckParam() error {
	if err := validator.Validate(g); err != nil {
		return err
	}
	return nil
}

func (g *GetExperimentListQry) CheckParam() error {
	if err := validator.Validate(g); err != nil {
		return err
	}
	return nil
}

func (g *GetSubjectRecordQry) CheckParam() error {
	if err := validator.Validate(g); err != nil {
		return err
	}
	return nil
}

func (g *GetSubjectRecordListQry) CheckParam() error {
	if err := validator.Validate(g); err != nil {
		return err
	}
	return nil
}