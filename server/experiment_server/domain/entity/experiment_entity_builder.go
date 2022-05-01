package entity

type ExperimentBuilder struct {
	experimentID   string
	title          string
	internalName   string
	description    string
	userID         string
	experimentTime int32
	participantNum int32
	subjectRecords []*SubjectRecord
}

func (e *ExperimentBuilder) ExperimentID(id string) *ExperimentBuilder {
	e.experimentID = id
	return e
}

func (e *ExperimentBuilder) Title(title string) *ExperimentBuilder {
	e.title = title
	return e
}

func (e *ExperimentBuilder) InternalName(internalName string) *ExperimentBuilder {
	e.internalName = internalName
	return e
}

func (e *ExperimentBuilder) Description(description string) *ExperimentBuilder {
	e.description = description
	return e
}

func (e *ExperimentBuilder) UserID(userID string) *ExperimentBuilder {
	e.userID = userID
	return e
}

func (e *ExperimentBuilder) ExperimentTime(time int32) *ExperimentBuilder {
	e.experimentTime = time
	return e
}

func (e *ExperimentBuilder) ParticipantNum(num int32) *ExperimentBuilder {
	e.participantNum = num
	return e
}

func (e *ExperimentBuilder) SubjectRecords(subjectRecords []*SubjectRecord) *ExperimentBuilder {
	e.subjectRecords = subjectRecords
	return e
}

func (e *ExperimentBuilder) Build() *Experiment {
	return &Experiment{
		experimentID:   e.experimentID,
		title:          e.title,
		internalName:   e.internalName,
		description:   e.description,
		userID:         e.userID,
		experimentTime: e.experimentTime,
		subjectRecords: e.subjectRecords,
	}
}

type SubejctRecordBuilder struct {
	subjectRecordID string
	experimentID    string
	userID          string
	state           int32
}

func (s *SubejctRecordBuilder) SubjectRecordID(subjectRecord string) *SubejctRecordBuilder {
	s.subjectRecordID = subjectRecord
	return s
}

func (s *SubejctRecordBuilder) ExperimentID(experimentID string) *SubejctRecordBuilder {
	s.experimentID = experimentID
	return s
}

func (s *SubejctRecordBuilder) UserID(userId string) *SubejctRecordBuilder {
	s.userID = userId
	return s
}

func (s *SubejctRecordBuilder) State(state int32) *SubejctRecordBuilder {
	s.state = state
	return s
}

func (s *SubejctRecordBuilder) Build() *SubjectRecord {
	return &SubjectRecord{
		subjectRecordID: s.subjectRecordID,
		experimentID:    s.experimentID,
		userID:          s.userID,
		state:           s.state,
	}
}
