package dao

import (
	"sync"

	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/infrastructure/mysql/po"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	once sync.Once
)

type ExperimentDao struct {
	experimentPO po.ExperimentPO
}

func (e *ExperimentDao) TableName() string {
	return "experiment_info"
}

func (e *ExperimentDao) Client() *gorm.DB {
	var db *gorm.DB
	once.Do(func() {
		dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
		db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		db = db.Table(e.experimentPO.TableName())
	})
	return db
}
