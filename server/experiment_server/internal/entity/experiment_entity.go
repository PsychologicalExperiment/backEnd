package entity

import (
	"time"

	"gorm.io/gorm"
)

type ExperimentEntity struct {
	gorm.Model
	ExperimentId   string `gorm:"primaryKey;type:varchar(32);unique;no null"`
	Title          string `gorm:"type:varchar(512);unique;no null"`
	Description    string `gorm:"type:varchar(4096);no null"`
	ResearcherId   int64  `gorm:"column:researcher_id"`
	ExperimentTime int32  `gorm:"no null"`
	ParticipantNum int32  `gorm:"no null"`
	CurType        int32  `gorm:"column:cur_type"`
	Price          int64  `gorm:"column:price"`
	State          int32
	EndTime        time.Time `gorm:"column:end_time"`
	Url            string    `gorm:"column:url"`
}
