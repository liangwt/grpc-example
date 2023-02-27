package resolver

import "google.golang.org/grpc/resolver"

var _ resolver.Builder = Builder{}

type Builder struct {
	addrsStore map[string][]string
}

func NewResolverBuilder(addrsStore map[string][]string) *Builder {
	return &Builder{addrsStore: addrsStore}
}

func (b Builder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &Resolver{
		target:     target,
		cc:         cc,
		addrsStore: b.addrsStore,
	}
	r.Start()
	return r, nil
}

func (b Builder) Scheme() string {
	return "example"
}
