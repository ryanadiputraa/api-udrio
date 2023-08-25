package database

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	rdb *redis.Client
}

type Redis interface {
	Set(ctx context.Context, key string, value interface{}, expiresTime time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}

func InitRedis(dsn string) (Redis, error) {
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(opt)
	return &RedisClient{rdb: rdb}, nil
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiresTime time.Duration) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = r.rdb.Set(ctx, key, jsonData, expiresTime).Err()
	return err
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, err
}
