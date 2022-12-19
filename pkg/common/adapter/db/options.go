package db

import (
	"context"
)

func WithForUpdate(tables ...string) Option {
	return func(ctx context.Context, db DB) DB {
		tx, ok := db.(*tx)
		if ok {
			tx.useForUpdate = true
			tx.lockTables = tables
		}
		return db
	}
}

func WithSkipLocked() Option {
	return func(ctx context.Context, db DB) DB {
		tx, ok := db.(*tx)
		if ok {
			tx.useForUpdateSkipLocked = true
		}
		return db
	}
}

func WithForNoKeyUpdate() Option {
	return func(ctx context.Context, db DB) DB {
		tx, ok := db.(*tx)
		if ok {
			tx.useForNoKeyUpdate = true
		}
		return db
	}
}
