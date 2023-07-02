package balancer

import (
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
)

var _ base.PickerBuilder = &Builder{}

func NewBalancerBuilder() balancer.Builder {
	return base.NewBalancerBuilder("random_picker", &Builder{}, base.Config{HealthCheck: true})
}

type Builder struct {
}

func (b *Builder) Build(info base.PickerBuildInfo) balancer.Picker {
	if len(info.ReadySCs) == 0 {
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
	}

	var scs []balancer.SubConn
	for subConn := range info.ReadySCs {
		scs = append(scs, subConn)
	}

	return &Picker{
		subConns: scs,
	}
}
