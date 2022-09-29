package validator

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewFloat64(condition string) Validator {
	return &float64Validator{
		condition: condition,
	}
}

func NewFloat64Gt(min float64) Validator {
	return NewFloat64(fmt.Sprintf("gt=%f", min))
}

func NewFloat64Lt(max float64) Validator {
	return NewFloat64(fmt.Sprintf("lt=%f", max))
}

type float64Validator struct {
	condition string
}

func (f float64Validator) Validate(ctx context.Context, value protoreflect.Value) error {
	return validate.VarCtx(ctx, value.Float(), f.condition)
}
