package dao

import (
	"sync"

	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/infrastructure/persistence/mysql/po"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	once    sync.Once
	writeDB *gorm.DB
	readDB  *gorm.DB
)

type ExperimentDao struct {
	experimentPO po.ExperimentPO
}

func (e *ExperimentDao) TableName() string {
	return "experiment_info"
}

// WriteClient 写主机
func (e *ExperimentDao) WriteClient() *gorm.DB {
	once.Do(func() {
		dsn := "root:qianhaiwaibao@tcp(118.195.204.214:3306)/psychological_experiment?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			println("db error!")
			return
		}
		writeDB = db.Table(e.TableName())
		// expDB = expDB
	})
	// dsn := "root:qianhaiwaibao@tcp(127.0.0.1:3306)/psychological_experiment?charset=utf8mb4&parseTime=True&loc=Local"
	// expDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	println("db error!")
	// 	return &gorm.DB{}
	// }
	// expDB = expDB.Table(e.experimentPO.TableName())
	return writeDB
}

// ReadClient 读备机
func (e *ExperimentDao) ReadClient() *gorm.DB {
	once.Do(func() {
		dsn := "root:qianhaiwaibao@tcp(159.75.15.177:3306)/psychological_experiment?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			println("db error!")
			return
		}
		readDB = db.Table(e.TableName())
	})
	return readDB
}
