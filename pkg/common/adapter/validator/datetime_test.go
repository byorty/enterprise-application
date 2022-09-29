package validator_test

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/validator"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestDatetimeValidatorSuite(t *testing.T) {
	suite.Run(t, new(DatetimeValidatorSuite))
}

type DatetimeValidatorSuite struct {
	suite.Suite
}

func (s *DatetimeValidatorSuite) TestAll() {
	user := &pbv1.User{}
	now := timestamppb.Now().ProtoReflect()
	v := validator.NewMinDatetime()
	s.NotNil(v.Validate(context.Background(), protoreflect.ValueOf(user.ProtoReflect())))
	s.NotNil(v.Validate(context.Background(), protoreflect.ValueOf(timestamppb.New(time.Time{}).ProtoReflect())))
	s.Nil(v.Validate(context.Background(), protoreflect.ValueOf(now)))

	v = validator.NewDeliveredAt()
	s.NotNil(v.Validate(context.Background(), protoreflect.ValueOf(now)))
}
