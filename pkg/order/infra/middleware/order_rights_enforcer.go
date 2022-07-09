package middleware

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/auth"
	"github.com/byorty/enterprise-application/pkg/common/adapter/ctxutil"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	ordersrv "github.com/byorty/enterprise-application/pkg/order/domain/service"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewFxOrderRightsEnforcer(
	orderService ordersrv.OrderService,
) auth.RightsEnforcerDescriptorOut {
	return auth.RightsEnforcerDescriptorOut{
		Descriptor: auth.RightsEnforcerDescriptor{
			Name: "order_uuid",
			RightsEnforcer: &orderRightsEnforcer{
				orderService: orderService,
			},
		},
	}
}

var _ auth.RightsEnforcer = (*orderRightsEnforcer)(nil)

type orderRightsEnforcer struct {
	orderService ordersrv.OrderService
}

func (r orderRightsEnforcer) Enforce(ctx context.Context, session pbv1.Session, value protoreflect.Value) (context.Context, error) {
	order, err := r.orderService.GetByUUID(ctx, value.String())
	if err != nil {
		return nil, err
	}

	if order.UserUuid != session.Uuid {
		return nil, grpc.ErrSessionNotOwnEntity
	}

	return ctxutil.Set(ctx, ctxutil.Order, order), nil
}
