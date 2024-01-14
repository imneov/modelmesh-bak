package picker

import (
	"math/rand"
	"sync"
	"time"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/resolver"
)

type RWRRPickerBuilder struct{}

func (pb *RWRRPickerBuilder) Build(info PickerBuildInfo) balancer.Picker {
	if len(info.ReadySCs) == 0 {
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
	}

	scs := []balancer.SubConn{}
	scToAddr := make(map[balancer.SubConn]resolver.Address)
	rwrr := newRwrr()
	var weight int32
	for sc, scInfo := range info.ReadySCs {
		scs = append(scs, sc)
		scToAddr[sc] = scInfo.Address
		weight = 1
		if scInfo.Address.Attributes != nil {
			val := scInfo.Address.Attributes.Value(WeightAttributeKey)
			weight = val.(int32)
		}
		rwrr.add(weight)
	}

	return &rwrrPicker{
		subConns: scs,
		scToAddr: scToAddr,
		// Start at a random index, as the same WRR balancer rebuilds a new
		// picker when SubConn states change, and we don't want to apply excess
		// load to the first server in the list.
		next: rwrr.next(),
		rwrr: rwrr,
	}
}

type rwrrPicker struct {
	// subConns is the snapshot of the weightedroundrobin balancer when this picker was
	// created. The slice is immutable. Each Get() will do a round robin
	// selection from it and return the selected SubConn.
	subConns []balancer.SubConn

	scToAddr map[balancer.SubConn]resolver.Address

	mu sync.Mutex

	rwrr *rwrr
	next int
}

func (p *rwrrPicker) Pick(opts balancer.PickInfo) (balancer.PickResult, error) {
	p.mu.Lock()
	n := len(p.subConns)
	p.mu.Unlock()
	if n == 0 {
		return balancer.PickResult{}, balancer.ErrNoSubConnAvailable
	}

	p.mu.Lock()
	sc := p.subConns[p.next]
	p.next = p.rwrr.next()
	p.mu.Unlock()
	return balancer.PickResult{SubConn: sc}, nil
}

type rwrr struct {
	weights      []int32
	sumOfWeights int32
	n            int
	rand         *rand.Rand
}

func newRwrr() *rwrr {
	return &rwrr{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (r *rwrr) next() int {
	if r.n == 0 || r.sumOfWeights == 0 {
		return 0
	}

	randomWeight := r.rand.Int31n(r.sumOfWeights) + 1
	for i, weight := range r.weights {
		randomWeight -= weight
		if randomWeight <= 0 {
			return i
		}
	}

	return r.rand.Intn(r.n)
}

func (r *rwrr) add(weight int32) {
	r.weights = append(r.weights, weight)
	r.sumOfWeights += weight
	r.n++
}
