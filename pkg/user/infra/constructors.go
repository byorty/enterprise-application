package useradap

import (
	userapp "github.com/byorty/enterprise-application/pkg/user/infra/app"
	"github.com/byorty/enterprise-application/pkg/user/infra/middleware"
	usersrvimpl "github.com/byorty/enterprise-application/pkg/user/infra/service"
	"go.uber.org/fx"
)

var Constructors = fx.Provide(
	userapp.NewFxUserServiceServer,
	usersrvimpl.NewFxUserService,
	usersrvimpl.NewFxUserProductService,
	middleware.NewFxUserRightsEnforcer,
	middleware.NewFxUserProductRightsEnforcer,
)
