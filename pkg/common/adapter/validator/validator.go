package validator

import (
	"context"
	"github.com/go-playground/validator/v10"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	validate = validator.New()
)

type Validator interface {
	Validate(ctx context.Context, value protoreflect.Value) error
}

type Form map[string]Validator

func (f Form) Validate(ctx context.Context, message protoreflect.Message) error {
	fields := message.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		v, ok := f[string(field.Name())]
		if !ok {
			continue
		}

		err := v.Validate(ctx, message.Get(field))
		if err != nil {
			return err
		}
	}

	return nil
}
