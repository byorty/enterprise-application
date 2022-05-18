package usersrv

import (
	"context"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
)

type UserProductService interface {
	GetProductsByUserUUID(ctx context.Context, userUUID string) ([]*pbv1.UserProduct, error)
	Put(ctx context.Context, userUUID string, params *pbv1.PutProductRequestParams) ([]*pbv1.UserProduct, error)
	Change(ctx context.Context, userUUID string, userProductUUID string, params *pbv1.ChangeProductRequestParams) ([]*pbv1.UserProduct, error)
}
