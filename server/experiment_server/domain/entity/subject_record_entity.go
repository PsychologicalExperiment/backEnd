package entity

import (
	uuid "github.com/satori/go.uuid"
)

type SubjectRecord struct {
	id              uint
	subjectRecordId string
	experimentId    string
	participantId   string
	state           int32
	timeTaken       string
}

func (s *SubjectRecord) ID() uint {
	return s.id
}

func (s *SubjectRecord) setID(id uint) {
	s.id = id
}

func (s *SubjectRecord) SubjectRecordId() string {
	return s.subjectRecordId
}

func (s *SubjectRecord) setSubjectRecordId(id string) {
	s.subjectRecordId = id
}

func (s *SubjectRecord) setExperimentId(id string) {
	s.experimentId = id
}

func (s *SubjectRecord) ExperimentId() string {
	return s.experimentId
}

func (s *SubjectRecord) setParticipantId(id string) {
	s.participantId = id
}

func (s *SubjectRecord) ParticipantId() string {
	return s.participantId
}

func (s *SubjectRecord) setState(state int32) {
	s.state = state
}

func (s *SubjectRecord) State() int32 {
	return s.state
}

func (s *SubjectRecord) TimeTaken() string {
	return s.timeTaken
}

func (s *SubjectRecord) setFinishTime(t string) {
	s.timeTaken = t
}

func (s *SubjectRecord) GenSubjectRecordId() {
	s.subjectRecordId = uuid.NewV4().String()
}
