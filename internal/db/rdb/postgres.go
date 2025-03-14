package rdb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"server/internal/config"
	"server/internal/log"
	"server/internal/repository"
	"time"
)

type rDB struct {
	Pool    *sql.DB
	Log     log.Log
	Config  config.Config
	ConnStr string
	ctx     context.Context // background context
	cancel  func()          // cancel background context
}

func NewDB(cfg config.Config, log log.Log) RDB {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%v/%s?sslmode=disable",
		cfg.RDB.Username, cfg.RDB.Password, cfg.RDB.Host, cfg.RDB.Port, cfg.RDB.Database)
	db := &rDB{
		ConnStr: dsn,
		Log:     log,
		Config:  cfg,
	}
	db.ctx, db.cancel = context.WithCancel(context.Background())
	return db
}

// Open opens the database connection.
func (db *rDB) Open() error {
	// Ensure a ConnStr is set before attempting to open the database.
	if db.ConnStr == "" {
		return errors.New("empty connection string")
	}

	pool, err := sql.Open("pgx", db.ConnStr)
	if err != nil {
		db.Log.Error("Open Postgres connection pool", "error", err)
		return err
	}
	db.Pool = pool
	db.Log.Info("Open Postgres connection pool")

	duration, err := time.ParseDuration(db.Config.RDB.MaxIdleConnectionLifetime)
	db.Pool.SetMaxOpenConns(db.Config.RDB.MaxOpenConnections)
	db.Pool.SetMaxIdleConns(db.Config.RDB.MaxIdleConnections)
	db.Pool.SetConnMaxLifetime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.Pool.PingContext(ctx)
	if err != nil {
		db.Log.Error("Ping Postgres", "error", err)
		return err
	}
	db.Log.Info("Ping Postgres")

	return nil
}

// Close closes the database connection.
func (db *rDB) Close() error {
	// Cancel background context.
	db.cancel()

	// Close database.
	if db.Pool == nil {
		return nil
	}

	err := db.Pool.Close()
	if err != nil {
		db.Log.Error("Close Postgres", "error", err)
		return err
	}
	return nil
}

func (db *rDB) NewRepo() *repository.Queries {
	return repository.New(db.Pool)
}

func (db *rDB) NewRepoWithTx() (*repository.Queries, *sql.Tx, error) {
	tx, err := db.Pool.Begin()
	if err != nil {
		return nil, nil, err
	}
	repo := repository.New(db.Pool).WithTx(tx)
	return repo, tx, nil
}
