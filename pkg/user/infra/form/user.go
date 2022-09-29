package form

import "github.com/byorty/enterprise-application/pkg/common/adapter/validator"

var (
	PhoneNumber = validator.Form{
		"phone_number": validator.NewPhoneNumber(),
	}
	UserUUID = validator.Form{
		"user_uuid": validator.NewUUID(),
	}
)
