package db

import (
	"context"

	pg "github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type Option func(ctx context.Context, db DB) DB

type DB interface {
	Begin(ctx context.Context, options ...Option) (DB, error)
	Commit() error
	Rollback() error
	Close() error
	Model(...interface{}) *orm.Query
	Insert(...interface{}) error
	Update(interface{}) error
	CreateTable(interface{}, *orm.CreateTableOptions) error
	DropTable(interface{}, *orm.DropTableOptions) error
	Exec(interface{}, ...interface{}) (pg.Result, error)
	WithCtx(context.Context) DB
	WithForUpdate(...string) DB
	WithForNoKeyUpdate(...string) DB
	WithSkipLocked() DB
	Stats() *pg.PoolStats
}
