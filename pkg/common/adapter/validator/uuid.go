package validator

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation"
	is "github.com/go-ozzo/ozzo-validation/is"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewFxUUID(names ...string) DescriptorOutFunc {
	return func() DescriptorOut {
		return DescriptorOut{
			Descriptor: Descriptor{
				Names:     names,
				Validator: new(uuidValidator),
			},
		}
	}
}

type uuidValidator struct{}

func (v uuidValidator) Validate(ctx context.Context, message protoreflect.Message, field protoreflect.FieldDescriptor) error {
	return validation.Validate(message.Get(field).String(), validation.Required, is.UUID)
}
