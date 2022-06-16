package productsadap

import (
	productsrcimpl "github.com/byorty/enterprise-application/pkg/product/infra/service"
	"go.uber.org/fx"
)

var Constructors = fx.Provide(
	productsrcimpl.NewFxProductService,
)
