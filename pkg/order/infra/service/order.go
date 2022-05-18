package ordersrvimpl

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/collection"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	"github.com/byorty/enterprise-application/pkg/order/domain/service"
	productsrv "github.com/byorty/enterprise-application/pkg/product/domain/service"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewOrderService(
	productService productsrv.ProductService,
) ordersrv.OrderService {
	return &orderService{
		orders:         collection.NewMap[string, *pbv1.Order](),
		productService: productService,
	}
}

type orderService struct {
	orders         collection.Map[string, *pbv1.Order]
	productService productsrv.ProductService
}

func (s *orderService) Create(ctx context.Context, request *pbv1.CreateOrderRequest) (*pbv1.Order, error) {
	order := &pbv1.Order{
		Uuid:        uuid.NewString(),
		Address:     request.Address,
		Status:      pbv1.OrderStatusCreated,
		Products:    request.Products,
		CreatedAt:   timestamppb.Now(),
		DeliveredAt: request.DeliveredAt,
	}

	for _, userProduct := range request.Products {
		product, err := s.productService.GetByUUID(ctx, userProduct.ProductUuid)
		if err != nil {
			return nil, err
		}

		order.Amount += product.Price * float64(userProduct.Count)
	}

	s.orders.Set(order.Uuid, order)
	return order, nil
}

func (s *orderService) Checkout(ctx context.Context, orderUUID string, params *pbv1.CheckoutOrderRequestParams) (*pbv1.Order, error) {
	order, ok := s.orders.Get(orderUUID)
	if !ok {
		return nil, ordersrv.ErrOrderNotFound
	}

	if order.Amount != params.Amount {
		return nil, ordersrv.ErrOrderAmountNotEqual
	}

	order.Status = params.Status

	return order, nil
}
