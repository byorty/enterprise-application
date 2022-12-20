package userrpimpl_test

import (
	"testing"

	"github.com/Pallinder/go-randomdata"
	"github.com/byorty/enterprise-application/pkg/common/adapter/suite_helper"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	userent "github.com/byorty/enterprise-application/pkg/user/domain/entity"
	userrpimpl "github.com/byorty/enterprise-application/pkg/user/infra/repo"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

func TestUserRepoSuite(t *testing.T) {
	suite.Run(t, new(UserRepoSuite))
}

type UserRepoSuite struct {
	suite_helper.RepoSuite
	activeUser      *userent.User
	unconfirmedUser *userent.User
}

func (s *UserRepoSuite) SetupSuite() {
	s.activeUser = &userent.User{
		UUID:        uuid.NewString(),
		Group:       pbv1.UserGroupCustomer,
		Status:      pbv1.UserStatusActive,
		PhoneNumber: randomdata.PhoneNumber(),
	}
	s.unconfirmedUser = &userent.User{
		UUID:        uuid.NewString(),
		Group:       pbv1.UserGroupCustomer,
		Status:      pbv1.UserStatusUnconfirmed,
		PhoneNumber: randomdata.PhoneNumber(),
	}

	s.RepoSuite.SetupSuite()
	s.CreateTable(&userent.User{}, "test")
	s.Nil(s.DB.Insert(s.activeUser, s.unconfirmedUser))
}

func (s *UserRepoSuite) TestGetActiveByPhoneNumber() {
	factory := userrpimpl.NewFxUserRepoFactory()
	repo := factory.Create(s.Ctx, s.DB)

	user, err := repo.GetActiveByPhoneNumber(s.Ctx, s.unconfirmedUser.PhoneNumber)
	s.Nil(user)
	s.NotNil(err)

	user, err = repo.GetActiveByPhoneNumber(s.Ctx, s.activeUser.PhoneNumber)
	s.Nil(err)
	s.NotNil(user)
	s.Equal(s.activeUser.PhoneNumber, user.PhoneNumber)
}
