package broker

import (
	"context"
	"fmt"
	v1 "github.com/imneov/modelmesh/api/modelfulx/v1alpha"
	"github.com/imneov/modelmesh/internal/broker/config"
	proto "github.com/imneov/modelmesh/mindspore_serving/proto"
	"google.golang.org/grpc"
	"k8s.io/klog/v2"
	"net"
	"strings"
)

type Server struct {
	Config        *config.Config
	listener      net.Listener
	BrokerServer  *grpc.Server //Broker grpc server
	ServingClient proto.MSServiceClient
	Queue         *Queue
	done          <-chan struct{}
}

func (s *Server) PrepareRun(done <-chan struct{}) (err error) {
	s.done = done
	if s.ServingClient == nil {
		return fmt.Errorf("serving client is nil")
	}

	v1.RegisterMFServiceServer(s.BrokerServer, s.Queue)

	return nil
}

func (s *Server) Run(ctx context.Context) (err error) {
	//err = s.waitForResourceSync(ctx)
	//if err != nil {
	//	return err
	//}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		select {
		case <-ctx.Done():
			s.BrokerServer.GracefulStop()
		case <-s.done:
			s.BrokerServer.GracefulStop()
		}
	}()

	addr := s.Config.BrokerServer.Addr
	klog.V(0).Infof("Start listening on %s", addr)

	network, address := ExtractNetAddress(addr)
	if network != "tcp" {
		return fmt.Errorf("unsupported protocol %s: %v", network, addr)
	}

	if s.listener, err = net.Listen(network, address); err != nil {
		return fmt.Errorf("unable to listen: %w", err)
	}

	err = s.BrokerServer.Serve(s.listener)
	return err
}

func ExtractNetAddress(apiAddress string) (string, string) {
	idx := strings.Index(apiAddress, "://")
	if idx < 0 {
		return "tcp", apiAddress
	}

	return apiAddress[:idx], apiAddress[idx+3:]
}
