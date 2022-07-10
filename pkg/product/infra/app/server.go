package productapp

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	productsrv "github.com/byorty/enterprise-application/pkg/product/domain/service"
)

var _ pbv1.ProductServiceServer = (*server)(nil)

func NewFxProductServiceServer(
	productService productsrv.ProductService,
) grpc.Descriptor {
	return grpc.Descriptor{
		Server: &server{
			productService: productService,
		},
		GRPCRegistrar:        pbv1.RegisterProductServiceServer,
		GRPCGatewayRegistrar: pbv1.RegisterProductServiceHandlerFromEndpoint,
		MethodDescriptors: []grpc.MethodDescriptor{
			{
				Method:     (*server).GetAllByFilter,
				Role:       pbv1.RoleProduct,
				Permission: pbv1.PermissionRead,
			},
			{
				Method:     (*server).GetByUUID,
				Role:       pbv1.RoleProduct,
				Permission: pbv1.PermissionRead,
			},
		},
	}
}

type server struct {
	productService productsrv.ProductService
}

func (s server) GetAllByFilter(ctx context.Context, request *pbv1.ProductsRequest) (*pbv1.ProductResponse, error) {
	products, count, err := s.productService.GetAllByFilter(ctx, request.Filter, request.Paginator)
	return &pbv1.ProductResponse{
		Items: products,
		Count: count,
	}, err
}

func (s server) GetByUUID(ctx context.Context, request *pbv1.GetByProductUUIDRequest) (*pbv1.Product, error) {
	return s.productService.GetByUUID(ctx, request.ProductUuid)
}
