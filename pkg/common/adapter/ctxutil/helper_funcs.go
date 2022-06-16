package ctxutil

import "context"

const (
	KeySession = "session"
	KeyUser    = "user"
)

func Set(ctx context.Context, key string, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}

func Get[T any](ctx context.Context, key string) (T, bool) {
	v, ok := ctx.Value(key).(T)
	return v, ok
}
