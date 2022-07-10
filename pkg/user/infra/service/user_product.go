package usersrvimpl

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/collection"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	productsrv "github.com/byorty/enterprise-application/pkg/product/domain/service"
	usersrv "github.com/byorty/enterprise-application/pkg/user/domain/service"
	"github.com/google/uuid"
	"sort"
)

func NewFxUserProductService(
	userService usersrv.UserService,
	productService productsrv.ProductService,
) usersrv.UserProductService {
	return &userProductService{
		userService:    userService,
		productService: productService,
		userProducts:   collection.NewMap[string, collection.Slice[*pbv1.UserProduct]](),
	}
}

type userProductService struct {
	userService    usersrv.UserService
	productService productsrv.ProductService
	userProducts   collection.Map[string, collection.Slice[*pbv1.UserProduct]]
}

func (s *userProductService) GetAllByFilter(ctx context.Context, params pbv1.GetUserProductRequestParams) ([]*pbv1.UserProduct, error) {
	userProducts := collection.NewSlice[*pbv1.UserProduct]()
	for userUUID, products := range s.userProducts.Entries() {
		if len(params.UserUuidIn) > 0 {
			x := sort.SearchStrings(params.UserUuidIn, userUUID)
			if (len(params.UserUuidIn) > x && params.UserUuidIn[x] == userUUID) == false {
				continue
			}
		}

		for _, product := range products.Entries() {
			if len(params.UuidIn) > 0 {
				y := sort.SearchStrings(params.UuidIn, product.Uuid)
				if len(params.UuidIn) > y && params.UuidIn[y] == product.Uuid {
					userProducts.Add(product)
				}
			} else {
				userProducts.Add(product)
			}
		}
	}

	if userProducts.Len() == 0 {
		return nil, usersrv.ErrUserProductNotFound
	}

	return userProducts.Entries(), nil
}

func (s *userProductService) Put(ctx context.Context, userUUID string, params *pbv1.PutProductRequestParams) ([]*pbv1.UserProduct, error) {
	_, err := s.userService.GetByUUID(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	userProducts, ok := s.userProducts.Get(userUUID)
	if !ok {
		userProducts = collection.NewSlice[*pbv1.UserProduct]()
		s.userProducts.Set(userUUID, userProducts)
	}

	_, err = s.productService.GetByUUID(ctx, params.ProductUuid)
	if err != nil {
		return nil, err
	}

	userProducts.Add(&pbv1.UserProduct{
		Uuid:        uuid.NewString(),
		ProductUuid: params.ProductUuid,
		Count:       params.ProductCount,
	})

	return userProducts.Entries(), nil
}

func (s *userProductService) Change(ctx context.Context, userUUID string, userProductUUID string, params *pbv1.ChangeProductRequestParams) ([]*pbv1.UserProduct, error) {
	_, err := s.userService.GetByUUID(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	userProducts, ok := s.userProducts.Get(userUUID)
	if !ok {
		return nil, usersrv.ErrUserProductNotFound
	}

	for i, userProduct := range userProducts.Entries() {
		if userProduct.Uuid == userProductUUID {
			userProduct.Count = params.Count

			if userProduct.Count <= 0 {
				userProducts.Remove(i)
			}
		}
	}

	return userProducts.Entries(), nil
}
