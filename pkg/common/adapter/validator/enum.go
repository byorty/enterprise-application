package validator

import (
	"context"
	"github.com/go-playground/validator/v10"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewEnum(values map[int32]string) Validator {
	return &enumValidator{
		validator: validator.New(),
		values:    values,
	}
}

type enumValidator struct {
	validator *validator.Validate
	values    map[int32]string
}

func (e enumValidator) Validate(ctx context.Context, value protoreflect.Value) error {
	enumValue := int32(value.Enum())
	err := e.validator.VarCtx(ctx, enumValue, "gt=0")
	if err != nil {
		return err
	}

	_, ok := e.values[enumValue]
	return e.validator.VarCtx(ctx, ok, "eq=true")
}
