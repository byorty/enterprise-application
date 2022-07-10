package middleware_test

import (
	"context"
	"github.com/Pallinder/go-randomdata"
	"github.com/byorty/enterprise-application/pkg/common/adapter/auth"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	productsrcimpl "github.com/byorty/enterprise-application/pkg/product/infra/service"
	usersrv "github.com/byorty/enterprise-application/pkg/user/domain/service"
	"github.com/byorty/enterprise-application/pkg/user/infra/middleware"
	usersrvimpl "github.com/byorty/enterprise-application/pkg/user/infra/service"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/reflect/protoreflect"
	"testing"
)

func TestUserProductRightsEnforcerSuite(t *testing.T) {
	suite.Run(t, new(UserProductRightsEnforcerSuite))
}

type UserProductRightsEnforcerSuite struct {
	suite.Suite
	enforcer           auth.RightsEnforcer
	userProductService usersrv.UserProductService
}

func (s *UserProductRightsEnforcerSuite) SetupSuite() {
	productService := productsrcimpl.NewFxProductService()
	userService := usersrvimpl.NewFxUserService()
	s.userProductService = usersrvimpl.NewFxUserProductService(userService, productService)
	s.enforcer = middleware.NewUserProductRightsEnforcer(s.userProductService)
}

func (s *UserProductRightsEnforcerSuite) TestEnforce() {
	var session pbv1.Session
	ctx := context.Background()
	_, err := s.enforcer.Enforce(ctx, session, protoreflect.ValueOf(randomdata.Alphanumeric(32)))
	s.True(errors.Is(err, grpc.ErrSessionNotOwnEntity))

	session.Uuid = "387301f4-551c-4022-900a-80f6f76f3a10"
	_, err = s.enforcer.Enforce(ctx, session, protoreflect.ValueOf(randomdata.Alphanumeric(32)))
	s.True(errors.Is(err, grpc.ErrSessionNotOwnEntity))

	productUuid := "42d8d533-5041-4931-a8c1-f215ab69ffe7"
	userProducts, err := s.userProductService.Put(ctx, session.Uuid, &pbv1.PutProductRequestParams{
		ProductUuid:  productUuid,
		ProductCount: 1,
	})
	s.NotEmpty(userProducts)
	s.Nil(err)

	_, err = s.enforcer.Enforce(ctx, session, protoreflect.ValueOf(randomdata.Alphanumeric(32)))
	s.True(errors.Is(err, grpc.ErrSessionNotOwnEntity))

	_, err = s.enforcer.Enforce(ctx, session, protoreflect.ValueOf(userProducts[0].Uuid))
	s.Nil(err)
}
