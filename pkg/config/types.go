package xconfig

import (
	"time"
)

// RPCClient is RPC client config.
type GRPCClient struct {
	Addr    string        `json:"addr,omitempty"`
	Timeout time.Duration `json:"timeout,omitempty"`
}

// RPCServer is RPC server config.
type GRPCServer struct {
	Addr                 string        `json:"addr,omitempty"`
	Timeout              time.Duration `json:"timeout,omitempty"`
	IdleTimeout          time.Duration `json:"idleTimeout,omitempty"`
	MaxLifeTime          time.Duration `json:"maxLifeTime,omitempty"`
	ForceCloseWait       time.Duration `json:"forceCloseWait,omitempty"`
	KeepAliveInterval    time.Duration `json:"keepAliveInterval,omitempty"`
	KeepAliveTimeout     time.Duration `json:"keepAliveTimeout,omitempty"`
	MaxMessageSize       int32         `json:"maxMessageSize,omitempty"`
	MaxConcurrentStreams int32         `json:"maxConcurrentStreams,omitempty"`
	EnableOpenTracing    bool          `json:"enableOpenTracing,omitempty"`
}

type BaseConfig struct {
	ProfPathPrefix string   `json:"profPathPrefix,omitempty"`
	LogLevel       string   `json:"logLevel,omitempty"`
	ProfEnable     bool     `json:"profEnable,omitempty"`
	TracerEnable   bool     `json:"tracerEnable,omitempty"`
	WhiteList      []string `json:"whiteList,omitempty"`
	Endpoints      []string `json:"endpoints,omitempty"`
	BaseConfig     string   `json:"baseConfig,omitempty"`
}
