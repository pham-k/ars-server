package idb

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"server/internal/config"
	"server/internal/log"
	"time"
)

type iDB struct {
	Pool    *redis.Client
	Log     log.Log
	Config  config.Config
	ConnStr string
}

func NewIDB(cfg config.Config, log log.Log) IDB {
	pool := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		//Password: cfg.IDB.Password,
		DB: cfg.IDB.Database,
	})

	connStr := fmt.Sprintf("redis://%s:%s@%s:%d/%d",
		cfg.IDB.Username, cfg.IDB.Password, cfg.IDB.Host, cfg.IDB.Port, cfg.IDB.Database)

	return &iDB{
		Pool:    pool,
		Log:     log,
		Config:  cfg,
		ConnStr: connStr,
	}
}

func (db iDB) Open() error {
	ctx := context.Background()

	if err := db.Pool.Ping(ctx).Err(); err != nil {
		db.Log.Error("Ping Redis", "error", err)
		return err
	}
	db.Log.Info("Ping Redis")

	return nil
}

func (db iDB) Close() error {
	if db.Pool == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.Pool.Shutdown(ctx).Err(); err != nil {
		db.Log.Error("Close Redis", "error", err)
		return err
	}
	db.Log.Info("Close Redis")

	return nil
}

func (db iDB) NewRepo() *redis.Client {
	return db.Pool
}
