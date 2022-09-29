package form

import (
	"github.com/byorty/enterprise-application/pkg/common/adapter/validator"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
)

var (
	Checkout = validator.Form{
		"order_uuid": validator.NewUUID(),
		"params": validator.NewMessage(validator.Form{
			"amount": validator.NewFloat64Gt(0),
			"status": validator.NewEnum(pbv1.OrderStatus_name),
		}),
	}
)
