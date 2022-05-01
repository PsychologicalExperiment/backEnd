package po

import (
	"gorm.io/gorm"
)

type ExperimentPO struct {
	gorm.Model
	ExperimentID   string `gorm:"type:varchar(32);unique;no null"`
	Title          string `gorm:"type:varchar(512);unique;no null"`
	InternalName   string `gorm:"type:varchar(128);no null"`
	Description    string `gorm:"type:varcahr(4096);no null"`
	UserID         string `gorm:"type:varchar(128);no null"`
	ExperimentTime int32  `gorm:"no null"`
	ParticipantNum int32  
}

func (s *ExperimentPO) TableName() string {
	return "experiment_info"
}
