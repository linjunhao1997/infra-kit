package main

import (
	"context"
	"fmt"
	"infra-kit/apps/iam/pb"
	gateway "infra-kit/gateway/option"
	"log"
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
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(gateway.IncommingHeaderMatcher),
		runtime.WithOutgoingHeaderMatcher(gateway.OutgoingHeaderMatcher),
		runtime.WithMetadata(gateway.WithMetadata),
	)
	logger.Info("gateway starting", slog.Group("hello", slog.String("1", "2")))
	basicEndpoint := buildEndpoint(consulAddr, serviceNameBasic, consulTag)
	fmt.Println(basicEndpoint)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	conn, err := grpc.DialContext(ctx, basicEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	authClient := pb.NewIAMServiceClient(conn)
	if err != nil {
		panic(err)
	}

	dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithChainUnaryInterceptor(gateway.UnaryClientAuthInterceptor(authClient)),
		grpc.WithChainStreamInterceptor(gateway.StreamClientAuthInterceptor(authClient))}

	go func() {
		// IAMService
		ctx := context.TODO()
		if err := pb.RegisterIAMServiceHandlerFromEndpoint(ctx, mux, basicEndpoint, dialOptions); err != nil {
			log.Fatal(err)
		}
	}()

	http.ListenAndServe(":"+port, mux)
}

func buildEndpoint(consulAddr, endpoint, tag string) string {
	return fmt.Sprintf("consul://%s/%s?wait=14s&tag=%s", consulAddr, endpoint, tag)
}
