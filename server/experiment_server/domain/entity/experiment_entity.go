package entity

import (
	uuid "github.com/satori/go.uuid"
)

type Experiment struct {
	id               uint
	experimentId     string
	title            string
	description      string
	researcherId     string
	experimentTime   int32
	participantNum   int32
	state            int32
	createTime       string
	updateTime       string
	subjectRecordNum int32
	subjectRecords   []*SubjectRecord
}

func (e *Experiment) ID() uint {
	return e.id
}

func (e *Experiment) setID(id uint) {
	e.id = id
}

func (e *Experiment) ExperimentId() string {
	return e.experimentId
}

func (e *Experiment) setExperimentId(experimentId string) {
	e.experimentId = experimentId
}

func (e *Experiment) Title() string {
	return e.title
}

func (e *Experiment) setTitle(title string) {
	e.title = title
}

func (e *Experiment) Description() string {
	return e.description
}

func (e *Experiment) setDescription(description string) {
	e.description = description
}

func (e *Experiment) ResearcherId() string {
	return e.researcherId
}

func (e *Experiment) setResearcherId(id string) {
	e.researcherId = id
}

func (e *Experiment) ExperimentTime() int32 {
	return e.experimentTime
}

func (e *Experiment) setExperimentTime(time int32) {
	e.experimentTime = time
}

func (e *Experiment) ParticipantNum() int32 {
	return e.participantNum
}

func (e *Experiment) setParticipantName(num int32) {
	e.participantNum = num
}

func (e *Experiment) SubjectRecords() []*SubjectRecord {
	return e.subjectRecords
}

func (e *Experiment) setSubjectRecords(subjectRecords []*SubjectRecord) {
	e.subjectRecords = subjectRecords
}

//  生成实验ID
func (e *Experiment) GenExperimentID() {
	e.experimentId = uuid.NewV4().String()
}

func (e *Experiment) State() int32 {
	return e.state 
}

func (e *Experiment) setState(state int32) {
	e.state = state
}

func (e *Experiment) CreateTime() string {
	return e.createTime
}

func (e *Experiment) setCreateTime(time string) {
	e.createTime = time
}

func (e *Experiment) UpdateTime() string {
	return e.updateTime
}

func (e *Experiment) setUpdateTime(time string){
	e.updateTime = time 
}

func (e *Experiment) SubjectRecordNum() int32 {
	return e.subjectRecordNum
}

func (e *Experiment) setSubjectRecordNum(num int32) {
	e.subjectRecordNum = num
}