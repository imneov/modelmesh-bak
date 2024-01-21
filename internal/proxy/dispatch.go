package proxy

import (
	"context"
	"fmt"
	v1 "github.com/imneov/modelmesh/api/modelfulx/v1alpha"
	"github.com/imneov/modelmesh/internal/proxy/config"
	"github.com/imneov/modelmesh/internal/proxy/holder"
	xgrpc "github.com/imneov/modelmesh/pkg/transport/grpc"
	"io"
	"k8s.io/klog/v2"
)

type Dispatch struct {
	queue  *Queue
	client v1.MFServiceClient
	holder holder.Holder
}

func NewDispatch(cfg *config.Dispatch, holder holder.Holder) (*Dispatch, error) {
	client := xgrpc.NewBrokerClient(cfg.Client)
	if client == nil {
		return nil, fmt.Errorf("client is nil")
	}
	queue, err := NewQueue(cfg.Queue)
	return &Dispatch{
		client: client,
		holder: holder,
		queue:  queue,
	}, err
}

func (d *Dispatch) Run(ctx context.Context) error {
	stream, err := d.client.Predict(context.Background())
	if err != nil {
		return err
	}

	closeCh := make(chan struct{})
	go func() {
		idx := 0
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(closeCh)
				return
			}
			if err != nil {
				klog.V(0).Infof("Failed to receive a note : %v", err)
			}

			d.holder.Notify(&holder.Response{
				ID:     in.Id,
				Status: holder.StatusOK,
			})
			klog.V(0).Infof("Got message [%d]%v:%v", idx, &in, in.Id)
			idx++
		}
	}()

	for {
		ret := d.queue.Pop(ctx)
		if ret == nil {
			continue
		}

		in := &v1.PredictRequest{
			Id: ret.Id,
			//Mindspore: new(proto.PredictRequest),
		}
		err := stream.Send(in)
		klog.V(0).Infof("Send %v,%v", ret, err)
		select {
		case <-ctx.Done():
			return nil
		case <-closeCh:
			return nil
		default:
		}
	}

	stream.CloseSend()
	return nil
}

func (d *Dispatch) Wait(ctx context.Context, id string) holder.Response {
	return d.holder.Wait(ctx, id)
}
