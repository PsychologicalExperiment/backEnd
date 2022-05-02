package adapter

import (
	"context"
	"reflect"
	"testing"

	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/domain/entity"
)

func TestExperiment_Save(t *testing.T) {
	type args struct {
		ctx              context.Context
		experimentEntity *entity.Experiment
	}
	tests := []struct {
		name    string
		e       *Experiment
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				ctx: context.Background(),
				experimentEntity: (&entity.ExperimentBuilder{}).Title("实验标题").ParticipantNum(10).Build(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Experiment{}
			if err := e.Save(tt.args.ctx, tt.args.experimentEntity); (err != nil) != tt.wantErr {
				t.Errorf("Experiment.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExperiment_Find(t *testing.T) {
	type args struct {
		ctx    context.Context
		exp_id string
	}
	tests := []struct {
		name    string
		e       *Experiment
		args    args
		want    *entity.Experiment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Experiment{}
			got, err := e.Find(tt.args.ctx, tt.args.exp_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Experiment.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Experiment.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
