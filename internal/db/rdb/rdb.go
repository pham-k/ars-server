package rdb

import (
	"database/sql"
	"server/internal/repository"
)

type RDB interface {
	Open() error
	Close() error
	NewRepo() *repository.Queries
	NewRepoWithTx() (*repository.Queries, *sql.Tx, error)
}
