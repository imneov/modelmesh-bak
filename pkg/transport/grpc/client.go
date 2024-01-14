/*
 * Copyright (C) 2019 Yunify, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this work except in compliance with the License.
 * You may obtain a copy of the License in the LICENSE file, or at:
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package grpc

import (
	"context"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/imneov/modelmesh/internal/broker/config"
	proto "github.com/imneov/modelmesh/mindspore_serving/proto"
	"google.golang.org/grpc"
	"k8s.io/klog/v2"
	"time"
)

var opts = []grpc.DialOption{
	grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	grpc.WithInsecure(),
	grpc.WithUnaryInterceptor(
		grpc_opentracing.UnaryClientInterceptor(),
	),
	//grpc.WithTimeout(timeout),
	grpc.WithBackoffMaxDelay(time.Second * 3),
	grpc.WithInitialWindowSize(1 << 30),
	grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(64 * 1024 * 1024)),
}

func NewServingClient(conf *config.GRPCClient) proto.MSServiceClient {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(conf.Timeout))
	conn, err := grpc.DialContext(ctx, conf.Addr, opts...)
	if err != nil {
		klog.Errorf("connect fail, endpoint: %s, err %w", "endpoint", conf.Addr, err)
	}
	return proto.NewMSServiceClient(conn)
}
