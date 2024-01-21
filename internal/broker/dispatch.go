package broker

import (
	"context"
	"fmt"
	v1 "github.com/imneov/modelmesh/api/modelfulx/v1alpha"
	"github.com/imneov/modelmesh/internal/broker/config"
	proto "github.com/imneov/modelmesh/mindspore_serving/proto"
	xgrpc "github.com/imneov/modelmesh/pkg/transport/grpc"
	"k8s.io/klog/v2"
)

type Dispatch struct {
	client proto.MSServiceClient
	queue  *Queue
}

func NewDispatch(cfg *config.Dispatch) (*Dispatch, error) {
	client := xgrpc.NewServingClient(cfg.Client)
	if client == nil {
		return nil, fmt.Errorf("client is nil")
	}
	queue, err := NewQueue(cfg.Queue)
	return &Dispatch{
		client: client,
		queue:  queue,
	}, err
}

// Run dispatch request to serving server
// @TODO call serving server
// 1. pull Request from queues
// 2. call serving server
// 3. call stream send response
func (d *Dispatch) Run(ctx context.Context) error {
	for {
		ret := d.queue.Pop(ctx)
		if ret == nil {
			continue
		}
		err := ret.Send(&v1.PredictReply{Id: fmt.Sprintf("broker-%s", ret.Request.Id)})
		klog.V(0).Infof("Send %v,%v", ret, err)
		select {
		case <-ctx.Done():
			return nil
		default:
		}
	}

	return nil
}
