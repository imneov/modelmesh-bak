package proxy

import (
	"context"
	"fmt"
	"github.com/imneov/modelmesh/internal/proxy/config"
	"github.com/imneov/modelmesh/internal/proxy/holder"
	proto "github.com/imneov/modelmesh/mindspore_serving/proto"
	xgrpc "github.com/imneov/modelmesh/pkg/transport/grpc"
	"github.com/imneov/modelmesh/pkg/utils"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"k8s.io/klog/v2"
	"net"
	"sync"
)

type Server struct {
	Config   *config.Config
	listener net.Listener
	Endpoint *grpc.Server //Broker grpc server
	dispatch *Dispatch
	holder   holder.Holder
	done     <-chan struct{}
	lock     sync.RWMutex
}

func (s *Server) PrepareRun(done <-chan struct{}) (err error) {
	s.done = done
	s.Endpoint = xgrpc.NewServer(s.Config.ProxyServer)

	if s.holder != nil {
		s.holder.Cancel()
	}
	s.holder = holder.New(context.Background())

	s.dispatch, err = NewDispatch(s.Config.Dispatch, s.holder)
	if err != nil {
		return err
	}

	if s.Endpoint == nil {
		return fmt.Errorf("broker server is nil")
	}
	proto.RegisterMSServiceServer(s.Endpoint, s)

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
			s.Endpoint.GracefulStop()
		case <-s.done:
			s.Endpoint.GracefulStop()
		}
	}()

	addr := s.Config.ProxyServer.Addr
	klog.V(0).Infof("Start listening on %s", addr)

	network, address := utils.ExtractNetAddress(addr)
	if network != "tcp" {
		return fmt.Errorf("unsupported protocol %s: %v", network, addr)
	}

	if s.listener, err = net.Listen(network, address); err != nil {
		return fmt.Errorf("unable to listen: %w", err)
	}

	var eg errgroup.Group
	eg.Go(func() error {
		return s.dispatch.Run(ctx)
	})
	eg.Go(func() error {
		return s.Endpoint.Serve(s.listener)
	})

	return eg.Wait()
}
