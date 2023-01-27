package resolver

import (
	"google.golang.org/grpc/resolver"
)

var _ resolver.Resolver = &Resolver{}

// impl google.golang.org/grpc/resolver.Resolver
type Resolver struct {
	target resolver.Target
	cc     resolver.ClientConn

	addrsStore map[string][]string
}

func (r *Resolver) Start() {
	// 在静态路由表中查询此 Endpoint 对应 addrs
	var addrs []resolver.Address
	for _, addr := range r.addrsStore[r.target.URL.Opaque] {
		addrs = append(addrs, resolver.Address{Addr: addr})
	}

	r.cc.UpdateState(resolver.State{Addresses: addrs})
}

func (r *Resolver) ResolveNow(resolver.ResolveNowOptions) {

}

func (r *Resolver) Close() {

}
