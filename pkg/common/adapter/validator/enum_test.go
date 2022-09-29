package validator_test

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/validator"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/reflect/protoreflect"
	"testing"
)

func TestEnumValidatorSuite(t *testing.T) {
	suite.Run(t, new(EnumValidatorSuite))
}

type EnumValidatorSuite struct {
	suite.Suite
}

func (s *EnumValidatorSuite) TestEnum() {
	v := validator.NewEnum(pbv1.UserGroup_name)
	s.NotNil(v.Validate(context.Background(), protoreflect.ValueOf(pbv1.UserGroupGuest.Number())))
	s.NotNil(v.Validate(context.Background(), protoreflect.ValueOf(pbv1.UserGroup(9999).Number())))
	s.Nil(v.Validate(context.Background(), protoreflect.ValueOf(pbv1.UserGroupCustomer.Number())))
}
