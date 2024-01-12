package config

import xtime "github.com/imneov/modelmesh/pkg/utils/time"

// RPCClient is RPC client config.
type GRPCClient struct {
	Dial    xtime.Duration
	Timeout xtime.Duration
}

// RPCServer is RPC server config.
type GRPCServer struct {
	APPID                string
	Addr                 string
	Timeout              xtime.Duration
	IdleTimeout          xtime.Duration
	MaxLifeTime          xtime.Duration
	ForceCloseWait       xtime.Duration
	KeepAliveInterval    xtime.Duration
	KeepAliveTimeout     xtime.Duration
	MaxMessageSize       int32
	MaxConcurrentStreams int32
	EnableOpenTracing    bool
}

type BaseConfig struct {
	ProfPathPrefix string
	LogLevel       string
	ProfEnable     bool
	TracerEnable   bool
	WhiteList      []string
	Endpoints      []string
	BaseConfig     string
}

type ModelService struct {
	name string
}

type VirtualService struct {
	name string
}

type TrafficRule struct {
	name string
}
