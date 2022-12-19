package useradap

import (
	"github.com/byorty/enterprise-application/pkg/user/infra/middleware"
	userrpimpl "github.com/byorty/enterprise-application/pkg/user/infra/repo"
	usersrvimpl "github.com/byorty/enterprise-application/pkg/user/infra/service"
	"go.uber.org/fx"
)

var Constructors = fx.Provide(
	usersrvimpl.NewFxUserService,
	usersrvimpl.NewFxUserProductService,
	middleware.NewFxUserRightsEnforcer,
	middleware.NewFxUserProductRightsEnforcer,
	userrpimpl.NewFxUserRepoFactory,
)
