package productredis

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host string
	Port int
	DB   int
}

type DB struct {
	client *redis.Client
}

func New(config Config) DB {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.Host, config.Port),
		DB:   config.DB,
	})

	return DB{client: rdb}
}

func (d DB) Client() *redis.Client {
	return d.client
}
