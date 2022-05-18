package productsrcimpl

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/collection"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	productsrv "github.com/byorty/enterprise-application/pkg/product/domain/service"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
)

func NewProductService() productsrv.ProductService {
	return &productService{
		products: collection.ImportMap[string, *pbv1.Product](map[string]*pbv1.Product{
			"42d8d533-5041-4931-a8c1-f215ab69ffe7": {
				Uuid:         "42d8d533-5041-4931-a8c1-f215ab69ffe7",
				Status:       pbv1.ProductStatusActive,
				Name:         "iPhone 13 Pro 256GB",
				Price:        107000,
				Availability: true,
				CreatedAt:    timestamppb.Now(),
				Properties: []*pbv1.ProductProperty{
					{
						Name:  "Цвет товара",
						Value: "Графитовый",
					},
					{
						Name:  "Процессор",
						Value: "Apple A15 Bionic",
					},
					{
						Name:  "Встроенная память",
						Value: "256 ГБ",
					},
				},
			},
			"d51fcd7e-3899-4592-a7cf-06a04a623ed3": {
				Uuid:         "d51fcd7e-3899-4592-a7cf-06a04a623ed3",
				Status:       pbv1.ProductStatusActive,
				Name:         "iPhone 13 128GB",
				Price:        98000,
				Availability: true,
				CreatedAt:    timestamppb.Now(),
				Properties: []*pbv1.ProductProperty{
					{
						Name:  "Цвет товара",
						Value: "Сияющая звезда",
					},
					{
						Name:  "Процессор",
						Value: "Apple A15 Bionic",
					},
					{
						Name:  "Встроенная память",
						Value: "128 ГБ",
					},
				},
			},
			"9dcd0f0b-79c8-4649-b91f-9a216260c36f": {
				Uuid:         "9dcd0f0b-79c8-4649-b91f-9a216260c36f",
				Status:       pbv1.ProductStatusActive,
				Name:         "Apple Watch Series 7 45 мм Aluminium Case",
				Price:        42490,
				Availability: true,
				CreatedAt:    timestamppb.Now(),
				Properties: []*pbv1.ProductProperty{
					{
						Name:  "Цвет товара",
						Value: "Темная ночь",
					},
					{
						Name:  "Материал корпуса",
						Value: "Алюминий",
					},
					{
						Name:  "Размер корпуса",
						Value: "45 мм",
					},
				},
			},
		}),
	}
}

type productService struct {
	products collection.Map[string, *pbv1.Product]
}

func (s *productService) GetAllByFilter(ctx context.Context, params *pbv1.ProductsRequestFilter, paginator *pbv1.Paginator) ([]*pbv1.Product, uint32, error) {
	products := collection.NewSlice[*pbv1.Product]()
	for _, product := range s.products.Entries() {
		if params != nil {
			if params.PriceLt > 0 && product.Price > params.PriceLt {
				continue
			}

			if params.PriceGt > 0 && product.Price < params.PriceGt {
				continue
			}

			if len(params.NameContains) > 0 && !strings.Contains(product.Name, params.NameContains) {
				continue
			}
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
