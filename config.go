package mgdb

import "time"

type Config struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DBName       string
	MaxOpenConns int
	CtxTimeout   time.Duration
}

func NewConfig(host, port, username, password, dbname string, open int, ctxTimeout time.Duration) *Config {
	return &Config{Host: host, Port: port, Username: username, Password: password, DBName: dbname, MaxOpenConns: open, CtxTimeout: ctxTimeout}
}
