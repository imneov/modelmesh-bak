package broker

import (
	"context"
	"fmt"
	v1 "github.com/imneov/modelmesh/api/modelfulx/v1alpha"
	"github.com/imneov/modelmesh/internal/broker/config"
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
	Config        *config.Config
	listener      net.Listener
	Endpoint      *grpc.Server //Broker grpc server
	ServingClient proto.MSServiceClient
	queue         *QueuePool
	dispatch      *Dispatch
	scheduler     *Scheduler
	done          <-chan struct{}
	lock          sync.RWMutex
}

func (s *Server) PrepareRun(done <-chan struct{}) (err error) {
	s.done = done
	s.Endpoint = xgrpc.NewServer(s.Config.BrokerServer)

	s.queue, err = NewQueuePool(s.Config.Queue, s.Config.ServiceGroups)
	if err != nil {
		return fmt.Errorf("queues is nil")
	}

	s.scheduler, err = NewScheduler(s.Config.Schedule)
	if err != nil {
		return err
	}
	s.scheduler.RegeneratePicker(s.Config.ServiceGroups)

	s.dispatch, err = NewDispatch(s.Config.Dispatch)
	if err != nil {
		return err
	}

	if s.Endpoint == nil {
		return fmt.Errorf("broker server is nil")
	}
	v1.RegisterMFServiceServer(s.Endpoint, s)

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

	addr := s.Config.BrokerServer.Addr
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
		return s.Process(ctx)
	})
	eg.Go(func() error {
		return s.Endpoint.Serve(s.listener)
	})

	return eg.Wait()
}
