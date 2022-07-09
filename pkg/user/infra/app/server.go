package userapp

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	ordersrv "github.com/byorty/enterprise-application/pkg/order/domain/service"
	usersrv "github.com/byorty/enterprise-application/pkg/user/domain/service"
)

var _ pbv1.UserServiceServer = (*server)(nil)

func NewFxUserServiceServer(
	userService usersrv.UserService,
	userProductService usersrv.UserProductService,
	orderService ordersrv.OrderService,
) grpc.Descriptor {
	return grpc.Descriptor{
		Server: &server{
			userService:        userService,
			userProductService: userProductService,
			orderService:       orderService,
		},
		GRPCRegistrar:        pbv1.RegisterUserServiceServer,
		GRPCGatewayRegistrar: pbv1.RegisterUserServiceHandlerFromEndpoint,
		MethodDescriptors: []grpc.MethodDescriptor{
			{
				Method:     (*server).Register,
				Role:       pbv1.RoleUser,
				Permission: pbv1.PermissionRead,
			},
			{
				Method:     (*server).GetByUUID,
				Role:       pbv1.RoleUser,
				Permission: pbv1.PermissionRead,
			},
		},
	}
}

type server struct {
	userService        usersrv.UserService
	userProductService usersrv.UserProductService
	orderService       ordersrv.OrderService
}

func (s server) Register(ctx context.Context, request *pbv1.RegisterRequest) (*pbv1.User, error) {
	return s.userService.Register(ctx, request.PhoneNumber)
}

func (s server) GetByUUID(ctx context.Context, request *pbv1.GetByUserUUIDRequest) (*pbv1.User, error) {
	return s.userService.GetByUUID(ctx, request.UserUuid)
}

func (s server) GetUserProducts(ctx context.Context, request *pbv1.GetByUserUUIDRequest) (*pbv1.UserProductsResponse, error) {
	userProducts, err := s.userProductService.GetProductsByUserUUID(ctx, request.UserUuid)
	return &pbv1.UserProductsResponse{
		Products: userProducts,
	}, err
}

func (s server) PutProduct(ctx context.Context, request *pbv1.PutProductRequest) (*pbv1.UserProductsResponse, error) {
	userProducts, err := s.userProductService.Put(ctx, request.UserUuid, request.Params)
	return &pbv1.UserProductsResponse{
		Products: userProducts,
	}, err
}

func (s server) ChangeProduct(ctx context.Context, request *pbv1.ChangeProductRequest) (*pbv1.UserProductsResponse, error) {
	userProducts, err := s.userProductService.Change(ctx, request.UserUuid, request.UserProductUuid, request.Params)
	return &pbv1.UserProductsResponse{
		Products: userProducts,
	}, err
}

func (s *server) CreateOrder(ctx context.Context, request *pbv1.CreateOrderRequest) (*pbv1.Order, error) {
	return s.orderService.Create(ctx, request.UserUuid, request.Params)
}
