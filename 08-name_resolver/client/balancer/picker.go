package balancer

import (
	"math/rand"

	"google.golang.org/grpc/balancer"
)

var _ balancer.Picker = &Picker{}

type Picker struct {
	subConns []balancer.SubConn
}

func (p *Picker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	index := rand.Intn(len(p.subConns))
	sc := p.subConns[index]
	return balancer.PickResult{SubConn: sc}, nil
}