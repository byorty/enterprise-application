package order

import (
	"context"
	"github.com/Pallinder/go-randomdata"
	pbv1 "github.com/byorty/enterprise-application/gen/proto/v1"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

var _ pbv1.OrderServiceServer = (*server)(nil)

func NewOrderServiceServer() pbv1.OrderServiceServer {
	return new(server)
}

type server struct {
}

func (s *server) Create(ctx context.Context, request *pbv1.CreateOrderRequest) (*pbv1.Order, error) {
	return &pbv1.Order{
		Uuid:        uuid.New().String(),
		Amount:      randomdata.Decimal(100, 1000),
		Address:     randomdata.Address(),
		Status:      pbv1.OrderStatusCreated,
		Products:    nil,
		CreatedAt:   timestamppb.New(time.Now()),
		DeliveredAt: timestamppb.New(time.Now().Add(time.Hour * 24 * 3)),
	}, nil
}

func (s *server) Checkout(ctx context.Context, request *pbv1.CheckoutOrderRequest) (*pbv1.Order, error) {
	return &pbv1.Order{
		Uuid:        uuid.New().String(),
		Amount:      randomdata.Decimal(100, 1000),
		Address:     randomdata.Address(),
		Status:      pbv1.OrderStatusCreated,
		Products:    nil,
		CreatedAt:   timestamppb.New(time.Now()),
		DeliveredAt: timestamppb.New(time.Now().Add(time.Hour * 24 * 3)),
	}, nil
}
