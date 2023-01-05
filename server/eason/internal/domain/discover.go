package domain

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
)

type Discovery struct {}

func (s *Discovery) DiscoverServer(namespace, serverName string) (Addr string, err error) {
	cli, cerr := clientv3.NewFromURL("http://159.75.177:2379")
	if cerr != nil {
		return "", cerr
	}
	etcdResolver, err := resolver.NewBuilder(cli)
	conn, err := grpc.Dial(fmt.Sprintf("etcd:///%s/%s", namespace, serverName), grpc.WithResolvers(etcdResolver))
	if err != nil {
		return "", err
	}
	return conn.Target(), nil
}