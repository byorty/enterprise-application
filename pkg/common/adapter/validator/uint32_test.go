package validator_test

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/validator"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/reflect/protoreflect"
	"testing"
)

func TestUint32ValidatorSuite(t *testing.T) {
	suite.Run(t, new(Uint32ValidatorSuite))
}

type Uint32ValidatorSuite struct {
	suite.Suite
}

func (s *Uint32ValidatorSuite) TestMin() {
	v := validator.NewUint32Min(10)
	s.NotNil(v.Validate(context.Background(), protoreflect.ValueOf(uint32(0))))
	s.Nil(v.Validate(context.Background(), protoreflect.ValueOf(uint32(10))))
	s.Nil(v.Validate(context.Background(), protoreflect.ValueOf(uint32(15))))
}
