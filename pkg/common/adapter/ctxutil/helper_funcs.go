package ctxutil

import (
	"context"
	"github.com/pkg/errors"
)

const (
	Session     = "session"
	User        = "user"
	UserProduct = "user_product"
	Order       = "order"
)

var (
	ErrNotFound = errors.Errorf("Контекст не содержит ключ")
)

func Set(ctx context.Context, key string, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}

func Get[T any](ctx context.Context, key string) (T, error) {
	v, ok := ctx.Value(key).(T)
	if !ok {
		return v, ErrNotFound
	}

	return v, nil
}
