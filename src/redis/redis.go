package redis

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/namsral/flag"
	"log"
	"time"
)

var redisAddr = flag.String("redis_address", "", "address of redis pub/sub instance")

type Client struct {
	redis *redis.Client
}

func NewClient() *Client {
	if redisAddr == nil {
		*redisAddr = "127.0.0.1:6379"
	}

	return &Client{
		redis: redis.NewClient(&redis.Options{
			Addr: *redisAddr,
		}),
	}
}

func (c Client) GetRedis() *redis.Client {
	if c.redis == nil {
		log.Fatalln("client not initialised")
	}

	return c.redis
}

type setOptions struct {
	exp time.Duration
}

type SetOptions func(*setOptions)

func WithExpiration(d time.Duration) SetOptions {
	return func(o *setOptions) {
		o.exp = d
	}
}

func (c Client) Set(ctx context.Context, key, val string, opts ...SetOptions) error {
	var o setOptions
	for _, opt := range opts {
		opt(&o)
	}

	return c.redis.Set(ctx, key, val, o.exp).Err()
}
