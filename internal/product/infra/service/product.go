package productsrcimpl

import (
	"context"
	"github.com/Pallinder/go-randomdata"
	"github.com/byorty/enterprise-application/internal/pkg/collection"
	pbv1 "github.com/byorty/enterprise-application/internal/pkg/gen/api/proto/v1"
	productsrv "github.com/byorty/enterprise-application/internal/product/domain/service"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"time"
)

func NewProductService() productsrv.ProductService {
	products := collection.NewMap[string, *pbv1.Product]()
	for i := 0; i < 10; i++ {
		product := &pbv1.Product{
			Uuid:         uuid.NewString(),
			Status:       pbv1.ProductStatusActive,
			Name:         randomdata.Alphanumeric(16),
			Description:  randomdata.Alphanumeric(64),
			Price:        randomdata.Decimal(10, 1000),
			Availability: randomdata.Boolean(),
			CreatedAt:    timestamppb.New(time.Now()),
			Properties: []*pbv1.ProductProperty{
				{
					Name:  randomdata.Alphanumeric(16),
					Value: randomdata.Alphanumeric(32),
				},
				{
					Name:  randomdata.Alphanumeric(16),
					Value: randomdata.Alphanumeric(32),
				},
				{
					Name:  randomdata.Alphanumeric(16),
					Value: randomdata.Alphanumeric(32),
				},
			},
		}

		products.Set(product.Uuid, product)
	}

	return &productService{
		products: products,
	}
}

type productService struct {
	products collection.Map[string, *pbv1.Product]
}

func (s *productService) GetAllByFilter(ctx context.Context, params *pbv1.ProductsRequestFilter, paginator *pbv1.Paginator) ([]*pbv1.Product, uint32, error) {
	products := collection.NewList[*pbv1.Product]()
	for _, product := range s.products.Entries() {
		if params.PriceLt > 0 && product.Price > params.PriceLt {
			continue
		}

		if params.PriceGt > 0 && product.Price < params.PriceGt {
			continue
		}

		if len(params.NameContains) > 0 && !strings.Contains(product.Name, params.NameContains) {
			continue
		}

		products.Add(product)
	}

	return products.Entries(), uint32(products.Len()), nil
}

func (s *productService) GetByUUID(ctx context.Context, uuid string) (*pbv1.Product, error) {
	product, ok := s.products.Get(uuid)
	if !ok {
		return nil, productsrv.ErrProductNotFound
	}

	return product, nil
}
