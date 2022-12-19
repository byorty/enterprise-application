package db

import "time"

type Config struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	Database        string        `yaml:"database"`
	Schema          string        `yaml:"schema"`
	Username        string        `yaml:"username"`
	Password        string        `yaml:"password"`
	MaxOpenConns    int           `yaml:"max_open_conns"`
	MaxRetries      int           `yaml:"max_retries"`
	MinIdleConns    int           `yaml:"min_idle_conns"`
	MaxConnLifetime time.Duration `yaml:"max_conn_lifetime"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`
	WriteTimeout    time.Duration `yaml:"write_timeout"`
}
