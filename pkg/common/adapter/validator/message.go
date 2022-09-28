package validator

import (
	"context"
	"github.com/go-playground/validator/v10"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewMessage(form Form) Validator {
	return &messageValidator{
		validator: validator.New(),
		form:      form,
	}
}

type messageValidator struct {
	validator *validator.Validate
	form      Form
}

func (v messageValidator) Validate(ctx context.Context, value protoreflect.Value) error {
	return v.form.Validate(ctx, value.Message())
}
