package usersrvimpl

import (
	"context"

	"github.com/byorty/enterprise-application/pkg/common/adapter/auth"
	"github.com/byorty/enterprise-application/pkg/common/adapter/db"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	userent "github.com/byorty/enterprise-application/pkg/user/domain/entity"
	userrp "github.com/byorty/enterprise-application/pkg/user/domain/repo"
	usersrv "github.com/byorty/enterprise-application/pkg/user/domain/service"
	"github.com/google/uuid"
)

func NewFxUserService(
	sessionManager auth.SessionManager,
	pool db.Pool,
	userRepoFactory userrp.UserRepoFactory,
) (usersrv.UserService, error) {
	conn, err := pool.Get(db.UserKind)
	if err != nil {
		return nil, err
	}

	return &userService{
		sessionManager:  sessionManager,
		userRepoFactory: userRepoFactory,
		conn:            conn,
	}, nil
}

type userService struct {
	sessionManager  auth.SessionManager
	userRepoFactory userrp.UserRepoFactory
	conn            db.DB
}

func (s *userService) Register(ctx context.Context, phoneNumber string) (string, error) {
	userRepo := s.userRepoFactory.Create(ctx, s.conn)
	user, _ := userRepo.GetActiveByPhoneNumber(ctx, phoneNumber)
	if user != nil {
		return "", usersrv.ErrUserAlreadyExists
	}

	user = &userent.User{
		UUID:        uuid.NewString(),
		Group:       pbv1.UserGroupCustomer,
		Status:      pbv1.UserStatusActive,
		PhoneNumber: phoneNumber,
	}
	err := userRepo.Save(ctx, user)
	if err != nil {
		return "", err
	}

	token, err := s.sessionManager.CreateToken(ctx, user.UUID, user.Group)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) Authorize(ctx context.Context, phoneNumber string) (string, error) {
	userRepo := s.userRepoFactory.Create(ctx, s.conn)
	user, err := userRepo.GetActiveByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		return "", usersrv.ErrUserNotFound
	}

	token, err := s.sessionManager.CreateToken(ctx, user.UUID, user.Group)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) GetByUUID(ctx context.Context, uuid string) (*userent.User, error) {
	userRepo := s.userRepoFactory.Create(ctx, s.conn)
	user, err := userRepo.GetByUUID(ctx, uuid)
	if err != nil {
		return nil, usersrv.ErrUserNotFound
	}

	return user, nil
}
