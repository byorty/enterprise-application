package grpc_option

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/auth"
	"github.com/byorty/enterprise-application/pkg/common/adapter/ctxutil"
	"github.com/byorty/enterprise-application/pkg/common/adapter/jwtutil"
	"github.com/byorty/enterprise-application/pkg/common/adapter/log"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	gRPC "google.golang.org/grpc"
)

func NewFxAuthOption(
	logger log.Logger,
	jwtHelper jwtutil.Helper,
) grpc.MiddlewareOut {
	return grpc.MiddlewareOut{
		GrpcMiddleware: grpc.Middleware{
			Priority: 99,
			GrpcOption: func(ctx context.Context, req interface{}, info *gRPC.UnaryServerInfo, handler gRPC.UnaryHandler) (resp interface{}, err error) {
				l := logger.WithCtx(ctx, "middleware", "authentication")
				token, err := grpc_auth.AuthFromMD(ctx, "bearer")

				if len(token) == 0 {
					ctx = context.WithValue(ctx, "session", pbv1.Session{
						Group: pbv1.UserGroupGuest,
					})
					l.Debug("guest session")
				} else {
					claims := new(auth.SessionClaims)
					err = jwtHelper.Parse(token, claims)
					if err != nil {
						l.Error(err)
						return nil, grpc.ErrUnauthenticated(grpc.ErrSessionNotFound)
					}

					ctx = ctxutil.Set(ctx, ctxutil.Session, claims.Session)
					l.Debugf("b2c session=%s with group=%s", claims.Session.Uuid, claims.Session.Group)
				}

				return handler(ctx, req)
			},
		},
	}
}
