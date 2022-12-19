package db

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/log"
	"io"
	"strings"
	"time"

	pg "github.com/go-pg/pg/v10"

	"github.com/pkg/errors"
)

type queryStartTime struct {
	start time.Time
}
type LogHook struct {
	Logger log.Logger
}

func (d LogHook) BeforeQuery(ctx context.Context, _ *pg.QueryEvent) (context.Context, error) {
	return context.WithValue(ctx, queryStartTimeKey, queryStartTime{start: time.Now()}), nil
}

func (d LogHook) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	queryTime := ctx.Value(queryStartTimeKey).(queryStartTime)
	logger := d.Logger.WithCtx(ctx, "duration", time.Now().Sub(queryTime.start).Milliseconds())
	if q.Err != nil {
		switch {
		case !errors.Is(q.Err, pg.ErrTxDone) && !errors.Is(q.Err, io.EOF):
			logger.Error(q.Err, q.Query)
		default:
			logger.Warn(q.Err)
		}
	}

	sql, err := q.FormattedQuery()
	if err != nil {
		logger.Error(err)
		return nil
	}
	logger.Info(strings.ReplaceAll(string(sql), `"`, ``))
	return nil
}
