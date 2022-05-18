package userapp

import (
	"context"
	pbv1 "github.com/byorty/enterprise-application/internal/pkg/gen/api/proto/v1"
	usersrv "github.com/byorty/enterprise-application/internal/user/domain/service"
)

var _ pbv1.UserServiceServer = (*server)(nil)

func NewServer(
	userService usersrv.UserService,
	userProductService usersrv.UserProductService,
) pbv1.UserServiceServer {
	return &server{
		userService:        userService,
		userProductService: userProductService,
	}
}

type server struct {
	userService        usersrv.UserService
	userProductService usersrv.UserProductService
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
