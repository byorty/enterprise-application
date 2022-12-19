package db_test

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/db"
	"github.com/byorty/enterprise-application/pkg/common/adapter/log"
	"github.com/byorty/enterprise-application/pkg/common/adapter/suite_helper"
	"sync"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
)

func TestPoolSuite(t *testing.T) {
	suite.Run(t, new(PoolSuite))
}

type PoolSuite struct {
	suite_helper.RepoSuite
	ctx context.Context
}

func (s *PoolSuite) TestFailureConnection() {
	p, err := db.NewPool(
		map[string]*db.Config{
			"default": {
				Host:     "127.0.0.1",
				Port:     0,
				Database: s.RepoSuite.DBName,
				Username: s.RepoSuite.DBUsername,
				Password: s.RepoSuite.DBPassword,
			},
		},
		log.NewFxLogger(),
		nil,
	)
	s.Nil(p)
	s.NotNil(err)
}

func (s *PoolSuite) TestFailureGetConnection() {
	p, err := db.NewPool(
		map[string]*db.Config{
			"kind1": {
				Host:     s.RepoSuite.DBHost,
				Port:     s.RepoSuite.DBPort,
				Database: s.RepoSuite.DBName,
				Username: s.RepoSuite.DBUsername,
				Password: s.RepoSuite.DBPassword,
			},
		},
		log.NewFxLogger(),
		nil,
	)
	s.NotNil(p)
	s.Nil(err)

	conn, err := p.Get("key-not-exist")
	s.Nil(conn)
	s.NotNil(err)
	s.True(errors.Is(err, db.ErrInvalidKind))
}

func (s *PoolSuite) TestSuccessGetConnection() {
	p, err := db.NewPool(
		map[string]*db.Config{
			"kind1": {
				Host:            s.RepoSuite.DBHost,
				Port:            s.RepoSuite.DBPort,
				Database:        s.RepoSuite.DBName,
				Username:        s.RepoSuite.DBUsername,
				Password:        s.RepoSuite.DBPassword,
				MaxOpenConns:    5,
				MaxRetries:      5,
				MinIdleConns:    2,
				WriteTimeout:    15 * time.Second,
				ReadTimeout:     15 * time.Second,
				MaxConnLifetime: 60 * time.Second,
			},
		},
		log.NewFxLogger(),
		nil,
	)
	s.NotNil(p)
	s.Nil(err)

	maxConnections := 10
	wg := &sync.WaitGroup{}
	wg.Add(maxConnections)
	for i := 0; i < maxConnections; i++ {
		go func(n int) {
			defer wg.Done()
			conn, err := p.Get("kind1")
			s.NotNil(conn)
			s.Nil(err)
			res, err := conn.Exec("SELECT pg_sleep(1)")
			s.NotNil(res)
			s.Nil(err)
		}(i)
	}
	wg.Wait()
}
