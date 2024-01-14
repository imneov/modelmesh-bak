package config

import (
	xconfig "github.com/imneov/modelmesh/pkg/config"
)

// RPCClient is RPC client config.
type GRPCClient = xconfig.GRPCClient

// RPCServer is RPC server config.
type GRPCServer = xconfig.GRPCServer

type BaseConfig = xconfig.BaseConfig

type SchedulingMethod string

const (
	// SP (Strict Priority—Traffic) scheduling for the selected queue and all higher queues is based strictly on the queue priority.
	SP SchedulingMethod = "SP"
	// WRR (WRR—Traffic) scheduling for the selected queue is based on WRR. The period time is divided between the WRR queues that are not empty, meaning that they have descriptors to egress. It happens only if Strict Priority queues are empty.
	WRR SchedulingMethod = "WRR"
)

// ServiceQueue 封装底层 ServingServer 并提供 ServiceQueue 抽象，一个 ServiceQueue 可以设置对应的调度策略
type ServiceQueue struct {
	Name             string
	ServingClient    *GRPCClient
	SchedulingMethod SchedulingMethod
}

// ServiceGroup 表示一个服务组，服务组的资源分配基于权重，权重高的服务组可以获得更多的资源，具体调度策略由 ServiceQueue 的 SchedulingMethod 决定
type ServiceGroup struct {
	Name string
	//reclaimable表示该queue在资源使用量超过该queue所应得的资源份额时，是否允许其他queue回收该queue使用超额的资源，默认值为true
	Reclaimable bool
	//weight表示该ServiceQueue在集群资源划分中所占的相对比重
	//该ServiceQueue应得资源总量为 (weight/total-weight) * total-resource。
	//其中， total-weight表示所有的queue的weight总和，total-resource表示集群的资源总量。
	//weight是一个软约束，取值范围为[1, 2^31-1]
	Weight uint16
}
