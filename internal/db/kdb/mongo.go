package kdb

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"server/internal/config"
	"server/internal/log"
	"time"
)

type kDB struct {
	Pool *mongo.Client
	Log  log.Log
}

func NewKDB(cfg config.Config, log log.Log) KDB {
	pool, _ := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))

	return &kDB{
		Pool: pool,
		Log:  log,
	}
}

func (db *kDB) Open() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := db.Pool.Ping(ctx, readpref.Primary())
	if err != nil {
		db.Log.Error("Ping Mongo", "error", err)
		return err
	}

	db.Log.Info("Ping Mongo")
	return nil
}

func (db *kDB) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := db.Pool.Disconnect(ctx)
	if err != nil {
		db.Log.Error("Close Mongo", "error", err)
		return err
	}

	db.Log.Info("Close Mongo")
	return nil
}
