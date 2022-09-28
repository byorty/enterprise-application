package validator

import (
	"context"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewEnum(values map[int32]string) Validator {
	return &enumValidator{
		values: values,
	}
}

type enumValidator struct {
	values map[int32]string
}

func (e enumValidator) Validate(ctx context.Context, value protoreflect.Value) error {
	enumValue := int32(value.Enum())
	err := validate.VarCtx(ctx, enumValue, "gt=0")
	if err != nil {
		return err
	}

	_, ok := e.values[enumValue]
	return validate.VarCtx(ctx, ok, "eq=true")
}
