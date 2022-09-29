package form

import "github.com/byorty/enterprise-application/pkg/common/adapter/validator"

var (
	minProductCountValidator = validator.NewUint32Min(1)
	PutProduct               = validator.Form{
		"user_uuid": validator.NewUUID(),
		"params": validator.NewMessage(validator.Form{
			"product_uuid":  validator.NewUUID(),
			"product_count": minProductCountValidator,
		}),
	}
	ChangeProduct = validator.Form{
		"user_uuid":         validator.NewUUID(),
		"user_product_uuid": validator.NewUUID(),
		"params": validator.NewMessage(validator.Form{
			"count": minProductCountValidator,
		}),
	}
)
