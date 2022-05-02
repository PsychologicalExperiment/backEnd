package grpc

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/PsychologicalExperiment/backEnd/api/experiment_server"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/application/service"
	"github.com/PsychologicalExperiment/backEnd/server/experiment_server/infrastructure/adapter"
)

func TestExperimentServiceImpl_NewExperiment(t *testing.T) {
	type fields struct {
		ApplicationService                   *service.ApplicationService
		UnimplementedExperimentServiceServer pb.UnimplementedExperimentServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.NewExperimentReq
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *pb.NewExperimentResp
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			fields: fields{
				ApplicationService: &service.ApplicationService{
					ExperimentPort: &adapter.Experiment{},
				},
			},
			args: args{
				ctx: context.Background(),
				req: &pb.NewExperimentReq{
					RequestId: "12121",
					ExpInfo: &pb.ExperimentInfo{
						Description: "实验描述",
						Title:       "实验名称",
					},
				},
			},
			wantResp: &pb.NewExperimentResp{

			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ExperimentServiceImpl{
				ApplicationService:                   tt.fields.ApplicationService,
				UnimplementedExperimentServiceServer: tt.fields.UnimplementedExperimentServiceServer,
			}
			gotResp, err := e.NewExperiment(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExperimentServiceImpl.NewExperiment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("ExperimentServiceImpl.NewExperiment() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestExperimentServiceImpl_QueryExperiment(t *testing.T) {
	type fields struct {
		ApplicationService                   *service.ApplicationService
		UnimplementedExperimentServiceServer pb.UnimplementedExperimentServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.QueryExperimentReq
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *pb.QueryExperimentResp
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ExperimentServiceImpl{
				ApplicationService:                   tt.fields.ApplicationService,
				UnimplementedExperimentServiceServer: tt.fields.UnimplementedExperimentServiceServer,
			}
			gotResp, err := e.QueryExperiment(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExperimentServiceImpl.QueryExperiment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("ExperimentServiceImpl.QueryExperiment() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestExperimentServiceImpl_QueryExperimentList(t *testing.T) {
	type fields struct {
		ApplicationService                   *service.ApplicationService
		UnimplementedExperimentServiceServer pb.UnimplementedExperimentServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.QueryExperimentListReq
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *pb.QueryExperimentListResp
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ExperimentServiceImpl{
				ApplicationService:                   tt.fields.ApplicationService,
				UnimplementedExperimentServiceServer: tt.fields.UnimplementedExperimentServiceServer,
			}
			gotResp, err := e.QueryExperimentList(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExperimentServiceImpl.QueryExperimentList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("ExperimentServiceImpl.QueryExperimentList() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestExperimentServiceImpl_CreateSubjectRecord(t *testing.T) {
	type fields struct {
		ApplicationService                   *service.ApplicationService
		UnimplementedExperimentServiceServer pb.UnimplementedExperimentServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.CreateSubjectRecordReq
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *pb.CreateSubjectRecordResp
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ExperimentServiceImpl{
				ApplicationService:                   tt.fields.ApplicationService,
				UnimplementedExperimentServiceServer: tt.fields.UnimplementedExperimentServiceServer,
			}
			gotResp, err := e.CreateSubjectRecord(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExperimentServiceImpl.CreateSubjectRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("ExperimentServiceImpl.CreateSubjectRecord() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestExperimentServiceImpl_UpdateSubjectRecord(t *testing.T) {
	type fields struct {
		ApplicationService                   *service.ApplicationService
		UnimplementedExperimentServiceServer pb.UnimplementedExperimentServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.UpdateSubjectRecordReq
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *pb.UpdateSubjectRecordResp
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ExperimentServiceImpl{
				ApplicationService:                   tt.fields.ApplicationService,
				UnimplementedExperimentServiceServer: tt.fields.UnimplementedExperimentServiceServer,
			}
			gotResp, err := e.UpdateSubjectRecord(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExperimentServiceImpl.UpdateSubjectRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("ExperimentServiceImpl.UpdateSubjectRecord() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestExperimentServiceImpl_QuerySubjectRecord(t *testing.T) {
	type fields struct {
		ApplicationService                   *service.ApplicationService
		UnimplementedExperimentServiceServer pb.UnimplementedExperimentServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.QuerySubjectRecordReq
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *pb.QuerySubjectRecordResp
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ExperimentServiceImpl{
				ApplicationService:                   tt.fields.ApplicationService,
				UnimplementedExperimentServiceServer: tt.fields.UnimplementedExperimentServiceServer,
			}
			gotResp, err := e.QuerySubjectRecord(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExperimentServiceImpl.QuerySubjectRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("ExperimentServiceImpl.QuerySubjectRecord() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestExperimentServiceImpl_QuerySubjectRecordList(t *testing.T) {
	type fields struct {
		ApplicationService                   *service.ApplicationService
		UnimplementedExperimentServiceServer pb.UnimplementedExperimentServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.QuerySubjectRecordListReq
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *pb.QuerySubjectRecordListResp
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ExperimentServiceImpl{
				ApplicationService:                   tt.fields.ApplicationService,
				UnimplementedExperimentServiceServer: tt.fields.UnimplementedExperimentServiceServer,
			}
			gotResp, err := e.QuerySubjectRecordList(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExperimentServiceImpl.QuerySubjectRecordList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("ExperimentServiceImpl.QuerySubjectRecordList() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
