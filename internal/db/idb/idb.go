package idb

import "github.com/redis/go-redis/v9"

type IDB interface {
	Open() error
	Close() error
	NewRepo() *redis.Client
}
