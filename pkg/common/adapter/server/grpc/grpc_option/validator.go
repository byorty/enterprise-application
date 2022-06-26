package grpc_option

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/log"
	"github.com/byorty/enterprise-application/pkg/common/adapter/protoutil"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	"github.com/byorty/enterprise-application/pkg/common/adapter/validator"
	gRPC "google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewFxValidatorOption(
	logger log.Logger,
	validators *protoutil.Map[validator.Validator],
) grpc.MiddlewareOut {
	return grpc.MiddlewareOut{
		GrpcMiddleware: grpc.Middleware{
			Priority: 97,
			GrpcOption: func(ctx context.Context, req interface{}, info *gRPC.UnaryServerInfo, handler gRPC.UnaryHandler) (resp interface{}, err error) {
				l := logger.WithCtx(ctx, "middleware", "validator")
				protoMessage, ok := req.(protoreflect.ProtoMessage)
				if ok {
					message := protoMessage.ProtoReflect()
					fields := message.Descriptor().Fields()
					for i := 0; i < fields.Len(); i++ {
						field := fields.Get(i)
						vld, err := validators.Get(message, field)
						if err != nil {
							continue
						}

						err = vld.Validate(ctx, message, field)
						if err != nil {
							l.Error(err)
							return nil, grpc.ErrPermissionDenied(grpc.ErrSessionNotOwnEntity)
						}
					}
				}

				return handler(ctx, req)
			},
		},
	}
}
