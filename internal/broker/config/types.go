package config

import (
	xconfig "github.com/imneov/modelmesh/pkg/config"
)

// RPCClient is RPC client config.
type GRPCClient = xconfig.GRPCClient

// RPCServer is RPC server config.
type GRPCServer = xconfig.GRPCServer

type BaseConfig = xconfig.BaseConfig
type Schedule = xconfig.Schedule
type Queue = xconfig.Queue
type Dispatch = xconfig.Dispatch
type ServiceGroup = xconfig.ServiceGroup

// Config defines everything needed for apiserver to deal with external services
type Config struct {
	BaseOptions   *BaseConfig     `yaml:"base" json:"baseOptions,omitempty"`
	BrokerServer  *GRPCServer     `yaml:"brokerServer" json:"brokerServer,omitempty"`
	Schedule      *Schedule       `yaml:"schedule" json:"schedule,omitempty"`
	Queue         *Queue          `yaml:"queue" json:"queue,omitempty"`
	Dispatch      *Dispatch       `yaml:"dispatch" json:"dispatch,omitempty"`
	ServiceGroups []*ServiceGroup `yaml:"serviceGroups" json:"serviceGroups,omitempty"`
}
