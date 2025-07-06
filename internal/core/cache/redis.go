package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"server/internal/core/config"
	ce "server/internal/core/error"
	"time"
)

type cache struct {
	Client *redis.Client
	logger *zap.Logger
	Config config.Config
}

func New(cfg config.Config) Cache {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		//Password: cfg.Cache.Password,
		DB: cfg.IDB.Database,
	})

	return &cache{
		Client: client,
		logger: zap.L(),
		Config: cfg,
	}
}

func (c cache) Open() ce.CoreError {
	ctx := context.Background()

	if redisErr := c.Client.Ping(ctx).Err(); redisErr != nil {
		err := ce.New("ping redis", ce.WithErr(redisErr))
		err.AddOp("open redis")
		return err
	}

	return nil
}

func (c cache) Close() ce.CoreError {
	if c.Client == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if redisErr := c.Client.Shutdown(ctx).Err(); redisErr != nil {
		err := ce.New("shutdown redis", ce.WithErr(redisErr))
		err.AddOp("shutdown redis")
		return err
	}

	return nil
}
