package validator_test

import (
	"context"
	"github.com/Pallinder/go-randomdata"
	"github.com/byorty/enterprise-application/pkg/common/adapter/validator"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/reflect/protoreflect"
	"testing"
)

func TestStringValidatorSuite(t *testing.T) {
	suite.Run(t, new(StringValidatorSuite))
}

type StringValidatorSuite struct {
	suite.Suite
}

func (s *StringValidatorSuite) TestUUID() {
	v := validator.NewUUID()
	s.NotNil(v.Validate(context.Background(), protoreflect.ValueOf("")))
	s.NotNil(v.Validate(context.Background(), protoreflect.ValueOf(randomdata.Alphanumeric(32))))
	s.Nil(v.Validate(context.Background(), protoreflect.ValueOf(uuid.NewString())))
}

func (s *StringValidatorSuite) TestPhoneNumber() {
	v := validator.NewPhoneNumber()
	s.NotNil(v.Validate(context.Background(), protoreflect.ValueOf("")))
	s.NotNil(v.Validate(context.Background(), protoreflect.ValueOf(randomdata.Alphanumeric(32))))
	s.Nil(v.Validate(context.Background(), protoreflect.ValueOf("+79008007060")))
}

func (s *StringValidatorSuite) TestMaxLen() {
	v := validator.NewStringWithMaxLen(10)
	s.NotNil(v.Validate(context.Background(), protoreflect.ValueOf("")))
	s.NotNil(v.Validate(context.Background(), protoreflect.ValueOf(randomdata.Alphanumeric(32))))
	s.Nil(v.Validate(context.Background(), protoreflect.ValueOf(randomdata.Alphanumeric(9))))
}

func (s *StringValidatorSuite) TestAlphanumeric() {
	v := validator.NewAlphanumeric()
	s.NotNil(v.Validate(context.Background(), protoreflect.ValueOf("")))
	s.NotNil(v.Validate(context.Background(), protoreflect.ValueOf("qazwsx!@#%%^")))
	s.Nil(v.Validate(context.Background(), protoreflect.ValueOf(randomdata.Alphanumeric(32))))
}
