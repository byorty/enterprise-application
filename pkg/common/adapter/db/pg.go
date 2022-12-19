package db

import (
	"context"
	"github.com/go-pg/pg/v10/orm"

	pg "github.com/go-pg/pg/v10"
)

type db struct {
	*pg.DB
}

func (d *db) Insert(model ...interface{}) error {
	_, err := d.DB.Model(model...).Insert()
	return err
}

func (d *db) Update(model interface{}) error {
	_, err := d.DB.Model(model).WherePK().Update()
	return err
}

func (d *db) CreateTable(model interface{}, options *orm.CreateTableOptions) error {
	return d.DB.Model(model).CreateTable(options)
}

func (d *db) DropTable(model interface{}, options *orm.DropTableOptions) error {
	return d.DB.Model(model).DropTable(options)
}

func (d *db) Stats() *pg.PoolStats {
	return d.DB.PoolStats()
}

func (d *db) WithForUpdate(_ ...string) DB {
	return d
}

func (d *db) WithForNoKeyUpdate(_ ...string) DB {
	return d
}

func (d *db) WithSkipLocked() DB {
	return d
}

func NewPG(conn *pg.DB) DB {
	return &db{
		DB: conn,
	}
}

func NewPGWithCtx(ctx context.Context, d *db) DB {
	return &db{
		DB: d.DB.WithContext(ctx),
	}
}

func (d *db) Begin(ctx context.Context, options ...Option) (DB, error) {
	pgTx, err := d.DB.WithContext(ctx).Begin()
	if err != nil {
		return nil, err
	}

	var newTx DB = &tx{
		db: d,
		Tx: pgTx,
	}

	for _, option := range options {
		newTx = option(ctx, newTx)
	}

	return newTx, nil
}

func (d *db) Commit() error {
	return nil
}

func (d *db) Rollback() error {
	return nil
}

func (d *db) WithCtx(ctx context.Context) DB {
	return NewPGWithCtx(ctx, d)
}
