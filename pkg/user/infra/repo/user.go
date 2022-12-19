package userrpimpl

import (
	"context"

	"github.com/byorty/enterprise-application/pkg/common/adapter/db"
	userent "github.com/byorty/enterprise-application/pkg/user/domain/entity"
	userrp "github.com/byorty/enterprise-application/pkg/user/domain/repo"
)

func NewFxUserRepoFactory() userrp.UserRepoFactory {
	return &userRepoFactory{}
}

type userRepoFactory struct {
}

func (u userRepoFactory) Create(ctx context.Context, conn db.DB) userrp.UserRepo {
	return &userRepo{
		Repo: db.NewRepo[userent.User](ctx, conn),
	}
}

type userRepo struct {
	db.Repo[userent.User]
}

func (u userRepo) GetActiveByPhoneNumber(ctx context.Context, phoneNumber string) (*userent.User, error) {
	var model userent.User
	err := u.DB().Model(&model).Where("phone_number = ?", phoneNumber).Select()
	if err != nil {
		return nil, err
	}

	return &model, nil
}
