package form

import "github.com/byorty/enterprise-application/pkg/common/adapter/validator"

var (
	CreateOrder = validator.Form{
		"user_uuid": validator.NewUUID(),
		"params": validator.NewMessage(validator.Form{
			"products": validator.NewReqList(
				validator.NewMessage(validator.Form{
					"product_uuid": validator.NewUUID(),
					"count":        minProductCountValidator,
				}),
			),
			"address":      validator.NewStringWithMaxLen(255),
			"delivered_at": validator.NewDeliveredAt(),
		}),
	}
)
