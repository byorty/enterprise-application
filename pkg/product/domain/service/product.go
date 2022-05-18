package productsrv

import (
	"context"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
)

type ProductService interface {
	GetAllByFilter(ctx context.Context, params *pbv1.ProductsRequestFilter, paginator *pbv1.Paginator) ([]*pbv1.Product, uint32, error)
	GetByUUID(ctx context.Context, uuid string) (*pbv1.Product, error)
}
