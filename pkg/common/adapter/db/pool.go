package db

import (
	"context"
	"fmt"
	"github.com/byorty/enterprise-application/pkg/common/adapter/log"
	"net"
	"sync"

	pg "github.com/go-pg/pg/v10"

	"go.uber.org/fx"
)

const (
	DefaultKind Kind = "default"
	UserKind    Kind = "user"

	queryStartTimeKey = "queryDuration"
)

type Kind string

type Pool interface {
	Get(Kind) (DB, error)
}

func NewPool(dbs map[string]*Config, logger log.Logger, lc fx.Lifecycle) (Pool, error) {
	p := &pool{
		connections: make(map[Kind]DB),
		logger:      logger.Named("pool"),
		mtx:         new(sync.Mutex),
	}
	for key, cfg := range dbs {
		conn := pg.Connect(&pg.Options{
			Addr:                  net.JoinHostPort(cfg.Host, fmt.Sprintf("%d", cfg.Port)),
			User:                  cfg.Username,
			Password:              cfg.Password,
			Database:              cfg.Database,
			PoolSize:              cfg.MaxOpenConns,
			MaxConnAge:            cfg.MaxConnLifetime,
			ReadTimeout:           cfg.ReadTimeout,
			WriteTimeout:          cfg.WriteTimeout,
			MaxRetries:            cfg.MaxRetries,
			MinIdleConns:          cfg.MinIdleConns,
			RetryStatementTimeout: true,
			ApplicationName:       key,
		})

		_, err := conn.Exec("SELECT 1")
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		conn.AddQueryHook(LogHook{Logger: logger.Named("query_logger")})
		p.connections[Kind(key)] = NewPG(conn)
	}

	if lc != nil {
		lc.Append(fx.Hook{
			OnStop: func(_ context.Context) error {
				p.mtx.Lock()
				defer p.mtx.Unlock()
				p.logger.Info("db pools are closing")
				for kind, conn := range p.connections {
					p.logger.Infof("closing %s", kind)
					err := conn.Close()
					if err != nil {
						p.logger.Error(err)
					}
				}
				p.logger.Info("db pools has been closed")
				return nil
			},
		})
	}

	return p, nil
}

type pool struct {
	connections map[Kind]DB
	mtx         *sync.Mutex
	logger      log.Logger
}

func (p *pool) Get(kind Kind) (DB, error) {
	logger := p.logger.With("method", "get")
	p.mtx.Lock()
	defer p.mtx.Unlock()
	conn, ok := p.connections[kind]
	if !ok {
		logger.Errorf("db kind=%s not found", kind)
		return nil, ErrInvalidKind
	}
	return conn, nil
}
