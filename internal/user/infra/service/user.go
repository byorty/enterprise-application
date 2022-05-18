package usersrvimpl

import (
	"context"
	"github.com/Pallinder/go-randomdata"
	"github.com/byorty/enterprise-application/internal/pkg/collection"
	pbv1 "github.com/byorty/enterprise-application/internal/pkg/gen/api/proto/v1"
	usersrv "github.com/byorty/enterprise-application/internal/user/domain/service"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewUserService() usersrv.UserService {
	return &userService{
		users: collection.ImportMap[string, *pbv1.User](map[string]*pbv1.User{
			"387301f4-551c-4022-900a-80f6f76f3a10": {
				Uuid:        "387301f4-551c-4022-900a-80f6f76f3a10",
				Group:       pbv1.UserGroupCustomer,
				Status:      pbv1.UserStatusActive,
				PhoneNumber: "+79008007060",
				Lastname:    "Иванов",
				Firstname:   "Иван",
				CreatedAt:   timestamppb.Now(),
			},
		}),
	}
}

type userService struct {
	users collection.Map[string, *pbv1.User]
}

func (s *userService) Register(ctx context.Context, phoneNumber string) (*pbv1.User, error) {
	for _, user := range s.users.Entries() {
		if user.PhoneNumber == phoneNumber {
			return nil, usersrv.ErrUserAlreadyExists
		}
	}

	user := &pbv1.User{
		Uuid:        uuid.NewString(),
		Group:       pbv1.UserGroupCustomer,
		Status:      pbv1.UserStatusActive,
		PhoneNumber: phoneNumber,
		Lastname:    randomdata.LastName(),
		Firstname:   randomdata.FirstName(1),
		CreatedAt:   timestamppb.Now(),
	}

	s.users.Set(user.Uuid, user)
	return user, nil
}

func (s *userService) GetByUUID(ctx context.Context, uuid string) (*pbv1.User, error) {
	user, ok := s.users.Get(uuid)
	if !ok {
		return nil, usersrv.ErrUserNotFound
	}

	return user, nil
}
