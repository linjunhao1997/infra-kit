package gateway

import (
	"context"
	"errors"
	"fmt"
	"infra-kit/apps/iam/pb"
	"infra-kit/lib/authlib"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type OrgRequest interface {
	GetOrgCode() string
}

const (
	OrgAdmin = "admin"
)

var (
	errOrgCodeIsInvalid = errors.New("orgCode is invalid")
)

func UnaryClientAuthInterceptor(client pb.IAMServiceClient) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, resp interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if method == pb.IAMService_Auth_FullMethodName {
			return invoker(ctx, method, req, resp, cc, opts...)
		}

		md, _ := metadata.FromOutgoingContext(ctx)

		tokenMD := md.Get(HeaderAccessToken)
		var token string
		if len(tokenMD) > 0 {
			token = tokenMD[0]
		} else {
			return status.Error(codes.Unauthenticated, "auth failed")
		}
		authStatus, err := client.CheckAuthStatus(ctx, &pb.CheckAuthStatusRequest{
			Token: token,
		})
		if err != nil {
			return err
		}

		if req, ok := req.(OrgRequest); ok {
			reqOrgCode := req.GetOrgCode()
			userOrgCode := authStatus.OrgCode
			if userOrgCode != OrgAdmin {
				if reqOrgCode != userOrgCode {
					return status.Error(codes.InvalidArgument, errOrgCodeIsInvalid.Error())
				}
			}
		}
		urlMethod := md[metadata_grpc_gateway_http_method][0]
		path, _ := runtime.HTTPPathPattern(ctx)
		authorityCode := authlib.GetUrlPathAuthorityCode(path, urlMethod)
		res, err := client.ValidateAuthorities(ctx, &pb.ValidateAuthoritiesRequest{
			Token:       token,
			Authorities: []string{authorityCode},
		})
		fmt.Println(err)

		if !res.Result[authorityCode] {
			return status.Error(codes.PermissionDenied, "")
		}

		return invoker(ctx, method, req, resp, cc, opts...)
	}
}

func StreamClientAuthInterceptor(client pb.IAMServiceClient) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		md, _ := metadata.FromOutgoingContext(ctx)
		tokenMD := md.Get(HeaderAccessToken)
		var token string
		if len(tokenMD) > 0 {
			token = tokenMD[0]
		} else {
			return nil, status.Error(codes.Unauthenticated, "auth failed")
		}

		pattern := md["grpcgateway-http-path"][0]
		res, err := client.ValidateAuthorities(ctx, &pb.ValidateAuthoritiesRequest{
			Token:       token,
			Authorities: []string{pattern},
		})
		if err != nil {
			return nil, err
		}
		if !res.Result[pattern] {
			return nil, status.Error(codes.PermissionDenied, "validate authority failed")
		}

		return streamer(ctx, desc, cc, method, opts...)
	}
}
