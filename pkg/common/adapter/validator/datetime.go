package validator

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func NewDatetime(duration time.Duration) Validator {
	return &datetimeValidator{
		duration: duration,
	}
}

func NewMinDatetime() Validator {
	return NewDatetime(-time.Hour * 24 * 365 * 100)
}

func NewDeliveredAt() Validator {
	return NewDatetime(time.Hour * 48)
}

type datetimeValidator struct {
	duration time.Duration
}

func (v datetimeValidator) Validate(ctx context.Context, value protoreflect.Value) error {
	actualDatetime, ok := value.Message().Interface().(*timestamppb.Timestamp)
	err := validate.VarCtx(ctx, ok, "eq=true")
	if err != nil {
		return err
	}

	expectedDatetime := time.Now().Add(v.duration)
	return validate.VarCtx(ctx, actualDatetime.AsTime().Unix(), fmt.Sprintf("gte=%d", expectedDatetime.Unix()))
}
