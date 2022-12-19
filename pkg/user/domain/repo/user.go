package userrp

import (
	"context"

	"github.com/byorty/enterprise-application/pkg/common/adapter/db"
	userent "github.com/byorty/enterprise-application/pkg/user/domain/entity"
)

type UserRepoFactory interface {
	Create(ctx context.Context, conn db.DB) UserRepo
}

type UserRepo interface {
	db.Repo[userent.User]
	GetActiveByPhoneNumber(ctx context.Context, phoneNumber string) (*userent.User, error)
}
