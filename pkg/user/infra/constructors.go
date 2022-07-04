package useradap

import (
	"github.com/byorty/enterprise-application/pkg/user/infra/middleware"
	usersrvimpl "github.com/byorty/enterprise-application/pkg/user/infra/service"
	"go.uber.org/fx"
)

var Constructors = fx.Provide(
	usersrvimpl.NewFxUserService,
	usersrvimpl.NewFxUserProductService,
	middleware.NewFxUserRightsEnforcer,
)
