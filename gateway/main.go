package main

import (
	"context"
	"fmt"
	"infra-kit/apps/iam/pb"
	gateway "infra-kit/gateway/option"
	"infra-kit/lib/grpclib"
	"infra-kit/lib/loglib"
	"infra-kit/lib/metriclib"
	"infra-kit/lib/tracelib"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	//timeout            = 10 * time.Second
	consulAddr       = os.Getenv("consul_addr")
	consulTag        = os.Getenv("consul_tag")
	port             = os.Getenv("port")
	serviceNameBasic = os.Getenv("service_name_basic")
)

func main() {
	loglib.Init()
	err := tracelib.Init("grpc-gateway")
	if err != nil {
		slog.Error("tracelib Init failed", "err", err)
		return
	}
	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(gateway.IncommingHeaderMatcher),
		runtime.WithOutgoingHeaderMatcher(gateway.OutgoingHeaderMatcher),
		runtime.WithMetadata(gateway.WithMetadata),
	)
	iamEndpoint := buildEndpoint(consulAddr, serviceNameBasic, consulTag)
	slog.Info("build iam endpoint", "endpoint", iamEndpoint)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	conn, err := grpc.DialContext(ctx, iamEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	authClient := pb.NewIAMServiceClient(conn)
	if err != nil {
		slog.Error("new iam service client", "error", err)
		return
	}

	dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithChainUnaryInterceptor(gateway.UnaryClientAuthInterceptor(authClient), grpclib.LoggingUnaryClientInterceptor, grpclib.MetricUnaryClientInterceptor),
		grpc.WithChainStreamInterceptor(gateway.StreamClientAuthInterceptor(authClient), grpclib.LoggingUnaryStreamInterceptor, grpclib.MetricStreamClientInterceptor),
		grpc.WithStatsHandler(grpclib.StatsClientHandler),
	}

	if err := pb.RegisterIAMServiceHandlerFromEndpoint(ctx, mux, iamEndpoint, dialOptions); err != nil {
		slog.Error("register iam service handler from endpoint", "error", err)
	}

	mux.HandlePath("GET", "/metrics", MetricHandler)
	slog.Info("http gateway server listening", "port", port)
	slog.Error("listen and server", "err", http.ListenAndServe(":"+port, mux))
}

func buildEndpoint(consulAddr, endpoint, tag string) string {
	return fmt.Sprintf("consul://%s/%s?wait=14s&tag=%s", consulAddr, endpoint, tag)
}

func MetricHandler(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
	metriclib.DefaultMetricsHttpHandler().ServeHTTP(w, req)
}
