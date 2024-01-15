package picker

import (
	"math/rand"
	"sync"

	"google.golang.org/grpc/balancer"
)

type RRPickerBuilder struct {
}

func (pb *RRPickerBuilder) Build(info PickerBuildInfo) Picker {
	if len(info.ReadySCs) == 0 {
		return NewErrPicker(balancer.ErrNoSubConnAvailable)
	}
	scs := make([]PickerKey, 0)
	//scToAddr := make(map[PickerKey]resolver.Address)
	//for sc, scInfo := range info.ReadySCs {
	for sc, _ := range info.ReadySCs {
		scs = append(scs, sc)
		//scToAddr[sc] = scInfo.Address
	}

	return &rrPicker{
		subConns: scs,
		//scToAddr: scToAddr,
		// Start at a random index, as the same RR balancer rebuilds a new
		// picker when SubConn states change, and we don't want to apply excess
		// load to the first server in the list.
		next: rand.Intn(len(scs)),
	}
}

type rrPicker struct {
	// subConns is the snapshot of the roundrobin balancer when this picker was
	// created. The slice is immutable. Each Get() will do a round robin
	// selection from it and return the selected SubConn.
	subConns []PickerKey

	//scToAddr map[PickerKey]resolver.Address

	mu   sync.Mutex
	next int
}

func (p *rrPicker) Pick(opts balancer.PickInfo) (balancer.PickResult, error) {
	p.mu.Lock()
	n := len(p.subConns)
	p.mu.Unlock()
	if n == 0 {
		return balancer.PickResult{}, balancer.ErrNoSubConnAvailable
	}

	p.mu.Lock()
	sc := p.subConns[p.next]
	// picked := p.scToAddr[sc]
	p.next = (p.next + 1) % len(p.subConns)
	p.mu.Unlock()

	done := func(info balancer.DoneInfo) {
		// log.Println("rrpicker done", picked.Addr)
	}

	return balancer.PickResult{SubConn: sc, Done: done}, nil
}
