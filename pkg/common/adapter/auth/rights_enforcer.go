package auth

import (
	"context"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	"go.uber.org/fx"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type RightsEnforcer interface {
	Enforce(ctx context.Context, session pbv1.Session, value protoreflect.Value) (context.Context, error)
}

type RightsEnforcerDescriptorOut struct {
	fx.Out
	Descriptor RightsEnforcerDescriptor `group:"rights_enforcer"`
}

type RightsEnforcerDescriptor struct {
	Name           string
	RightsEnforcer RightsEnforcer
}
