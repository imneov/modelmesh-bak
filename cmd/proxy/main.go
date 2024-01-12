package main

import (
	"context"
	"flag"
	mindspore_serving_proto "github.com/imneov/modelmesh/mindspore_serving/proto"
	"github.com/imneov/modelmesh/pkg/transport/grpc"
	"k8s.io/klog/v2"
	"os"
	"os/signal"
	"syscall"

	"github.com/imneov/modelmesh/pkg/service"
	"github.com/tkeel-io/kit/app"
	"github.com/tkeel-io/kit/log"
	"github.com/tkeel-io/kit/transport"
	// User import.
)

var (
	// Name app.
	Name string
	// HTTPAddr string.
	HTTPAddr string
	// GRPCAddr string.
	GRPCAddr string
)

func init() {
	flag.StringVar(&Name, "app", "demo", "app name.")
	flag.StringVar(&Name, "weight", "weight", "app weight.")
	flag.StringVar(&GRPCAddr, "grpc_addr", "0.0.0.0:5100", "grpc listen address.")
}

func main() {

	klog.InitFlags(nil)
	flag.Parse()

	//httpSrv := server.NewHTTPServer(HTTPAddr)
	//serverList := []transport.Server{httpSrv, grpcSrv}
	grpcSrv := grpc.New(GRPCAddr)
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
		service := service.NewMSServiceService(nil)
		mindspore_serving_proto.RegisterMSServiceServer(grpcSrv.GetServe(), service)
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
