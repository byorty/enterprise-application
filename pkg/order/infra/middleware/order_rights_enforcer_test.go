package middleware_test

import (
	"context"
	"github.com/Pallinder/go-randomdata"
	"github.com/byorty/enterprise-application/pkg/common/adapter/auth"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	ordersrv "github.com/byorty/enterprise-application/pkg/order/domain/service"
	"github.com/byorty/enterprise-application/pkg/order/infra/middleware"
	ordersrvimpl "github.com/byorty/enterprise-application/pkg/order/infra/service"
	productsrcimpl "github.com/byorty/enterprise-application/pkg/product/infra/service"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
)

func TestOrderRightsEnforcerSuite(t *testing.T) {
	suite.Run(t, new(OrderRightsEnforcerSuite))
}

type OrderRightsEnforcerSuite struct {
	suite.Suite
	enforcer     auth.RightsEnforcer
	orderService ordersrv.OrderService
}

func (s *OrderRightsEnforcerSuite) SetupSuite() {
	productService := productsrcimpl.NewFxProductService()
	s.orderService = ordersrvimpl.NewFxOrderService(productService)
	s.enforcer = middleware.NewOrderRightsEnforcer(s.orderService)
}

func (s *OrderRightsEnforcerSuite) TestEnforce() {
	ctx := context.Background()
	userUUID := "387301f4-551c-4022-900a-80f6f76f3a10"
	order, err := s.orderService.Create(ctx, userUUID, &pbv1.CreateOrderRequestParams{
		Products: []*pbv1.UserProduct{
			{
				Uuid:        uuid.NewString(),
				ProductUuid: "42d8d533-5041-4931-a8c1-f215ab69ffe7",
				Count:       1,
			},
		},
		Address:     randomdata.Address(),
		DeliveredAt: timestamppb.Now(),
	})
	s.NotNil(order)
	s.Nil(err)

	var session pbv1.Session
	_, err = s.enforcer.Enforce(ctx, session, protoreflect.ValueOf(randomdata.Alphanumeric(32)))
	s.True(errors.Is(err, grpc.ErrSessionNotOwnEntity))

	_, err = s.enforcer.Enforce(ctx, session, protoreflect.ValueOf(order.Uuid))
	s.True(errors.Is(err, grpc.ErrSessionNotOwnEntity))

	session.Uuid = userUUID
	_, err = s.enforcer.Enforce(ctx, session, protoreflect.ValueOf(randomdata.Alphanumeric(32)))
	s.True(errors.Is(err, grpc.ErrSessionNotOwnEntity))

	_, err = s.enforcer.Enforce(ctx, session, protoreflect.ValueOf(order.Uuid))
	s.Nil(err)
}
