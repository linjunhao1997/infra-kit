package grpclib

import "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

var (
	StatsClientHandler = otelgrpc.NewClientHandler()
	StatsServerHandler = otelgrpc.NewServerHandler()
)
