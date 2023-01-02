package etcd

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sync"
)

var (
	once sync.Once
	cli  *clientv3.Client
)

// GetEtcdClient get etcd client
func GetEtcdClient() (*clientv3.Client, error) {
	once.Do(func() {
		var err error
		cli, err = clientv3.NewFromURL("http://159.75.15.177:2379")
		if err != nil {
			fmt.Println("get etcd client error: ", err)
			panic("get ectd client erorr")
		}
	})
	return cli, nil
}
