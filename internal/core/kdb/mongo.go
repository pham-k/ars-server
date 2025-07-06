package kdb

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"go.uber.org/zap"
	"server/internal/core/config"
	ce "server/internal/core/error"
	"time"
)

type keyValStore struct {
	Pool   *mongo.Client
	logger *zap.Logger
}

func New(cfg config.Config) KeyValStore {
	pool, _ := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))

	return &keyValStore{
		Pool:   pool,
		logger: zap.L(),
	}
}

func (db *keyValStore) Open() ce.CoreError {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	mongoErr := db.Pool.Ping(ctx, readpref.Primary())
	if mongoErr != nil {
		err := ce.New("ping mongo", ce.WithErr(mongoErr))
		err.AddOp("open mongo")
		return err
	}

	return nil
}

func (db *keyValStore) Close() ce.CoreError {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	mongoErr := db.Pool.Disconnect(ctx)
	if mongoErr != nil {
		err := ce.New("close mongo", ce.WithErr(mongoErr))
		err.AddOp("close mongo")
		return err
	}

	return nil
}
