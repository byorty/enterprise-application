package userent

import (
	"time"

	"github.com/byorty/enterprise-application/pkg/common/collection"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
)

type UserProduct struct {
	UUID        string `pg:"type:uuid,pk,default:gen_random_uuid()"`
	ProductUUID string `pg:"type:uuid"`
	UserUUID    string `pg:"type:uuid"`
	Count       uint32
	CreatedAt   time.Time `pg:"default:now()"`
}

func (u UserProduct) GetUUID() string {
	return u.UUID
}

func (u UserProduct) ToProto() *pbv1.UserProduct {
	return &pbv1.UserProduct{
		Uuid:        u.UUID,
		ProductUuid: u.ProductUUID,
		Count:       u.Count,
	}
}

type UserProductSlice collection.Slice[*UserProduct]

func NewUserProductSlice() UserProductSlice {
	return collection.NewSlice[*UserProduct]()
}
