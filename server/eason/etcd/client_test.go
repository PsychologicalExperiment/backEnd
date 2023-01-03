package etcd

import (
	"reflect"
	"testing"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestGetEtcdClient(t *testing.T) {
	tests := []struct {
		name    string
		want    *clientv3.Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetEtcdClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEtcdClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEtcdClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
