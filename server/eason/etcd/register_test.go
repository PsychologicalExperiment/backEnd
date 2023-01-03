package etcd

import "testing"

func TestRegister(t *testing.T) {
	type args struct {
		namespace string
		svrName   string
		addr      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "register experiment_server",
			args: args{
				namespace: "eason",
				svrName:   "experiment_server",
				addr:      "159.75.15.177",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Register(tt.args.namespace, tt.args.svrName, tt.args.addr); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_withAlive(t *testing.T) {
	type args struct {
		key  string
		addr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := withAlive(tt.args.key, tt.args.addr); (err != nil) != tt.wantErr {
				t.Errorf("withAlive() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnRegister(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UnRegister(tt.args.key)
		})
	}
}
