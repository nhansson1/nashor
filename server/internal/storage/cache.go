package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
    client *redis.Client
    ctx context.Context
}

func NewRedisClient() *RedisClient {
    rdb := redis.NewClient(&redis.Options{
        Addr: "cache:6379",
    })

    return &RedisClient{
        client: rdb,
        ctx: context.Background(),
    }
}

func (rc *RedisClient) Set(key, data string, ex time.Duration) {
    if err := rc.client.Set(rc.ctx, key, data, ex).Err(); err != nil {
        fmt.Println("Failed to set value in cache: ", key, data, err)
    }
}

func (rc *RedisClient) Get(key string) (string, error) {
    val, err := rc.client.Get(rc.ctx, key).Result()

    if err != nil {
        return "", err
    }

    return val, nil 
}

func (rc *RedisClient) GetJson(key string, out any) error {
    val, err := rc.client.Get(rc.ctx, key).Result()

    if err != nil {
        return err
    }

    err = json.Unmarshal([]byte(val), out)

    if err != nil {
        return err
    }

    return nil
}
