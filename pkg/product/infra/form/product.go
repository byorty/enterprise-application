package form

import "github.com/byorty/enterprise-application/pkg/common/adapter/validator"

var (
	ProductUUID = validator.Form{
		"product_uuid": validator.NewUUID(),
	}
	Products = validator.Form{
		"filter": validator.NewMessage(validator.Form{
			"uuid_in":       validator.NewOptList(validator.NewUUID()),
			"name_contains": validator.NewAlphanumeric(),
			"price_gt":      validator.NewFloat64Gt(0),
			"price_lt":      validator.NewFloat64Lt(1000000),
			"properties_eq": validator.NewOptList(
				validator.NewMessage(validator.Form{
					"name":  validator.NewAlphanumeric(),
					"value": validator.NewAlphanumeric(),
				}),
			),
		}),
	}
)
