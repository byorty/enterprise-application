package grpc_option

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	gRPC "google.golang.org/grpc"
)

func NewFxPanicOption() grpc.MiddlewareOut {
	return grpc.MiddlewareOut{
		GrpcMiddleware: grpc.Middleware{
			Priority: int(^uint32(0) >> 1),
			GrpcOption: func(ctx context.Context, req interface{}, info *gRPC.UnaryServerInfo, handler gRPC.UnaryHandler) (resp interface{}, err error) {
				defer func() {
					if r := recover(); r != nil {
						err = grpc.ErrInternal(err)
					}
				}()

				resp, err = handler(ctx, req)
				return resp, err
			},
		},
	}
}
