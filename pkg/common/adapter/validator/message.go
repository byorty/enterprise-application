package validator

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/protoutil"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewFxMessage(
	validators *protoutil.Map[Validator],
) DescriptorOut {
	return DescriptorOut{
		Descriptor: Descriptor{
			Names: []string{
				"message",
			},
			Validator: &messageValidator{
				validators: validators,
			},
		},
	}
}

type messageValidator struct {
	validators *protoutil.Map[Validator]
}

func (v messageValidator) Validate(ctx context.Context, message protoreflect.Message, _ protoreflect.FieldDescriptor) error {
	fields := message.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)

		validator, err := v.validators.Get(message, field)
		if err != nil {
			continue
		}

		err = validator.Validate(ctx, message, field)
		if err != nil {
			return err
		}
	}

	return nil
}
