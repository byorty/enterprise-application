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

func NewFxUserProductRightsEnforcer(
	userProductService usersrv.UserProductService,
) auth.RightsEnforcerDescriptorOut {
	return auth.RightsEnforcerDescriptorOut{
		Descriptor: auth.RightsEnforcerDescriptor{
			Name:           "user_product_uuid",
			RightsEnforcer: NewUserProductRightsEnforcer(userProductService),
		},
	}
}

func NewUserProductRightsEnforcer(
	userProductService usersrv.UserProductService,
) auth.RightsEnforcer {
	return &userProductRightsEnforcer{
		userProductService: userProductService,
	}
}

type userProductRightsEnforcer struct {
	userProductService usersrv.UserProductService
}

func (r userProductRightsEnforcer) Enforce(ctx context.Context, session pbv1.Session, value protoreflect.Value) (context.Context, error) {
	userProducts, err := r.userProductService.GetAllByFilter(ctx, pbv1.GetUserProductRequestParams{
		UuidIn:     []string{value.String()},
		UserUuidIn: []string{session.Uuid},
	})
	if err != nil || len(userProducts) == 0 {
		return nil, grpc.ErrSessionNotOwnEntity
	}

	return ctxutil.Set(ctx, ctxutil.UserProduct, userProducts[0]), nil
}
