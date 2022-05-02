package port

import (
	"context"

	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/domain/entity"
)

type ExperimentPort interface {
	Save(ctx context.Context, exp *entity.Experiment) error 
	Find(ctx context.Context, exp_id string) (*entity.Experiment, error)
}
