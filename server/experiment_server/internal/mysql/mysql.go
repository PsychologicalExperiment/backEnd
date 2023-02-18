package mysql

import (
	"context"
	"fmt"

	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/internal/entity"
	"github.com/PsychologicalExperiment/backEnd/util/plugins/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	masterDB *gorm.DB
	slaveDB  *gorm.DB
)

const (
	experimentInfoTableName = "experiment_info"
	subjectRecordTableName  = "subject_record_info"
)

func MasterClient() (*gorm.DB, error) {
	if masterDB == nil {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/psychological_experiment?charset=utf8mb4&parseTime=True&loc=Local",
			config.Config.Db.Master.User, config.Config.Db.Master.Passwd,
			config.Config.Db.Master.IP, config.Config.Db.Master.Port)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		masterDB = db
	}
	return masterDB, nil
}

func SlaveClient() (*gorm.DB, error) {
	if slaveDB == nil {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/psychological_experiment?charset=utf8mb4&parseTime=True&loc=Local",
			config.Config.Db.Slave.User, config.Config.Db.Slave.Passwd,
			config.Config.Db.Slave.IP, config.Config.Db.Slave.Port)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		slaveDB = db
	}
	return slaveDB, nil
}

type ExperimentDao interface {
	SaveExperiment(ctx context.Context, exp *entity.ExperimentEntity) (string, error)
	SaveSubjectRecord(ctx context.Context, record *entity.SubjectRecordEntity) (string, error)

	UpdateExperiment(ctx context.Context, exp *entity.ExperimentEntity) error
	UpdateSubjectRecord(ctx context.Context, record *entity.SubjectRecordEntity) error

	FindExperiment(context.Context, string) (*entity.ExperimentEntity, error)
	FindExperimentsB(context.Context, QueryExperimentReq) ([]*entity.ExperimentEntity, int64, error)

	FindSubjectRecord(context.Context, string) (*entity.SubjectRecordEntity, error)
	FindSubjectRecords(context.Context, QuerySubjectRecordReq) ([]*entity.SubjectRecordEntity, int32, error)

	CheckSubscribe(context.Context, string, int64) (int32, error)
}

// QueryExperimentReq 查询条件
type QueryExperimentReq struct {
	ResearcherId  int64
	Offset        int
	Limit         int
	MinPrice      int64
	EndTime       int64
	OnlySeeMyself int32
}

// QuerySubjectRecordReq 查询条件
type QuerySubjectRecordReq struct {
	ExperimentId string
	Offset       int
	Limit        int
}
