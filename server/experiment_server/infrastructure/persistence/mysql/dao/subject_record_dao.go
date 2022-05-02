package dao

import (
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/infrastructure/persistence/mysql/po"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)


type SubjectRecordDAO struct {
	subjectRecordPO po.SubjectRecordPO
}

func (s *SubjectRecordDAO) Client() *gorm.DB {
	var db *gorm.DB
	once.Do(func() {
		dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
		db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		db = db.Table(s.subjectRecordPO.TableName())
	})
	return db 
}