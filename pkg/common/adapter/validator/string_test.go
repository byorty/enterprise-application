package validator_test

import (
	"context"
	"github.com/Pallinder/go-randomdata"
	"github.com/byorty/enterprise-application/pkg/common/adapter/validator"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestStringValidatorSuite(t *testing.T) {
	suite.Run(t, new(StringValidatorSuite))
}

type StringValidatorSuite struct {
	suite.Suite
}

func (s *StringValidatorSuite) TestUUID() {
	msg := &pbv1.User{}
	protoMsg := msg.ProtoReflect()
	uuidField := protoMsg.Descriptor().Fields().Get(0)
	v := validator.NewUUID()

	s.NotNil(v.Validate(context.Background(), protoMsg.Get(uuidField)))

	msg.Uuid = randomdata.Alphanumeric(32)
	s.NotNil(v.Validate(context.Background(), protoMsg.Get(uuidField)))

	msg.Uuid = uuid.NewString()
	s.Nil(v.Validate(context.Background(), protoMsg.Get(uuidField)))
}

func (s *StringValidatorSuite) TestPhoneNumber() {
	msg := &pbv1.User{}
	protoMsg := msg.ProtoReflect()
	phoneNumber := protoMsg.Descriptor().Fields().Get(3)
	v := validator.NewPhoneNumber()

	s.NotNil(v.Validate(context.Background(), protoMsg.Get(phoneNumber)))

	msg.PhoneNumber = randomdata.Alphanumeric(32)
	s.NotNil(v.Validate(context.Background(), protoMsg.Get(phoneNumber)))

	msg.PhoneNumber = "+79998887766"
	s.Nil(v.Validate(context.Background(), protoMsg.Get(phoneNumber)))
}

func (s *StringValidatorSuite) TestMaxLen() {
	msg := &pbv1.Order{}
	protoMsg := msg.ProtoReflect()
	address := protoMsg.Descriptor().Fields().Get(3)
	v := validator.NewStringWithMaxLen(10)

	s.NotNil(v.Validate(context.Background(), protoMsg.Get(address)))

	msg.Address = randomdata.Alphanumeric(32)
	s.NotNil(v.Validate(context.Background(), protoMsg.Get(address)))

	msg.Address = randomdata.Alphanumeric(9)
	s.Nil(v.Validate(context.Background(), protoMsg.Get(address)))
}
