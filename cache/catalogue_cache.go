package cache

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis successfully!")
}

func SetCache(key string, value interface{}) {
	jsonData, err := json.Marshal(value)
	if err != nil {
		log.Fatalf("Failed to marshal catalogue: %v", err)
	}
	err = rdb.Set(ctx, key, jsonData, 1*time.Minute).Err()
	if err != nil {
		log.Fatalf("Failed to set cache: %v", err)
	}
}

func GetCache(key string) (json.RawMessage, error) {
	value, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return json.RawMessage(value), nil
}
