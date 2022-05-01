package query

//  查看单个实验详情
type GetExperimentQry struct {
	ExperimentID string
	UserID       string
}

//  查看实验列表
type GetExperimentListQry struct {
	UserID string
}

//  获取被试记录
type GetSubjectRecordsQry struct {
	ExperimentID string
}
