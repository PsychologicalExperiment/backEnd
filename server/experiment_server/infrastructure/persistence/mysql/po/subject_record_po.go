package po

import (
	"time"

	"gorm.io/gorm"
)

type SubjectRecordPO struct {
	gorm.Model
	SubjectRecordId string    `gorm:"primaryKey;column:subject_record_id;type:varchar(36);unique;no null"`
	ExperimentId    string    `gorm:"column:experiment_id;type:varchar(36);no null"`
	ParticipantId   string    `gorm:"column:participant_id;type:varchar(128);no null"`
	State           int32     `gorm:"column:state"`
	FinishTime      time.Time `gorm:"column:finished_at"`
}

func (s *SubjectRecordPO) TableName() string {
	return "subject_record_info"
}
