package redis

import (
	"context"
	"finance-tracker/pkg/config"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg config.App) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisConfig.Host + ":" + cfg.RedisConfig.Port,
		Password: cfg.RedisConfig.Password,
		DB:       cfg.RedisConfig.DB,
	})

	// Health check menggunakan PING
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Successfully connected to Redis")
	return client
}

// BlacklistToken menambahkan token ke blacklist dengan TTL
func BlacklistToken(client *redis.Client, token string, expiry time.Duration) error {
	ctx := context.Background()
	return client.Set(ctx, token, "blacklisted", expiry).Err()
}

// IsTokenBlacklisted memeriksa apakah token ada di blacklist
func IsTokenBlacklisted(client *redis.Client, token string) (bool, error) {
	ctx := context.Background()
	val, err := client.Get(ctx, token).Result()
	if err == redis.Nil {
		return false, nil
	}
	return val == "blacklisted", err
}
