package database

import (
	"database/sql"
	ce "server/internal/core/error"
)

type Database interface {
	Open() ce.CoreError
	Close() ce.CoreError
	GetPool() *sql.DB
}
