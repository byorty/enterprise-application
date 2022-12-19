package userent

import (
	"time"

	"github.com/byorty/enterprise-application/pkg/common/collection"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type User struct {
	UUID        string `pg:"type:uuid,pk,default:gen_random_uuid()"`
	Group       pbv1.UserGroup
	Status      pbv1.UserStatus
	PhoneNumber string `pg:",unique"`
	Lastname    string
	Firstname   string
	CreatedAt   time.Time `pg:"default:now()"`
}

func (u User) GetUUID() string {
	return u.UUID
}

func (u User) ToProto() *pbv1.User {
	return &pbv1.User{
		Uuid:        u.UUID,
		Group:       u.Group,
		Status:      u.Status,
		PhoneNumber: u.PhoneNumber,
		Lastname:    u.Lastname,
		Firstname:   u.Firstname,
		CreatedAt:   timestamppb.New(u.CreatedAt),
	}
}

type UserSlice collection.PrototypedSlice[*User]

func NewUserSlice() UserSlice {
	return collection.NewPrototypedSlice[*User]()
}
