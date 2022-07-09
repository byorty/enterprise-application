package middleware

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/auth"
	"github.com/byorty/enterprise-application/pkg/common/adapter/ctxutil"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	usersrv "github.com/byorty/enterprise-application/pkg/user/domain/service"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewFxUserRightsEnforcer(
	userService usersrv.UserService,
) auth.RightsEnforcerDescriptorOut {
	return auth.RightsEnforcerDescriptorOut{
		Descriptor: auth.RightsEnforcerDescriptor{
			Name:           "user_uuid",
			RightsEnforcer: NewUserRightsEnforcer(userService),
		},
	}
}

func NewUserRightsEnforcer(
	userService usersrv.UserService,
) auth.RightsEnforcer {
	return &userRightsEnforcer{
		userService: userService,
	}
}

type userRightsEnforcer struct {
	userService usersrv.UserService
}

func (r userRightsEnforcer) Enforce(ctx context.Context, session pbv1.Session, value protoreflect.Value) (context.Context, error) {
	user, err := r.userService.GetByUUID(ctx, value.String())
	if err != nil {
		return nil, err
	}

	if user.Status != pbv1.UserStatusActive {
		return nil, grpc.ErrSessionHasNotPermissions
	}

	if session.Uuid != user.Uuid {
		return nil, grpc.ErrSessionNotOwnEntity
	}

	return ctxutil.Set(ctx, ctxutil.User, user), nil
}
