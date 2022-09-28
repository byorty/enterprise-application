package validator_test

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/validator"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
)

func TestDatetimeValidatorSuite(t *testing.T) {
	suite.Run(t, new(DatetimeValidatorSuite))
}

type DatetimeValidatorSuite struct {
	suite.Suite
}

func (s *DatetimeValidatorSuite) TestAll() {
	msg := &pbv1.User{
		CreatedAt: timestamppb.Now(),
	}
	protoMsg := msg.ProtoReflect()
	createdAtField := protoMsg.Descriptor().Fields().Get(protoMsg.Descriptor().Fields().Len() - 1)
	v := validator.NewMinDatetime()
	s.Nil(v.Validate(context.Background(), protoMsg.Get(createdAtField)))

	v = validator.NewDeliveredAt()
	s.NotNil(v.Validate(context.Background(), protoMsg.Get(createdAtField)))
}
