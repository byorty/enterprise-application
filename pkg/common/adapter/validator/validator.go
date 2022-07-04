package validator

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/protoutil"
	"go.uber.org/fx"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Validator interface {
	Validate(ctx context.Context, message protoreflect.Message, field protoreflect.FieldDescriptor) error
}

type DescriptorOutFunc func() DescriptorOut

type DescriptorOut struct {
	fx.Out
	Descriptor Descriptor `group:"validator"`
}

type Descriptor struct {
	Names     []string
	Validator Validator
}

func NewFxMap(
	descriptors []Descriptor,
) *protoutil.Map[Validator] {
	validatorMap := protoutil.NewMap[Validator]()
	for _, descriptor := range descriptors {
		for _, name := range descriptor.Names {
			validatorMap.Set(name, descriptor.Validator)
		}
	}

	return validatorMap
}
