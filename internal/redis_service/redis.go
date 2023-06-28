package redisservice

import (
	"context"
	"log"

	"github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/config"
	"github.com/redis/go-redis/v9"
)

func New(conf config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: conf.RedisAddr(),
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}

	return rdb
}
