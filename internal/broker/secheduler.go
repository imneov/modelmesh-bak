package broker

import "github.com/database-mesh/pisanix/pisa-controller/pkg/proxy"

type Secheduler interface {
}

type secheduler struct {
	proxy.QoSGroup
}

func NewSecheduler() Secheduler {
	return &secheduler{}
}
