package picker

import "google.golang.org/grpc/balancer"

type Picker = balancer.Picker
type PickerKey = balancer.SubConn
