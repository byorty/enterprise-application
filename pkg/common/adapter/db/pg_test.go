package db_test

import (
	"context"
	"fmt"
	"github.com/byorty/enterprise-application/pkg/common/adapter/db"
	"github.com/byorty/enterprise-application/pkg/common/adapter/suite_helper"
	"github.com/go-pg/pg/v10"
	"sync"
	"testing"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/go-pg/pg/v10/orm"
	"github.com/stretchr/testify/suite"
)

func TestPgSuite(t *testing.T) {
	suite.Run(t, new(PgSuite))
}

type Model struct {
	tableName struct{} `pg:"test.model"`

	Id   uint64 `pg:",pk"`
	Name string
}

type PgSuite struct {
	suite_helper.RepoSuite

	ctx context.Context
}

func (s *PgSuite) SetupSuite() {
	s.RepoSuite.SetupSuite()

	_, err := s.DB.Exec("create schema if not exists test")
	s.Nil(err)

	err = s.DB.CreateTable(new(Model), &orm.CreateTableOptions{})
	s.Nil(err)
}

func (s *PgSuite) TestSimpleConnection() {
	dbConn := s.DB.WithCtx(context.WithValue(context.Background(), "test", "test"))
	result, err := dbConn.Exec("SELECT 1")
	s.NotNil(result)
	s.Nil(err)

	defer dbConn.Rollback()
	err = dbConn.Commit()
	s.Nil(err)

	cnt, err := dbConn.Model(new(Model)).SelectAndCount()
	s.Zero(cnt)
	s.NotNil(err)
}

func (s *PgSuite) TestTransaction() {
	ctx := context.WithValue(context.Background(), "test", "test")

	tx, err := s.DB.Begin(context.Background())
	s.NotNil(tx)
	s.Nil(err)

	tx = tx.WithCtx(ctx)
	result, err := tx.Exec("SELECT 1")
	s.NotNil(result)
	s.Nil(err)

	cnt, err := tx.Model(new(Model)).SelectAndCount()
	s.Zero(cnt)
	s.NotNil(err)

	err = tx.Commit()
	s.Nil(err)

	// protect duplicate begin
	tx, err = tx.Begin(context.WithValue(ctx, "test2", "test2"))
	s.NotNil(tx)
	s.Nil(err)

	result, err = tx.Exec("set statement_timeout TO 100;")
	s.Nil(result)
	s.NotNil(err)

	s.Nil(tx.Commit())
	s.Nil(tx.Commit())

	err = tx.Rollback()
	s.Nil(err)
}

func (s *PgSuite) TestXOptionWithForUpdate() {
	newModel := new(Model)
	_, err := s.DB.Model(newModel).Insert(newModel)
	s.Nil(err)

	wg := new(sync.WaitGroup)
	wg.Add(4)

	go func() {
		tx, err := s.DB.Begin(context.Background(), db.WithForUpdate())
		s.NotNil(tx)
		s.Nil(err)
		existsModel := new(Model)
		err = tx.Model(existsModel).Where("id = ?", newModel.Id).Select()
		s.Nil(err)
		s.Equal(newModel.Id, existsModel.Id)

		txSkipLocked, err := s.DB.Begin(context.Background(), db.WithForNoKeyUpdate(), db.WithSkipLocked())
		s.NotNil(txSkipLocked)
		err = txSkipLocked.Model(existsModel).
			Where("id = ?", newModel.Id).Select()
		s.Equal(pg.ErrNoRows, err)
		s.Nil(txSkipLocked.Commit())

		time.Sleep(time.Second)
		s.Nil(tx.Commit())
		wg.Done()
		s.T().Log(1)
	}()

	go func() {
		time.Sleep(time.Second / 2)
		model := new(Model)
		err = s.DB.Model(model).Where("id = ?", newModel.Id).Select()
		s.Nil(err)
		s.Equal(newModel.Id, model.Id)

		err = s.DB.WithForUpdate().Model(model).Where("id = ?", newModel.Id).Select()
		s.Nil(err)
		s.Equal(newModel.Id, model.Id)

		err = s.DB.WithForNoKeyUpdate().Model(model).Where("id = ?", newModel.Id).Select()
		s.Nil(err)
		s.Equal(newModel.Id, model.Id)

		err = s.DB.WithSkipLocked().Model(model).Where("id = ?", newModel.Id).Select()
		s.Nil(err)
		s.Equal(newModel.Id, model.Id)

		err = s.DB.Update(newModel)
		s.Nil(err)

		wg.Done()
		s.T().Log(2)
	}()

	go func() {
		time.Sleep(time.Second / 2)
		tx, err := s.DB.Begin(context.Background(), db.WithForUpdate())
		s.NotNil(tx)
		s.Nil(err)

		existsModel := new(Model)
		err = tx.Model(existsModel).Where("id = ?", newModel.Id).Select()
		s.Nil(err)
		s.Equal(newModel.Id, existsModel.Id)

		existsModel.Name = randomdata.Letters(32)
		_, err = tx.Model(existsModel).Update(existsModel)
		s.NotNil(err)
		s.Nil(tx.Commit())
		wg.Done()
		s.T().Log(3)
	}()

	go func() {
		time.Sleep(time.Second / 2)
		tx, err := s.DB.Begin(context.Background())
		s.NotNil(tx)
		s.Nil(err)

		existsModel := new(Model)
		err = tx.Model(existsModel).Where("id = ?", newModel.Id).Select()
		s.Nil(err)
		s.Equal(newModel.Id, existsModel.Id)

		err = tx.WithForUpdate().Model(existsModel).Where("id = ?", newModel.Id).Select()
		s.Nil(err)
		s.Equal(newModel.Id, existsModel.Id)

		err = tx.WithForUpdate("model").Model(existsModel).Where("id = ?", newModel.Id).Select()
		s.Nil(err)
		s.Equal(newModel.Id, existsModel.Id)

		err = tx.WithForUpdate("model").WithSkipLocked().Model(existsModel).Where("id = ?", newModel.Id).Select()
		s.Nil(err)
		s.Equal(newModel.Id, existsModel.Id)

		err = tx.WithForNoKeyUpdate("model").Model(existsModel).Where("id = ?", newModel.Id).Select()
		s.Nil(err)
		s.Equal(newModel.Id, existsModel.Id)

		err = tx.Model(existsModel).Where("id = ?", newModel.Id).Select()
		s.Nil(err)
		s.Equal(newModel.Id, existsModel.Id)

		err = tx.Update(existsModel)
		s.Nil(err)

		existsModel.Name = randomdata.Letters(32)
		_, err = tx.Model(existsModel).Update(existsModel)
		s.NotNil(err)
		s.Equal(err, fmt.Errorf("pg: Update and Delete queries require Where clause (try WherePK)"))

		s.NotNil(tx.Commit())
		wg.Done()
		s.T().Log(4)
	}()

	wg.Wait()
}

func (s *PgSuite) TestStat() {
	stats := s.DB.Stats()
	s.NotNil(stats)
	s.True(stats.TotalConns > 0)
}
