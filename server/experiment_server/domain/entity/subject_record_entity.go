package entity

import "time"

type SubjectRecord struct {
	id              uint
	subjectRecordID string
	experimentID    string
	userID          string
	state           int32
	finishTime      time.Time
}

func (s *SubjectRecord) ID() uint {
	return s.id
}

func (s *SubjectRecord) setID(id uint) {
	s.id = id
}

func (s *SubjectRecord) SubjectRecordID() string {
	return s.subjectRecordID
}

func (s *SubjectRecord) setSubjectRecordID(id string) {
	s.subjectRecordID = id
}

func (s *SubjectRecord) setExperimentID(id string) {
	s.experimentID = id
}

func (s *SubjectRecord) ExperimentID() string {
	return s.experimentID
}

func (s *SubjectRecord) setUserID(id string) {
	s.userID = id
}

func (s *SubjectRecord) UserID() string {
	return s.userID
}

func (s *SubjectRecord) setState(state int32) {
	s.state = state
}

func (s *SubjectRecord) State() int32 {
	return s.state
}

func (s *SubjectRecord) FinishTime() time.Time {
	return s.finishTime
}

func (s *SubjectRecord) setFinishTime(t time.Time) {
	s.finishTime = t
}