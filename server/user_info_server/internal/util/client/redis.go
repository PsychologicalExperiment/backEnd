package client

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/grpclog"
)

type RedisInitOpt struct {
	Addr     string
	Password string
	DB       int
}

func InitRedisClient(opt RedisInitOpt) *redis.Client {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     opt.Addr,
		Password: opt.Password,
		DB:       opt.DB,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := redisCli.Ping(ctx).Result()
	if err != nil {
		grpclog.Fatalf("init redis cli failed, error: %+v", err)
		return nil
	}
	return redisCli
}
