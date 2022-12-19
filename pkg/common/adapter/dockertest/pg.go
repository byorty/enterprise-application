package dockertest

import (
	"fmt"
	"github.com/byorty/enterprise-application/pkg/common/adapter/db"
	"github.com/byorty/enterprise-application/pkg/common/adapter/log"
	"github.com/go-pg/pg/v10"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/go-pg/pg/v10/orm"
	dockertest "github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

const (
	postgresUser          = "ecosystem"
	postgresPassword      = "ecosystem"
	postgresContainerName = "pg-for-test"
	postgresDB            = "ecosystem"
)

type PG struct {
	DBHost     string
	DBPort     int
	DBName     string
	DBUsername string
	DBPassword string
	DB         db.DB
}

func NewPG(dbName string) (*PG, error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, err
	}
	if dbName == "" {
		dbName = "test_db_" + strings.ToLower(randomdata.Letters(20))
	}
	logger := log.NewFxLogger()
	resource, ok := pool.ContainerByName(postgresContainerName)
	if !ok {
		resource, err = pool.RunWithOptions(&dockertest.RunOptions{
			Name:       postgresContainerName,
			Repository: "postgres",
			Tag:        "alpine",
			Env: []string{
				fmt.Sprintf("POSTGRES_PASSWORD=%s", postgresPassword),
				fmt.Sprintf("POSTGRES_USER=%s", postgresUser),
				fmt.Sprintf("POSTGRES_DB=%s", postgresDB),
				"listen_addresses = '*'",
			},
			PortBindings: map[docker.Port][]docker.PortBinding{
				"5432/tcp": {
					{
						HostIP:   "127.0.0.1",
						HostPort: "5431/tcp",
					},
				},
			},
			Cmd: []string{"postgres", "-c", "wal_level=logical"},
		}, func(config *docker.HostConfig) {
			// set AutoRemove to true so that stopped container goes away by itself
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{Name: "no"}
		})
		if err != nil {
			return nil, err
		}
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	if err != nil {
		return nil, err
	}

	host, portStr, err := net.SplitHostPort(hostAndPort)
	if err != nil {
		return nil, err
	}
	if host == "localhost" {
		host = "127.0.0.1"
	}
	logger.Debugf("start container %s on %s:%s", postgresContainerName, host, portStr)
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}

	pool.MaxWait = time.Second * 30
	pgForTest := &PG{
		DBHost:     host,
		DBPort:     port,
		DBName:     dbName,
		DBPassword: postgresPassword,
		DBUsername: postgresUser,
	}
	err = pool.Retry(func() error {
		err = pgForTest.InitDB()
		if err != nil {
			logger.Debug(err)
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	err = resource.Expire(60 * 15)
	if err != nil {
		return nil, err
	}
	return pgForTest, nil
}

func (p *PG) InitDB() error {
	logger := log.NewFxLogger()
	common, err := p.Connect(logger, "postgres")
	if err != nil {
		return err
	}
	_, _ = common.Exec("CREATE DATABASE " + p.DBName)

	newDb, err := p.Connect(logger, p.DBName)
	if err != nil {
		return err
	}
	p.DB = newDb
	return nil
}

func (p *PG) Connect(logger log.Logger, dbName string) (db.DB, error) {
	conn := pg.Connect(&pg.Options{
		Addr:     net.JoinHostPort(p.DBHost, fmt.Sprintf("%d", p.DBPort)),
		User:     p.DBUsername,
		Password: p.DBPassword,
		Database: dbName,
	})

	_, err := conn.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}
	conn.AddQueryHook(db.LogHook{Logger: logger.Named("query_logger")})
	return db.NewPG(conn), nil
}

func (p *PG) CreateTable(model interface{}, schema string) error {
	if schema != "" {
		_, err := p.DB.Exec("create schema if not exists " + schema)
		if err != nil {
			return err
		}
	}

	return p.DB.CreateTable(model, &orm.CreateTableOptions{
		IfNotExists: true,
	})
}
