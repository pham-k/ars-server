package kdb

type KDB interface {
	Open() error
	Close() error
}
