package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

func mustInsertTokenBuckets(rdb *redis.Client) {
	regions := []string{
		"BR1",
		"EUN1",
		"EUW1",
		"JP1",
		"KR",
		"LA1",
		"LA2",
		"NA1",
		"OC1",
		"TR1",
		"RU",
		"PH2",
		"SG2",
		"TH2",
		"TW2",
		"VN2",
		"AMERICAS",
		"ASIA",
		"EUROPE",
		"SEA",
	}

	fl := os.Getenv("FAST_API_LIMIT")
	sl := os.Getenv("SLOW_API_LIMIT")

	tokenBucketsHash := []string{
		"fast_bucket", fl,
		"fast_bucket_max", fl,
		"fast_bucket_last_refill", "0",
		"fast_bucket_interval", "1",

		"slow_bucket", sl,
		"slow_bucket_max", sl,
		"slow_bucket_last_refill", "0",
		"slow_bucket_interval", "120",
	}

	for _, v := range regions {
		exists, err := rdb.HExists(context.Background(), fmt.Sprint("api:rate_limits_", strings.ToLower(v)), "fast_bucket").Result()

		if err != nil {
			log.Fatal("failed to get hash", err)
		}

		if !exists {
			_, err := rdb.HSet(context.Background(), fmt.Sprint("api:rate_limits_", strings.ToLower(v)), tokenBucketsHash).Result()

			if err != nil {
				log.Fatal("failed to set rate limit buckets", err)
			}
		}
	}
}

func NewRedisClient() *RedisClient {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URI"))

	if err != nil {
		log.Fatal("invalid redis connection string", err)
	}

	rdb := redis.NewClient(opt)

	mustInsertTokenBuckets(rdb)

	rc := &RedisClient{
		client: rdb,
		ctx:    context.Background(),
	}

	return rc
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

func (rc *RedisClient) GetClient() *redis.Client {
	return rc.client
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
