package kdb

import ce "server/internal/core/error"

type KeyValStore interface {
	Open() ce.CoreError
	Close() ce.CoreError
}
