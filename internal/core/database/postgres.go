package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
	configPkg "server/internal/core/config"
	ce "server/internal/core/error"
	"time"
)

type database struct {
	Pool *sql.DB

	logger           *zap.Logger
	config           configPkg.Config
	connectionString string
	ctx              context.Context // background context
	cancel           func()          // cancel background context
}

func New(config configPkg.Config) Database {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%v/%s?sslmode=disable",
		config.RDB.Username, config.RDB.Password, config.RDB.Host, config.RDB.Port, config.RDB.Database)
	db := &database{
		connectionString: dsn,
		logger:           zap.L(),
		config:           config,
	}
	db.logger.Warn(fmt.Sprintf("%+v", db.config.RDB))
	db.ctx, db.cancel = context.WithCancel(context.Background())
	return db
}

func (db *database) Open() ce.CoreError {
	if db.connectionString == "" {
		err := ce.New("empty connection string")
		err.AddOp("open postgres connection")
		return err
	}

	pool, sqlErr := sql.Open("pgx", db.connectionString)
	if sqlErr != nil {
		err := ce.New("open postgres connection", ce.WithErr(sqlErr))
		err.AddOp("open postgres connection")
		return err
	}
	db.Pool = pool

	duration, timeErr := time.ParseDuration(db.config.RDB.MaxIdleConnectionLifetime)

	if timeErr != nil {
		db.logger.Warn(db.config.RDB.MaxIdleConnectionLifetime)
		db.logger.Error("error", zap.Error(timeErr))
		err := ce.New("parse MaxIdleConnectionLifetime config", ce.WithErr(sqlErr))
		err.AddOp("open postgres connection")
		return err
	}
	db.Pool.SetMaxOpenConns(db.config.RDB.MaxOpenConnections)
	db.Pool.SetMaxIdleConns(db.config.RDB.MaxIdleConnections)
	db.Pool.SetConnMaxLifetime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sqlErr = db.Pool.PingContext(ctx)
	if sqlErr != nil {
		err := ce.New("ping postgres", ce.WithErr(sqlErr))
		err.AddOp("open postgres connection")
		return err
	}

	return nil
}

func (db *database) Close() ce.CoreError {
	db.cancel()

	if db.Pool == nil {
		return nil
	}

	sqlErr := db.Pool.Close()
	if sqlErr != nil {
		err := ce.New("close postgres connection", ce.WithErr(sqlErr))
		err.AddOp("close postgres connection")
		return err
	}
	return nil
}

func (db *database) GetPool() *sql.DB {
	return db.Pool
}
