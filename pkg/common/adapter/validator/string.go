package validator

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewString(format string) Validator {
	return &stringValidator{
		validator: validator.New(),
		format:    format,
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

type stringValidator struct {
	validator *validator.Validate
	format    string
}

func (v stringValidator) Validate(ctx context.Context, value protoreflect.Value) error {
	return v.validator.VarCtx(ctx, value.String(), v.format)
}
