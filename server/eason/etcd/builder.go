package etcd

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
)

type etcdBuilder struct {
	Client *clientv3.Client
}

func (s etcdBuilder) Scheme() string {
	return "etcd"
}

func (s *etcdBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
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

//func (s *etcdResolver) ResolveNow(rn resolver.ResolveNowOptions) {
//	// TODO: do something
//	return
//}
//
//func (s *etcdResolver) Close() {
//	// TODO: do something
//	return
//}
//
//func (s *etcdResolver) watch(keyPrefix string) {
//	var state resolver.State
//
//	etcdResp, err := cli.Get(context.Background(), keyPrefix, clientv3.WithPrefix())
//	if err != err {
//		fmt.Print(err)
//	} else {
//		for idx := range etcdResp.Kvs {
//			state.Addresses = append(state.Addresses, resolver.Address{Addr: strings.TrimPrefix(string(etcdResp.Kvs[idx].Key), keyPrefix)})
//		}
//	}
//	err = s.conn.UpdateState(state)
//	if err != nil {
//		fmt.Print(err)
//	}
//
//	rch := cli.Watch(context.Background(), keyPrefix, clientv3.WithPrefix())
//	for n := range rch {
//		for _, ev := range n.Events {
//			addr := strings.TrimPrefix(string(ev.Kv.Key), keyPrefix)
//			switch ev.Type {
//			case mvccpb.PUT:
//				if !exist(state.Addresses, addr) {
//					state.Addresses = append(state.Addresses, resolver.Address{Addr: addr})
//					err := s.conn.UpdateState(state)
//					if err != nil {
//						fmt.Print(err)
//					}
//				}
//			case mvccpb.DELETE:
//				if t, ok := remove(state, addr); ok {
//					state.Addresses = t
//					err := s.conn.UpdateState(state)
//					if err != nil {
//						fmt.Print(err)
//					}
//				}
//			}
//		}
//	}
//}
//
//func exist(l []resolver.Address, addr string) bool {
//	for i := range l {
//		if l[i].Addr == addr {
//			return true
//		}
//	}
//	return false
//}
//
//func remove(s resolver.State, addr string) ([]resolver.Address, bool) {
//	for i := range s.Addresses {
//		if s.Addresses[i].Addr == addr {
//			s.Addresses[i] = s.Addresses[len(s.Addresses)-1]
//			return s.Addresses[:len(s.Addresses)-1], true
//		}
//	}
//	return nil, false
//}
