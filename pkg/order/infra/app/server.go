package orderapp

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	ordersrv "github.com/byorty/enterprise-application/pkg/order/domain/service"
)

var _ pbv1.OrderServiceServer = (*server)(nil)

func NewFxOrderServiceServer(
	orderService ordersrv.OrderService,
) grpc.Descriptor {
	return grpc.Descriptor{
		Server: &server{
			orderService: orderService,
		},
		GRPCRegistrar:        pbv1.RegisterOrderServiceServer,
		GRPCGatewayRegistrar: pbv1.RegisterOrderServiceHandlerFromEndpoint,
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
