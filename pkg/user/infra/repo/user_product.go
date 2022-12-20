package userrpimpl

import (
	"context"

	"github.com/byorty/enterprise-application/pkg/common/adapter/db"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	userent "github.com/byorty/enterprise-application/pkg/user/domain/entity"
	userrp "github.com/byorty/enterprise-application/pkg/user/domain/repo"
	"github.com/go-pg/pg/v10"
)

func NewFxUserProductRepoFactory() userrp.UserProductRepoFactory {
	return &userProductRepoFactory{}
}

type userProductRepoFactory struct {
}

func (u userProductRepoFactory) Create(ctx context.Context, conn db.DB) userrp.UserProductRepo {
	return &userProductRepo{
		Repo: db.NewRepo[userent.UserProduct](ctx, conn),
	}
}

type userProductRepo struct {
	db.Repo[userent.UserProduct]
}

func (u userProductRepo) GetAllByFilter(ctx context.Context, params pbv1.GetUserProductRequestParams) (userent.UserProductSlice, error) {
	sl := userent.NewUserProductSlice()
	query := u.DB().Model(sl)
	if len(params.UuidIn) > 0 {
		query.Where("uuid IN (?)", pg.In(params.UuidIn))
	}

	if len(params.UserUuidIn) > 0 {
		query.Where("user_uuid IN (?)", pg.In(params.UserUuidIn))
	}

	if len(params.ProductUuidIn) > 0 {
		query.Where("product_uuid IN (?)", pg.In(params.ProductUuidIn))
	}

	err := query.Select()
	if err != nil {
		return nil, err
	}

	return sl, nil
}
