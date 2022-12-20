package userrp

import (
	"context"

	"github.com/byorty/enterprise-application/pkg/common/adapter/db"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	userent "github.com/byorty/enterprise-application/pkg/user/domain/entity"
)

type UserProductRepoFactory interface {
	Create(ctx context.Context, conn db.DB) UserProductRepo
}

type UserProductRepo interface {
	db.Repo[userent.UserProduct]
	GetAllByFilter(ctx context.Context, params pbv1.GetUserProductRequestParams) (userent.UserProductSlice, error)
}
