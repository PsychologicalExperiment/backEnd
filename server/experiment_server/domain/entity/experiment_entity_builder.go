package entity

type ExperimentBuilder struct {
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

func (e *ExperimentBuilder) ExperimentID(id string) *ExperimentBuilder {
	e.experimentId = id
	return e
}

func (e *ExperimentBuilder) Title(title string) *ExperimentBuilder {
	e.title = title
	return e
}

func (e *ExperimentBuilder) Description(description string) *ExperimentBuilder {
	e.description = description
	return e
}

func (e *ExperimentBuilder) ResearcherId(id string) *ExperimentBuilder {
	e.researcherId = id
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

func (e *ExperimentBuilder) State(state int32) *ExperimentBuilder {
	e.state = state
	return e
}

func (e *ExperimentBuilder) CreateTime(time string) *ExperimentBuilder {
	e.createTime = time
	return e
}

func (e *ExperimentBuilder) UpdateTime(time string) *ExperimentBuilder {
	e.updateTime = time
	return e
}

func (e *ExperimentBuilder) SubjectRecordNum(num int32) *ExperimentBuilder {
	e.subjectRecordNum = num
	return e
}

func (e *ExperimentBuilder) Build() *Experiment {
	return &Experiment{
		experimentId:     e.experimentId,
		title:            e.title,
		description:      e.description,
		researcherId:     e.researcherId,
		experimentTime:   e.experimentTime,
		participantNum:   e.participantNum,
		state:            e.state,
		createTime:       e.createTime,
		updateTime:       e.updateTime,
		subjectRecordNum: e.subjectRecordNum,
		subjectRecords:   e.subjectRecords,
	}
}

type SubjectRecordBuilder struct {
	subjectRecordId string
	experimentId    string
	participantId   string
	state           int32
	timeTaken       string
}

func (s *SubjectRecordBuilder) SubjectRecordID(subjectRecord string) *SubjectRecordBuilder {
	s.subjectRecordId = subjectRecord
	return s
}

func (s *SubjectRecordBuilder) ExperimentID(experimentID string) *SubjectRecordBuilder {
	s.experimentId = experimentID
	return s
}

func (s *SubjectRecordBuilder) ParticipantId(userId string) *SubjectRecordBuilder {
	s.participantId = userId
	return s
}

func (s *SubjectRecordBuilder) State(state int32) *SubjectRecordBuilder {
	s.state = state
	return s
}

func (s *SubjectRecordBuilder) TimeTaken(time string) *SubjectRecordBuilder {
	s.timeTaken = time
	return s
}

func (s *SubjectRecordBuilder) Build() *SubjectRecord {
	return &SubjectRecord{
		subjectRecordId: s.subjectRecordId,
		experimentId:    s.experimentId,
		participantId:   s.participantId,
		state:           s.state,
		timeTaken:       s.timeTaken,
	}
}
