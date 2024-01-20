package broker

import (
	"context"
	v1 "github.com/imneov/modelmesh/api/modelfulx/v1alpha"
	"github.com/imneov/modelmesh/internal/broker/picker"
	"io"
	"k8s.io/klog/v2"
	"time"
)

const (
	timestampFormat = time.StampNano
)

type PredictRequest struct {
	Request v1.PredictRequest
	Send    func(*v1.PredictReply) error
}

func (s *Server) Predict(stream v1.MFService_PredictServer) error {
	klog.V(7).Infof("--- Predict ---")
	//// Create trailer in defer to record function return time.
	//defer func() {
	//	trailer := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	//	stream.SetTrailer(trailer)
	//}()
	//
	//// Read metadata from client.
	//md, ok := metadata.FromIncomingContext(stream.Context())
	//if !ok {
	//	return status.Errorf(codes.DataLoss, "BidirectionalStreamingEcho: failed to get metadata")
	//}
	//
	//if t, ok := md["timestamp"]; ok {
	//	fmt.Printf("timestamp from metadata:\n")
	//	for i, e := range t {
	//		fmt.Printf(" %d. %s\n", i, e)
	//	}
	//}
	//
	////TODO read serviceGroup
	serviceGroup := getServiceGroup(stream)
	//
	//// Create and send header.
	//header := metadata.New(map[string]string{"location": "MTV", "timestamp": time.Now().Format(timestampFormat)})
	//stream.SendHeader(header)

	ctx := context.Background()
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		//应用按照配置推入不同的  group
		s.Push(ctx, serviceGroup, &PredictRequest{
			Request: *in,
			Send:    stream.Send,
		})
	}
}

func getServiceGroup(stream v1.MFService_PredictServer) string {
	//TODO read serviceGroup
	return "default"
}

func (s *Server) Process(ctx context.Context) error {
	for {
		ret, err := s.scheduler.Pick(picker.PickInfo{})
		if err != nil {
			klog.Errorf("scheduler pick error: %v", err)
			continue
		}
		group := ret.Resource.ResourceName()
		s.transGroupRequest(ctx, group)

		select {
		case <-ctx.Done():
			return nil
		default:
		}
	}

	return nil
}

func (s *Server) Push(ctx context.Context, group string, in *PredictRequest) error {
	s.lock.RLock()
	defer s.lock.RUnlock()

	queue, err := s.queue.Select(group)
	if err != nil {
		klog.Errorf("scheduler pick error: %v", err)
		return err
	}

	queue.Push(ctx, in)
	return nil
}

func (s *Server) transGroupRequest(ctx context.Context, group string) error {
	s.lock.RLock()
	defer s.lock.RUnlock()

	queue, err := s.queue.Select(group)
	if err != nil {
		klog.Errorf("scheduler pick error: %v", err)
		return err
	}

	request := queue.Pop(ctx)
	if request == nil {
		return nil
	}

	s.dispatch.queue.Push(ctx, request)

	return nil
}
