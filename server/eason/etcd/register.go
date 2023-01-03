package etcd

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	ttl = 1
)

func Register(namespace, svrName, addr string) error {
	fmt.Println("register")
	if cli == nil {
		_, err := GetEtcdClient()
		if err != nil {
			return err
		}
	}

	etcdKey := fmt.Sprintf("/%s/%s/%s", namespace, svrName, addr)
	fmt.Println(etcdKey)
	etcdVal, err := cli.Get(context.Background(), etcdKey)
	fmt.Println("etcdVal: ", etcdVal)
	if err != nil {
		fmt.Printf("server:[%s] register error.", svrName)
	} else if etcdVal.Count == 0 {
		fmt.Println("Count: ", etcdVal)
		err := withAlive(etcdKey, addr)
		if err != nil {
			fmt.Printf("keep alive error:%s", err)
		}
	}

	return nil
}

func withAlive(key, addr string) error {
	leaseRsp, err := cli.Grant(context.Background(), 1000)
	if err != nil {
		return nil
	}
	_, err = cli.Put(context.Background(), key, addr, clientv3.WithLease(leaseRsp.ID))
	if err != nil {
		return err
	}
	_, err = cli.KeepAlive(context.Background(), leaseRsp.ID)
	if err != nil {
		fmt.Printf("keep alive error: %s", err)
		return err
	}
	return nil
}

func UnRegister(key string) error {
	if cli != nil {
		_, err := cli.Delete(context.Background(), key)
		if err != nil {
			fmt.Printf("delete [%s] erorr", key)
			return err
		}
	}
	return nil
}
