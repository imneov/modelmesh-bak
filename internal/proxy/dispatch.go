package proxy

import (
	"context"
	"fmt"
	v1 "github.com/imneov/modelmesh/api/modelfulx/v1alpha"
	"github.com/imneov/modelmesh/internal/proxy/config"
	xgrpc "github.com/imneov/modelmesh/pkg/transport/grpc"
	"io"
	"k8s.io/klog/v2"
	"time"
)

type Dispatch struct {
	queue  *Queue
	client v1.MFServiceClient
}

func NewDispatch(cfg *config.Dispatch) (*Dispatch, error) {
	client := xgrpc.NewBrokerClient(cfg.Client)
	if client == nil {
		return nil, fmt.Errorf("client is nil")
	}
	return &Dispatch{
		client: client,
	}, nil
}

func (d *Dispatch) Run(ctx context.Context) error {
	stream, err := d.client.Predict(context.Background())
	if err != nil {
		return err
	}

	waitc := make(chan struct{})
	go func() {
		idx := 0
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				klog.V(0).Infof("Failed to receive a note : %v", err)
			}
			klog.V(0).Infof("Got message [%d]%v:%v", idx, &in, in.Id)
			idx++
		}
	}()
	idx := 0
	for {
		time.Sleep(1 * time.Millisecond)
		in := &v1.PredictRequest{
			Id: fmt.Sprintf("test-%d", idx),
			//Mindspore: new(proto.PredictRequest),
		}
		if err := stream.Send(in); err != nil {
			klog.V(0).Infof("Failed to send a note: %v", err)
		}
		klog.V(0).Infof("Send message [%d]%v:%v", idx, &in, in.Id)
		idx++
	}
	stream.CloseSend()
	<-waitc
	return nil
}
