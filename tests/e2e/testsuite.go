package e2e

import (
	"context"
	"fmt"
	broker "github.com/imneov/modelmesh/cmd/broker/app"
	brokeroptions "github.com/imneov/modelmesh/cmd/broker/app/options"
	proxy "github.com/imneov/modelmesh/cmd/proxy/app"
	proxyoptions "github.com/imneov/modelmesh/cmd/proxy/app/options"
	brokerconfig "github.com/imneov/modelmesh/internal/broker/config"
	proxyconfig "github.com/imneov/modelmesh/internal/proxy/config"
	mindspore_serving_proto "github.com/imneov/modelmesh/mindspore_serving/proto"
	xconfig "github.com/imneov/modelmesh/pkg/config"
	"github.com/imneov/modelmesh/pkg/transport/grpc"
	"github.com/imneov/modelmesh/pkg/utils"
	"k8s.io/klog/v2"
	"net"
	"time"
)

var _handler = &handler{}

type handler struct {
}

func (h *handler) Predict(ctx context.Context, in *mindspore_serving_proto.PredictRequest) (*mindspore_serving_proto.PredictReply, error) {
	klog.V(0).Info("Predict")
	return &mindspore_serving_proto.PredictReply{
		ServableSpec: &mindspore_serving_proto.ServableSpec{
			Name: "Predict test",
		},
	}, nil
}

func runFakeClent(ctx context.Context) error {
	cfg, err := proxyconfig.TryLoadFromDisk()
	if err != nil {
		return err
	}
	client := grpc.NewServingClient(&xconfig.GRPCClient{
		Addr: cfg.ProxyServer.Addr,
	})

	for {
		_, err := client.Predict(ctx, &mindspore_serving_proto.PredictRequest{
			ServableSpec: &mindspore_serving_proto.ServableSpec{
				Name: "Predict test",
			},
		})
		if err != nil {
			klog.Fatal(err)
		}
		klog.V(0).Info("Predict")
		time.Sleep(1 * time.Second)
	}
	return nil
}

func runFakeServingServer(ctx context.Context) error {
	cfg, err := brokerconfig.TryLoadFromDisk()
	if err != nil {
		return err
	}
	s := grpc.NewServer(cfg.BrokerServer)

	mindspore_serving_proto.RegisterMSServiceServer(s, _handler)

	addr := cfg.Dispatch.Client.Addr // mock serving server address
	klog.V(0).Infof("Start listening on %s", addr)

	network, address := utils.ExtractNetAddress(addr)
	if network != "tcp" {
		return fmt.Errorf("unsupported protocol %s: %v", network, addr)
	}

	listener, err := net.Listen(network, address)
	if err != nil {
		return fmt.Errorf("unable to listen: %w", err)
	}

	return s.Serve(listener)
}

func runBroker(ctx context.Context) error {
	s := brokeroptions.NewServerRunOptions()

	// Load configuration from file /etc/schedule/schedule.yaml
	conf, err := brokerconfig.TryLoadFromDisk()
	if err == nil {
		s = &brokeroptions.ServerRunOptions{
			Config: conf,
		}
	} else {
		klog.Fatal("Failed to load configuration from disk", err)
	}
	return broker.Run(s, ctx)
}

func runProxy(ctx context.Context) error {
	s := proxyoptions.NewServerRunOptions()

	// Load configuration from file /etc/schedule/schedule.yaml
	conf, err := proxyconfig.TryLoadFromDisk()
	if err == nil {
		s = &proxyoptions.ServerRunOptions{
			Config: conf,
		}
	} else {
		klog.Fatal("Failed to load configuration from disk", err)
	}
	return proxy.Run(s, ctx)
}

func Testsuite(ctx context.Context) {
	Go(func() {
		runFakeServingServer(ctx)
	}, 3)
	Go(func() {
		runBroker(ctx)
	}, 3)
	Go(func() {
		runProxy(ctx)
	}, 3)
	Go(func() {
		runFakeClent(ctx)
	}, 1)
}

func Go(fn func(), sleep time.Duration) {
	go fn()
	klog.V(0).Infof("sleeping %d second...", sleep)
	time.Sleep((sleep) * time.Second)
}
