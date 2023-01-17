package entity

import (
	"time"

	"gorm.io/gorm"
)

type SubjectRecordEntity struct {
	gorm.Model
	SubjectRecordId string    `gorm:"primaryKey;column:subject_record_id;type:varchar(36);unique;no null"`
	ExperimentId    string    `gorm:"column:experiment_id;type:varchar(36);no null"`
	ParticipantId   int64     `gorm:"column:participant_id"`
	State           int32     `gorm:"column:state"`
	FinishTime      time.Time `gorm:"column:finished_at"`
}
