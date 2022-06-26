package grpc_option

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/auth"
	"github.com/byorty/enterprise-application/pkg/common/adapter/ctxutil"
	"github.com/byorty/enterprise-application/pkg/common/adapter/log"
	"github.com/byorty/enterprise-application/pkg/common/adapter/protoutil"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	gRPC "google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewFxEnforcerOption(
	logger log.Logger,
	roleEnforcer auth.RoleEnforcer,
	methodDescriptorMap grpc.MethodDescriptorMap,
	rightsEnforcerDescriptors []auth.RightsEnforcerDescriptor,
) grpc.MiddlewareOut {
	rightsEnforcers := protoutil.NewMap[auth.RightsEnforcer]()
	for _, descriptor := range rightsEnforcerDescriptors {
		rightsEnforcers.Set(descriptor.Name, descriptor.RightsEnforcer)
	}

	return grpc.MiddlewareOut{
		GrpcMiddleware: grpc.Middleware{
			Priority: 98,
			GrpcOption: func(ctx context.Context, req interface{}, info *gRPC.UnaryServerInfo, handler gRPC.UnaryHandler) (resp interface{}, err error) {
				l := logger.WithCtx(ctx, "middleware", "enforcer")
				methodDescriptor, ok := methodDescriptorMap[info.FullMethod]
				if !ok {
					return nil, grpc.ErrUnauthenticated(grpc.ErrMethodDescriptorNotFound)
				}

				session, err := ctxutil.Get[pbv1.Session](ctx, ctxutil.Session)
				if err != nil {
					return nil, grpc.ErrPermissionDenied(grpc.ErrSessionNotFound)
				}

				ok, err = roleEnforcer.Enforce(session, methodDescriptor.Object, methodDescriptor.Operation)
				if !ok {
					l.Error(err)
					return nil, grpc.ErrPermissionDenied(grpc.ErrSessionHasNotPermissions)
				}

				protoMessage, ok := req.(protoreflect.ProtoMessage)
				if ok {
					message := protoMessage.ProtoReflect()
					fields := message.Descriptor().Fields()
					for i := 0; i < fields.Len(); i++ {
						field := fields.Get(i)
						rightsEnforcer, err := rightsEnforcers.Get(message, field)
						if err != nil {
							continue
						}

						ctx, err = rightsEnforcer.Enforce(ctx, session, message.Get(field))
						if err != nil {
							return nil, grpc.ErrPermissionDenied(grpc.ErrSessionNotOwnEntity)
						}
					}
				}

				return handler(ctx, req)
			},
		},
	}
}
