package usersrv

import (
	"context"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
)

type UserService interface {
	Register(ctx context.Context, phoneNumber string) (*pbv1.User, error)
	GetByUUID(ctx context.Context, uuid string) (*pbv1.User, error)
}
