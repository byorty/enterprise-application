package validator

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewFxOptionalUint32() DescriptorOut {
	return DescriptorOut{
		Descriptor: Descriptor{
			Names: []string{
				"offset",
				"limit",
			},
			Validator: &uint32Validator{
				rules: []validation.Rule{
					is.Int,
				},
			},
		},
	}
}

func NewFxRequiredUint32() DescriptorOut {
	return DescriptorOut{
		Descriptor: Descriptor{
			Names: []string{
				"count",
			},
			Validator: &uint32Validator{
				rules: []validation.Rule{
					validation.Required,
					is.Int,
					validation.Min(0),
				},
			},
		},
	}
}

type uint32Validator struct {
	rules []validation.Rule
}

func (v uint32Validator) Validate(ctx context.Context, message protoreflect.Message, field protoreflect.FieldDescriptor) error {
	return validation.Validate(uint32(message.Get(field).Uint()), v.rules...)
}
