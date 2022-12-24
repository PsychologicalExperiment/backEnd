package dao

import (
	"sync"

	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/infrastructure/persistence/mysql/po"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	once sync.Once
	expDB *gorm.DB
)

type ExperimentDao struct {
	experimentPO po.ExperimentPO
}

func (e *ExperimentDao) TableName() string {
	return "experiment_info"
}

func (e *ExperimentDao) Client() *gorm.DB {
	once.Do(func() {
		dsn := "root:qianhaiwaibao@tcp(127.0.0.1:3306)/psychological_experiment?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			println("db error!")
			return
		}
		expDB = db.Table(e.TableName())
		// expDB = expDB
	})
	// dsn := "root:qianhaiwaibao@tcp(127.0.0.1:3306)/psychological_experiment?charset=utf8mb4&parseTime=True&loc=Local"
	// expDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	println("db error!")
	// 	return &gorm.DB{}
	// }
	// expDB = expDB.Table(e.experimentPO.TableName())
	return expDB
}
