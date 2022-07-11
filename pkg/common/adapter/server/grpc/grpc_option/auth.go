package grpc_option

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/auth"
	"github.com/byorty/enterprise-application/pkg/common/adapter/ctxutil"
	"github.com/byorty/enterprise-application/pkg/common/adapter/log"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	gRPC "google.golang.org/grpc"
)

func NewFxAuthOption(
	logger log.Logger,
	sessionManager auth.SessionManager,
) grpc.MiddlewareOut {
	return grpc.MiddlewareOut{
		GrpcMiddleware: grpc.Middleware{
			Priority: 99,
			GrpcOption: func(ctx context.Context, req interface{}, info *gRPC.UnaryServerInfo, handler gRPC.UnaryHandler) (resp interface{}, err error) {
				logger := logger.WithCtx(ctx, "middleware", "auth")
				token, err := grpc_auth.AuthFromMD(ctx, "bearer")
				if err != nil {
					logger.Error(err)
				}

				session, err := sessionManager.GetSessionByToken(ctx, token)
				if err != nil {
					logger.Error(err)
					return nil, err
				}

				return handler(ctxutil.Set(ctx, ctxutil.Session, session), req)
			},
		},
	}
}
