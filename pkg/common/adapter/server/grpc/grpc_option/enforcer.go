package grpc_option

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/auth"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	gRPC "google.golang.org/grpc"
)

func NewFxEnforcerOption(
	enforcer auth.Enforcer,
) grpc.MiddlewareOut {
	return grpc.MiddlewareOut{
		GrpcMiddleware: grpc.Middleware{
			Priority: 1,
			GrpcOption: func(ctx context.Context, req interface{}, info *gRPC.UnaryServerInfo, handler gRPC.UnaryHandler) (resp interface{}, err error) {
				//enforcer.Enforce()
				return handler(ctx, req)
			},
		},
	}
}
