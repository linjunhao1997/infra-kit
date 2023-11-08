package gateway

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"
)

const (
	metadata_grpc_gateway_http_method = "grpcgateway-http-method"
)

func WithMetadata(_ context.Context, req *http.Request) metadata.MD {
	md := metadata.New(map[string]string{
		metadata_grpc_gateway_http_method: req.Method,
	})
	tokenCookie, _ := req.Cookie(HeaderAccessToken)
	if tokenCookie != nil {
		md.Append(HeaderAccessToken, tokenCookie.Value)
	}
	return md
}
