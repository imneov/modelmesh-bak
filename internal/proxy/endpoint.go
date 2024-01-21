package proxy

import (
	"context"
	v1 "github.com/imneov/modelmesh/api/modelfulx/v1alpha"
	proto "github.com/imneov/modelmesh/mindspore_serving/proto"
	"github.com/imneov/modelmesh/pkg/utils"
	"k8s.io/klog/v2"
)

func (s *Server) Predict(ctx context.Context, in *proto.PredictRequest) (*proto.PredictReply, error) {
	klog.V(0).Infof("--- Predict ---")

	requestID := utils.ReqUUID()
	request := &v1.PredictRequest{
		Id: requestID,
	}

	s.dispatch.queue.Push(ctx, request)
	// wait for response, dispatch will call back when response is ready.
	response := s.dispatch.Wait(ctx, requestID)
	klog.V(0).Infof("response: %v", response)
	//response.Data  @TODO: do something with the response data

	return &proto.PredictReply{}, nil
}
