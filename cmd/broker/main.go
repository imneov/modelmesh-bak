package main

import (
	"context"
	"flag"
	mindspore_proto "github.com/imneov/modelmesh/mindspore_serving/proto"
	"github.com/imneov/modelmesh/pkg/config"
	"k8s.io/klog/v2"
	"os"
	"os/signal"
	"syscall"

	"github.com/imneov/modelmesh/pkg/service"
	"github.com/imneov/modelmesh/pkg/transport/grpc"
	"github.com/tkeel-io/kit/app"
	"github.com/tkeel-io/kit/log"
	"github.com/tkeel-io/kit/transport"
	// User import.
)

var (
	// Name app.
	Name string
	// GRPCAddr string.
	GRPCAddr string
	// GRPCAddr string.
	ModelServingGRPCAddr string
)

func init() {
	flag.StringVar(&Name, "name", "model-service-broker", "app name.")
	flag.StringVar(&GRPCAddr, "grpc_addr", ":8080", "grpc listen address.")
	flag.StringVar(&ModelServingGRPCAddr, "serving_addr", "127.0.0.1:5100", "model serving grpc listen address.")
}

func main() {
	flag.Parse()

	cfg, err := config.TryLoadFromDisk()
	if err != nil {
		klog.Exit(err)
	}

	//httpSrv := server.NewHTTPServer(HTTPAddr)
	//serverList := []transport.Server{httpSrv, grpcSrv}
	grpcSrv := grpc.New(cfg.GRPCServer)
	serverList := []transport.Server{grpcSrv}

	app := app.New(Name,
		&log.Conf{
			App:   Name,
			Level: "debug",
			Dev:   true,
		},
		serverList...,
	)

	{ // User service
		conn, err := grpc.Dial(ModelServingGRPCAddr, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		client := mindspore_proto.NewMSServiceClient(conn)
		service := service.NewMSServiceService(client)
		mindspore_proto.RegisterMSServiceServer(grpcSrv.GetServe(), service)
	}

	if err := app.Run(context.TODO()); err != nil {
		panic(err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt)
	<-stop

	if err := app.Stop(context.TODO()); err != nil {
		panic(err)
	}
}
