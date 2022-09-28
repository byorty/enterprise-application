package validator

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Validator interface {
	Validate(ctx context.Context, value protoreflect.Value) error
}

type Form map[string]Validator

func (f Form) Validate(ctx context.Context, message protoreflect.Message) error {
	fields := message.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		validator, ok := f[string(field.Name())]
		if !ok {
			continue
		}

		err := validator.Validate(ctx, message.Get(field))
		if err != nil {
			return grpc.ErrInvalidArgument(err)
		}
	}

	return nil
}
