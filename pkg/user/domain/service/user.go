package usersrv

import (
	"context"

	userent "github.com/byorty/enterprise-application/pkg/user/domain/entity"
)

type UserService interface {
	Register(ctx context.Context, phoneNumber string) (string, error)
	Authorize(ctx context.Context, phoneNumber string) (string, error)
	GetByUUID(ctx context.Context, uuid string) (*userent.User, error)
}
