package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func New(addr, password string, db int) *redis.Client {
	if addr == "" {
		log.Println("[infra/redis] Redis not configured, skipping")
		return nil
	}

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Printf("[infra/redis] Redis ping failed: %v", err)
		return nil
	}

	log.Println("[infra/redis] Redis connected")
	return client
}
