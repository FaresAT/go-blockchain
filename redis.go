package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type redisDB struct {
	ctx context.Context
	rdb *redis.Client
}

func NewRedis() *redisDB {
	return &redisDB{
		ctx: context.Background(),
		rdb: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}
}
