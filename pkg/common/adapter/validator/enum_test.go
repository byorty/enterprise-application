package validator_test

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/validator"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestEnumValidatorSuite(t *testing.T) {
	suite.Run(t, new(EnumValidatorSuite))
}

type EnumValidatorSuite struct {
	suite.Suite
}

func (s *EnumValidatorSuite) TestEnum() {
	msg := &pbv1.User{}
	protoMsg := msg.ProtoReflect()
	groupField := protoMsg.Descriptor().Fields().Get(1)
	v := validator.NewEnum(pbv1.UserGroup_name)

	s.NotNil(v.Validate(context.Background(), protoMsg.Get(groupField)))

	msg.Group = pbv1.UserGroup(999)
	s.NotNil(v.Validate(context.Background(), protoMsg.Get(groupField)))

	msg.Group = pbv1.UserGroupCustomer
	s.Nil(v.Validate(context.Background(), protoMsg.Get(groupField)))
}
