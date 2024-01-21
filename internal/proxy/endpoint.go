package proxy

import (
	"context"
	v1 "github.com/imneov/modelmesh/api/modelfulx/v1alpha"
	proto "github.com/imneov/modelmesh/mindspore_serving/proto"
	"k8s.io/klog/v2"
)

func (s *Server) Predict(ctx context.Context, in *proto.PredictRequest) (*proto.PredictReply, error) {
	klog.V(7).Infof("--- Predict ---")
	request := &v1.PredictRequest{
		Id: "test",
	}
	s.dispatch.queue.Push(ctx, request)
	return &proto.PredictReply{}, nil
}
