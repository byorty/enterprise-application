package grpc_option

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	gRPC "google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewFxValidatorOption(
	methodDescriptorMap grpc.MethodDescriptorMap,
) grpc.MiddlewareOut {
	return grpc.MiddlewareOut{
		GrpcMiddleware: grpc.Middleware{
			Priority: 97,
			GrpcOption: func(ctx context.Context, req interface{}, info *gRPC.UnaryServerInfo, handler gRPC.UnaryHandler) (resp interface{}, err error) {
				protoMessage, ok := req.(protoreflect.ProtoMessage)
				if !ok {
					return handler(ctx, req)
				}

				methodDescriptor, ok := methodDescriptorMap.GetByFullName(info.FullMethod)
				if !ok {
					return handler(ctx, req)
				}

				if len(methodDescriptor.Form) == 0 {
					return handler(ctx, req)
				}

				err = methodDescriptor.Form.Validate(ctx, protoMessage.ProtoReflect())
				if err != nil {
					return nil, grpc.ErrInvalidArgument(err)
				}

				return handler(ctx, req)
			},
		},
	}
}
