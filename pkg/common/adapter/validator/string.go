package validator

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewString(condition string) Validator {
	return &stringValidator{
		condition: condition,
	}
}

func NewUUID() Validator {
	return NewString("uuid4")
}

func NewPhoneNumber() Validator {
	return NewString("e164")
}

func NewStringWithMaxLen(l int) Validator {
	return NewString(fmt.Sprintf("required,max=%d", l))
}

func NewReqString() Validator {
	return NewString("required")
}

func NewAlphanumeric() Validator {
	return NewString("alphanum")
}

type stringValidator struct {
	condition string
}

func (v stringValidator) Validate(ctx context.Context, value protoreflect.Value) error {
	return validate.VarCtx(ctx, value.String(), v.condition)
}
