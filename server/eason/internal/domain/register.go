package domain

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Register struct {

}

// EtcdRegister 注册
func (r *Register) EtcdRegister(ctx context.Context, server, addr string, ttl int64) error {
	cli, err := clientv3.NewFromURL("http://159.75.177:2379")
	if err != nil {
		return err
	}
	lease, err := cli.Grant(ctx, ttl)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("/%s/%s", server, addr)
	_, err = cli.Put(ctx, key, addr, clientv3.WithLease(lease.ID))
	if err != nil {
		return err
	}
	keepAlive, err := cli.KeepAlive(ctx, lease.ID)
	if err != nil {
		return err
	}
	go func() {
		for {
			<-keepAlive
		}
	}()
	return nil
}

// EtcdUnRegister 反注册
func (r *Register) EtcdUnRegister(ctx context.Context, server, addr string) error {
	return nil
}