package dao

import (
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/infrastructure/persistence/mysql/po"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var (
	subDb *gorm.DB
)

type SubjectRecordDAO struct {
	subjectRecordPO po.SubjectRecordPO
}

func (s *SubjectRecordDAO) Client() *gorm.DB {
	// once.Do(func() {
	// 	dsn := "root:qianhaiwaibao@tcp(127.0.0.1:3306)/psychological_experiment?charset=utf8mb4&parseTime=True&loc=Local"
	// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// 	if err != nil {
	// 		println("db error!")
	// 		return
	// 	}
	// 	subDb = db.Table(s.subjectRecordPO.TableName())
	// })
	dsn := "root:qianhaiwaibao@tcp(127.0.0.1:3306)/psychological_experiment?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		println("db error!")
		return &gorm.DB{}
	}
	subDb = db.Table(s.subjectRecordPO.TableName())
	return subDb 
}