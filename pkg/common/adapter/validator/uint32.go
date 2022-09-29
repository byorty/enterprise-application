package validator

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewUint32(condition string) Validator {
	return &uint32Validator{
		condition: condition,
	}
}

func NewUint32Min(min uint32) Validator {
	return NewUint32(fmt.Sprintf("gte=%d", min))
}

type uint32Validator struct {
	condition string
}

func (u uint32Validator) Validate(ctx context.Context, value protoreflect.Value) error {
	return validate.VarCtx(ctx, value.Uint(), u.condition)
}
