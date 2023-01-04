package etcd

import (
	"context"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	log "google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/resolver"
	"sync"
)

type Resolver struct {
	sync.RWMutex
	Client *clientv3.Client
	cc resolver.ClientConn
	prefix string
	addresses map[string]resolver.Address
}

func (r *Resolver) ResolveNow(options resolver.ResolveNowOptions) {
	// TODO
}

func (r *Resolver) Close() {
	// TODO
}

func (r *Resolver) watcher() {
	r.addresses = make(map[string]resolver.Address)
	response, err := r.Client.Get(context.Background(), r.prefix, clientv3.WithPrefix())
	if err != nil {
		log.Errorf("Get etcd:///%s error", r.prefix)
		return
	}
	for _, kv := range response.Kvs {
		r.setAddress(string(kv.Key), string(kv.Value))
	}

	if err := r.cc.UpdateState(resolver.State{
		Addresses: r.getAddresses(),
	}); err != nil {
		log.Errorf("etcd UpdateState error: %v", err)
		return
	}

	watch := r.Client.Watch(context.Background(), r.prefix, clientv3.WithPrefix())

	for resp := range watch {
		for _, event := range resp.Events {
			switch event.Type {
			case mvccpb.PUT:
				r.setAddress(string(event.Kv.Key), string(event.Kv.Value))
			case mvccpb.DELETE:
				r.delAddress(string(event.Kv.Key))
			}
		}
		if err := r.cc.UpdateState(resolver.State{
			Addresses: r.getAddresses(),
		}); err != nil {
			log.Errorf("etcd UpdateState error: %v", err)
		}
	}


}

func (r *Resolver) setAddress(key, address string) {
	r.Lock()
	defer r.Unlock()
	r.addresses[key] = resolver.Address{Addr: string(address)}
}

func (r *Resolver) delAddress(key string) {
	r.Lock()
	defer r.Unlock()
	delete(r.addresses, key)
}

func (r *Resolver) getAddresses() []resolver.Address {
	addresses := make([]resolver.Address, 0, len(r.addresses))

	for _, address := range r.addresses {
		addresses = append(addresses, address)
	}
	return addresses
}