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

func TestValidatorSuite(t *testing.T) {
	suite.Run(t, new(ValidatorSuite))
}

type ValidatorSuite struct {
	suite.Suite
}

func (s *ValidatorSuite) TestMessage() {
	form := validator.Form{
		"uuid": validator.NewUUID(),
		"properties": validator.NewReqList(
			validator.NewMessage(validator.Form{
				"name":  validator.NewReqString(),
				"value": validator.NewReqString(),
			}),
		),
	}
	msg := &pbv1.Product{}
	protoMsg := msg.ProtoReflect()

	s.NotNil(form.Validate(context.Background(), protoMsg))

	msg.Uuid = uuid.NewString()
	s.NotNil(form.Validate(context.Background(), protoMsg))

	msg.Properties = []*pbv1.ProductProperty{
		{},
	}
	s.NotNil(form.Validate(context.Background(), protoMsg))

	msg.Properties = []*pbv1.ProductProperty{
		{
			Name:  randomdata.Alphanumeric(10),
			Value: randomdata.Alphanumeric(20),
		},
		{
			Name:  randomdata.Alphanumeric(10),
			Value: randomdata.Alphanumeric(20),
		},
	}
	s.Nil(form.Validate(context.Background(), protoMsg))
}
