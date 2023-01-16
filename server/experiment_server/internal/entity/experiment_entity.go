package entity

import (
	"gorm.io/gorm"
)

type ExperimentEntity struct {
	gorm.Model
	ExperimentId   string `gorm:"primaryKey;type:varchar(32);unique;no null"`
	Title          string `gorm:"type:varchar(512);unique;no null"`
	Description    string `gorm:"type:varcahr(4096);no null"`
	ResearcherId   string `gorm:"column:researcher_id;type:varchar(128);no null"`
	ExperimentTime int32  `gorm:"no null"`
	ParticipantNum int32  `gorm:"no null"`
	CurType        int32  `gorm:"column:cur_type"`
	Price          int32  `gorm:"column:price"`
	State          int32
}
