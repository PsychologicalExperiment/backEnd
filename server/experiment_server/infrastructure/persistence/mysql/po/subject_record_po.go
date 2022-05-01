package po

import (
	"time"

	"gorm.io/gorm"
)

type SubjectRecordPO struct {
	gorm.Model
	SubjectRecordID string    `gorm:"type:varchar(32);unique;no null"`
	ExperimentID    string    `gorm:"type:varchar(32);no null"`
	UserID          string    `gorm:"type:varchar(128);no null"`
	State           int32     
	FinishTime      time.Time 
}

func (s *SubjectRecordPO) TableName() string {
	return "subject_record_info"
}

