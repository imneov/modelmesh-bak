package xconfig

import (
	"github.com/imneov/modelmesh/internal/broker/picker"
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

// ServiceQueue 封装底层 ServingServer 并提供 ServiceQueue 抽象，一个 ServiceQueue 可以设置对应的调度策略
type Schedule struct {
	Method picker.SchedulingMethod
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

func (sg *ServiceGroup) ResourceName() string {
	return sg.Name
}

type Queue struct {
	Size int64 // 服务队列大小
}

// ServiceQueue 封装底层 ServingServer 并提供 ServiceQueue 抽象，一个 ServiceQueue 可以设置对应的调度策略
type Dispatch struct {
	Queue  *Queue      `json:"queue,omitempty"` // 等待队列
	Client *GRPCClient `json:"client,omitempty"`
}
