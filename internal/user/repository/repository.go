package repository

import (
	"context"
	"database/sql"
	"go.uber.org/zap"
	"server/internal/core/database"
	"server/internal/user/repository/sgen"
)

type Repository interface {
	NewQuery() *sgen.Queries
	NewQueryTx() (*sgen.Queries, *sql.Tx, error)
	GetUsers(ctx context.Context) (*[]sgen.User, error)
}

type repository struct {
	logger *zap.Logger
	RDB    database.Database
}

func NewRepository(rDB database.Database) Repository {
	return &repository{
		logger: zap.L(),
		RDB:    rDB,
	}
}

func (repo *repository) NewQuery() *sgen.Queries {
	return sgen.New(repo.RDB.GetPool())
}

func (repo *repository) NewQueryTx() (*sgen.Queries, *sql.Tx, error) {
	pool := repo.RDB.GetPool()
	tx, err := pool.Begin()
	if err != nil {
		return nil, nil, err
	}
	query := sgen.New(pool).WithTx(tx)
	return query, tx, nil
}

func (repo *repository) GetUsers(ctx context.Context) (*[]sgen.User, error) {
	//TODO implement me
	panic("implement me")
}
