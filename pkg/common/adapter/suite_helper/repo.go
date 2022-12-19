package suite_helper

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/dockertest"

	"github.com/stretchr/testify/suite"
)

type RepoSuite struct {
	suite.Suite
	ctx context.Context
	*dockertest.PG
}

func (s *RepoSuite) SetupSuite() {
	s.ctx = context.Background()
	var err error
	s.PG, err = dockertest.NewPG("")
	s.Nil(err)
}

func (s *RepoSuite) CreateTable(model interface{}, schema string) {
	err := s.PG.CreateTable(model, schema)
	s.Nil(err)
}

func (s *RepoSuite) TearDownSuite() {

}
