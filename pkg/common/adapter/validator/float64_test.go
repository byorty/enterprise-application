package validator_test

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/validator"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/reflect/protoreflect"
	"testing"
)

func TestFloat64ValidatorSuite(t *testing.T) {
	suite.Run(t, new(Float64ValidatorSuite))
}

type Float64ValidatorSuite struct {
	suite.Suite
}

func (s *Float64ValidatorSuite) TestGt() {
	v := validator.NewFloat64Gt(10)
	s.NotNil(v.Validate(context.Background(), protoreflect.ValueOf(float64(0))))
	s.Nil(v.Validate(context.Background(), protoreflect.ValueOf(float64(11))))
}

func (s *Float64ValidatorSuite) TestLt() {
	v := validator.NewFloat64Lt(10)
	s.NotNil(v.Validate(context.Background(), protoreflect.ValueOf(float64(11))))
	s.Nil(v.Validate(context.Background(), protoreflect.ValueOf(float64(9))))
}
