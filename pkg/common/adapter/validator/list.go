package validator

import (
	"context"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewList(v Validator) Validator {
	return &listValidator{
		validator: v,
	}
}

func NewReqList(v Validator) Validator {
	return &listValidator{
		validator: v,
		required:  true,
	}
}

type listValidator struct {
	validator Validator
	required  bool
}

func (v listValidator) Validate(ctx context.Context, value protoreflect.Value) error {
	list := value.List()
	if v.required {
		err := validate.VarCtx(ctx, list.Len(), "gt=0")
		if err != nil {
			return err
		}
	}

	for i := 0; i < list.Len(); i++ {
		err := v.validator.Validate(ctx, list.Get(i))
		if err != nil {
			return err
		}
	}

	return nil
}
