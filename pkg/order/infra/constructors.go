package orderadap

import (
	ordersrvimpl "github.com/byorty/enterprise-application/pkg/order/infra/service"
	"go.uber.org/fx"
)

var Constructors = fx.Provide(
	ordersrvimpl.NewFxOrderService,
)
