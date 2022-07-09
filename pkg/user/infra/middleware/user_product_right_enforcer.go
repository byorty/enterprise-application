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
	userService usersrv.UserService,
	userProductService usersrv.UserProductService,
) auth.RightsEnforcerDescriptorOut {
	return auth.RightsEnforcerDescriptorOut{
		Descriptor: auth.RightsEnforcerDescriptor{
			Name: "user_product_uuid",
			RightsEnforcer: &userProductRightsEnforcer{
				userService:        userService,
				userProductService: userProductService,
			},
		},
	}
}

var _ auth.RightsEnforcer = (*userRightsEnforcer)(nil)

type userProductRightsEnforcer struct {
	userService        usersrv.UserService
	userProductService usersrv.UserProductService
}

func (r userProductRightsEnforcer) Enforce(ctx context.Context, session pbv1.Session, value protoreflect.Value) (context.Context, error) {
	userProductUUID := value.String()
	userProducts, err := r.userProductService.GetProductsByUserUUID(ctx, session.Uuid)
	if err != nil {
		return nil, err
	}

	for _, userProduct := range userProducts {
		if userProduct.Uuid == userProductUUID {
			return ctxutil.Set(ctx, ctxutil.UserProduct, userProduct), nil
		}
	}

	return nil, grpc.ErrSessionNotOwnEntity
}
