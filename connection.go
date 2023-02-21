package mgdb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	cfg      *Config
	clientDB *mongo.Database
}

func NewCtx(timeout time.Duration) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	return ctx
}

func NewClient(cfg *Config, clientDB *mongo.Database) *DB {
	return &DB{
		cfg:      cfg,
		clientDB: clientDB,
	}
}

func (d *DB) Close() error {
	return nil
}

func SetDefaultConfig() {

}
