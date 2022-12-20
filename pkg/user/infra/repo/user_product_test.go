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

func TestUserProductRepoSuite(t *testing.T) {
	suite.Run(t, new(UserProductRepoSuite))
}

type UserProductRepoSuite struct {
	suite_helper.RepoSuite
	product *userent.UserProduct
}

func (s *UserProductRepoSuite) SetupSuite() {
	s.product = &userent.UserProduct{
		UUID:        uuid.NewString(),
		ProductUUID: uuid.NewString(),
		UserUUID:    uuid.NewString(),
		Count:       uint32(randomdata.Number(1, 100)),
	}

	s.RepoSuite.SetupSuite()
	s.CreateTable(&userent.UserProduct{}, "test")
	s.Nil(s.DB.Insert(s.product))
}

func (s *UserProductRepoSuite) TestGetAllByFilter() {
	factory := userrpimpl.NewFxUserProductRepoFactory()
	repo := factory.Create(s.Ctx, s.DB)
	sl, err := repo.GetAllByFilter(s.Ctx, pbv1.GetUserProductRequestParams{
		UuidIn: []string{s.product.UUID},
	})
	s.Nil(err)
	s.NotNil(sl)
	s.Equal(sl.Len(), 1)
	s.Equal(sl.Get(0).UUID, s.product.UUID)

	sl, err = repo.GetAllByFilter(s.Ctx, pbv1.GetUserProductRequestParams{
		UuidIn: []string{uuid.NewString()},
	})
	s.Empty(sl)
	s.Nil(err)

	sl, err = repo.GetAllByFilter(s.Ctx, pbv1.GetUserProductRequestParams{
		ProductUuidIn: []string{s.product.ProductUUID},
	})
	s.Nil(err)
	s.NotNil(sl)
	s.Equal(sl.Len(), 1)
	s.Equal(sl.Get(0).UUID, s.product.UUID)

	sl, err = repo.GetAllByFilter(s.Ctx, pbv1.GetUserProductRequestParams{
		ProductUuidIn: []string{uuid.NewString()},
	})
	s.Empty(sl)
	s.Nil(err)

	sl, err = repo.GetAllByFilter(s.Ctx, pbv1.GetUserProductRequestParams{
		UserUuidIn: []string{s.product.UserUUID},
	})
	s.Nil(err)
	s.NotNil(sl)
	s.Equal(sl.Len(), 1)
	s.Equal(sl.Get(0).UUID, s.product.UUID)

	sl, err = repo.GetAllByFilter(s.Ctx, pbv1.GetUserProductRequestParams{
		UserUuidIn: []string{uuid.NewString()},
	})
	s.Empty(sl)
	s.Nil(err)
}
