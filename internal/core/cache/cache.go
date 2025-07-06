package cache

import ce "server/internal/core/error"

type Cache interface {
	Open() ce.CoreError
	Close() ce.CoreError
}
