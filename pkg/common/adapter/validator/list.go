package validator

import (
	"context"
	"github.com/go-playground/validator/v10"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewList(v Validator) Validator {
	return &listValidator{
		validator:     validator.New(),
		itemValidator: v,
	}
}

func NewReqList(v Validator) Validator {
	return &listValidator{
		validator:     validator.New(),
		itemValidator: v,
		required:      true,
	}
}

type listValidator struct {
	validator     *validator.Validate
	itemValidator Validator
	required      bool
}

func (v listValidator) Validate(ctx context.Context, value protoreflect.Value) error {
	list := value.List()
	if v.required {
		err := v.validator.VarCtx(ctx, list.Len(), "gt=0")
		if err != nil {
			return err
		}
	}

	for i := 0; i < list.Len(); i++ {
		err := v.itemValidator.Validate(ctx, list.Get(i))
		if err != nil {
			return err
		}
	}

	return nil
}
