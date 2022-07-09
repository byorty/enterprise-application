package middleware_test

import (
	"context"
	"errors"
	"github.com/Pallinder/go-randomdata"
	"github.com/byorty/enterprise-application/pkg/common/adapter/auth"
	"github.com/byorty/enterprise-application/pkg/common/adapter/ctxutil"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	usersrv "github.com/byorty/enterprise-application/pkg/user/domain/service"
	"github.com/byorty/enterprise-application/pkg/user/infra/middleware"
	usersrvimpl "github.com/byorty/enterprise-application/pkg/user/infra/service"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/reflect/protoreflect"
	"testing"
)

func TestUserRightsEnforcerSuiteSuite(t *testing.T) {
	suite.Run(t, new(UserRightsEnforcerSuite))
}

type UserRightsEnforcerSuite struct {
	suite.Suite
	enforcer    auth.RightsEnforcer
	userService usersrv.UserService
}

func (s *UserRightsEnforcerSuite) SetupSuite() {
	s.userService = usersrvimpl.NewFxUserService()
	s.enforcer = middleware.NewUserRightsEnforcer(s.userService)
}

func (s *UserRightsEnforcerSuite) TestEnforce() {
	ctx := context.Background()
	userUUID := "387301f4-551c-4022-900a-80f6f76f3a10"
	user, err := s.userService.GetByUUID(ctx, userUUID)
	s.Equal(userUUID, user.Uuid)
	s.Nil(err)

	var session pbv1.Session
	_, err = s.enforcer.Enforce(ctx, session, protoreflect.ValueOf(randomdata.Alphanumeric(32)))
	s.NotNil(err)
	s.True(errors.Is(err, usersrv.ErrUserNotFound))

	user.Status = pbv1.UserStatusBlocked
	_, err = s.enforcer.Enforce(ctx, session, protoreflect.ValueOf(userUUID))
	s.NotNil(err)
	s.True(errors.Is(err, grpc.ErrSessionHasNotPermissions))

	user.Status = pbv1.UserStatusActive
	user.Uuid = randomdata.Alphanumeric(32)
	_, err = s.enforcer.Enforce(ctx, session, protoreflect.ValueOf(userUUID))
	s.NotNil(err)
	s.True(errors.Is(err, grpc.ErrSessionNotOwnEntity))

	session.Uuid = userUUID
	user.Uuid = userUUID
	ctx, err = s.enforcer.Enforce(ctx, session, protoreflect.ValueOf(userUUID))
	s.Nil(err)

	ctxUser, err := ctxutil.Get[*pbv1.User](ctx, ctxutil.User)
	s.Nil(err)
	s.Equal(userUUID, ctxUser.Uuid)
}
