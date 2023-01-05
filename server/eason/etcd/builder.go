package etcd

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
)

type Builder struct {
	Client *clientv3.Client
}

func (s Builder) Scheme() string {
	return "etcd"
}

func (s *Builder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	prefix := fmt.Sprintf("/%s/", target.URL.Path)
	r := &Resolver{
		Client: s.Client,
		cc: cc,
		prefix: prefix,
	}
	go func() {
		r.watcher()
	}()
	r.ResolveNow(resolver.ResolveNowOptions{})
	return r, nil
}
