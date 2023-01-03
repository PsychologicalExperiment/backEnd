package dao

import (
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/infrastructure/persistence/mysql/po"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	subReadDB  *gorm.DB
	subWriteDB *gorm.DB
)

type SubjectRecordDAO struct {
	subjectRecordPO po.SubjectRecordPO
}

func (s *SubjectRecordDAO) WriteClient() *gorm.DB {
	dsn := "root:qianhaiwaibao@tcp(118.195.204.214:3306)/psychological_experiment?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		println("db error!")
		return &gorm.DB{}
	}
	subWriteDB = db.Table(s.subjectRecordPO.TableName())
	return subWriteDB
}

func (s *SubjectRecordDAO) ReadClient() *gorm.DB {
	dsn := "root:qianhaiwaibao@tcp(159.75.15.177:3306)/psychological_experiment?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		println("db error!")
		return &gorm.DB{}
	}
	subReadDB = db.Table(s.subjectRecordPO.TableName())
	return subReadDB
}
