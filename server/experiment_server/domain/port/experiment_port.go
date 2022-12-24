package port

import (
	"context"

	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/domain/entity"
)

type ExperimentPort interface {
	SaveExperiment(ctx context.Context, exp *entity.Experiment) error 
	SaveSubjectRecord(ctx context.Context, record *entity.SubjectRecord) error

	UpdateExperiment(ctx context.Context, exp *entity.Experiment) error
	UpdateSubjectRecord(ctx context.Context, record *entity.SubjectRecord) error

	FindExperiment(ctx context.Context, exp_id string) (*entity.Experiment, error)
	FindExperimentsByResearcherID(ctx context.Context, id string, page int32, size int32) ([]*entity.Experiment, int32, error)

	FindSubjectRecord(ctx context.Context, id string) (*entity.SubjectRecord, error)
	FindSubjectRecordsByExpID(ctx context.Context, id string, page int32, size int32) ([]*entity.SubjectRecord, int32, error)
}
