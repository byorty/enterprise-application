package orderapp

import (
	"context"
	ordersrv "github.com/byorty/enterprise-application/internal/order/domain/service"
	pbv1 "github.com/byorty/enterprise-application/internal/pkg/gen/api/proto/v1"
)

var _ pbv1.OrderServiceServer = (*server)(nil)

func NewOrderServiceServer(
	orderService ordersrv.OrderService,
) pbv1.OrderServiceServer {
	return &server{
		orderService: orderService,
	}
}

type server struct {
	orderService ordersrv.OrderService
}

func (s *server) Create(ctx context.Context, request *pbv1.CreateOrderRequest) (*pbv1.Order, error) {
	return s.orderService.Create(ctx, request)
}

func (s *server) Checkout(ctx context.Context, request *pbv1.CheckoutOrderRequest) (*pbv1.Order, error) {
	return s.orderService.Checkout(ctx, request.OrderUuid, request.Params)
}
