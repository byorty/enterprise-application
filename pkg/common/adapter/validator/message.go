package validator

import (
	"context"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewMessage(form Form) Validator {
	return &messageValidator{
		form: form,
	}
}

type messageValidator struct {
	form Form
}

func (v messageValidator) Validate(ctx context.Context, value protoreflect.Value) error {
	return v.form.Validate(ctx, value.Message())
}
