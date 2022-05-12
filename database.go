package main

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Database struct {
	Client *redis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis database")
	CTX    = context.TODO()
)

func NewDB(address string) (*Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if err := client.Ping(CTX).Err(); err != nil {
		return nil, err
	}
	return &Database{
		Client: client,
	}, nil
}

func initRouter(database *Database) *gin.Engine {
	r := gin.Default()
	return r
}
