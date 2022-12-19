package db

import (
	"context"
	"strings"

	pg "github.com/go-pg/pg/v10"

	"github.com/go-pg/pg/v10/orm"
)

type tx struct {
	*db
	*pg.Tx
	lockTables             []string
	useForKeyShare         bool
	useForShare            bool
	useForNoKeyUpdate      bool
	useForUpdate           bool
	useForUpdateSkipLocked bool
	useForUpdateNoWait     bool
	done                   bool
}

func (t *tx) Stats() *pg.PoolStats {
	return t.db.PoolStats()
}

func (t *tx) Model(obj ...interface{}) *orm.Query {
	query := t.Tx.Model(obj...)

	withFor := ""
	if t.useForKeyShare {
		withFor = "KEY SHARE"
	}

	if t.useForShare {
		withFor = "SHARE"
	}

	if t.useForNoKeyUpdate {
		withFor = "NO KEY UPDATE"
	}

	if t.useForUpdate {
		withFor = "UPDATE"
	}

	if len(t.lockTables) > 0 {
		withFor += " OF " + strings.Join(t.lockTables, ", ")
	}

	if t.useForUpdateSkipLocked {
		withFor += " SKIP LOCKED"
	}

	if t.useForUpdateNoWait {
		withFor += " NOWAIT"
	}

	if withFor != "" {
		return query.For(withFor)
	}

	return query
}

func (t *tx) WithCtx(context.Context) DB {
	return t
}

func (t *tx) WithForUpdate(tables ...string) DB {
	newTx := *t
	newTx.useForUpdate = true
	newTx.lockTables = tables
	return &newTx
}

func (t *tx) WithForNoKeyUpdate(tables ...string) DB {
	newTx := *t
	newTx.useForNoKeyUpdate = true
	newTx.lockTables = tables
	return &newTx
}

func (t *tx) WithSkipLocked() DB {
	newTx := *t
	newTx.useForUpdateSkipLocked = true
	return &newTx
}

func (t *tx) Begin(ctx context.Context, options ...Option) (DB, error) {
	return t, nil
}

func (t *tx) Commit() error {
	if t.done {
		return nil
	}

	if err := t.Tx.Commit(); err != nil {
		return err
	}
	t.done = true
	return nil
}

func (t *tx) Rollback() error {
	if t.done {
		return nil
	}
	return t.Tx.Rollback()
}
