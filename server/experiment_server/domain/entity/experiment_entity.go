package entity

type Experiment struct {
	experimentID   string
	title          string
	internalName   string
	description    string
	userID         string
	experimentTime int32
	participantNum int32
	subjectRecords []*SubjectRecord
}

func (e *Experiment) ExperimentID() string {
	return e.experimentID
}

func (e *Experiment) setExperimentID(experimentID string) {
	e.experimentID = experimentID
}

func (e *Experiment) Title() string {
	return e.title
}

func (e *Experiment) setTitle(title string) {
	e.title = title
}

func (e *Experiment) InternalName() string {
	return e.internalName
}

func (e *Experiment) setInternalName(internalName string) {
	e.internalName = internalName
}

func (e *Experiment) Description() string {
	return e.description
}

func (e *Experiment) setDescription(description string) {
	e.description = description
}

func (e *Experiment) UserID() string {
	return e.userID
}

func (e *Experiment) setUserID(userid string) {
	e.userID = userid
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
