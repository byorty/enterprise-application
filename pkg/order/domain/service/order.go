package ordersrv

import (
	"context"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
)

type OrderService interface {
	Create(ctx context.Context, userUUID string, params *pbv1.CreateOrderRequestParams) (*pbv1.Order, error)
	Checkout(ctx context.Context, orderUUID string, params *pbv1.CheckoutOrderRequestParams) (*pbv1.Order, error)
	GetByUUID(ctx context.Context, orderUUID string) (*pbv1.Order, error)
}
