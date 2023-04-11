package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type RedisClient struct {
	rdb *redis.Client
}

type Redis interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) (string, error)
}

func InitRedis() Redis {
	opt, err := redis.ParseURL(viper.GetString("REDIS_URL"))
	if err != nil {
		panic(fmt.Sprintf("fail to connect to redis: %s", err))
	}

	rdb := redis.NewClient(opt)
	return &RedisClient{rdb: rdb}
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = r.rdb.Set(ctx, key, jsonData, 10*time.Minute).Err()
	return err
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, err
}
