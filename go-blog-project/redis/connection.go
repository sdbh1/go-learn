package redis

import (
	"log"
	"sdbh/config"

	"github.com/go-redis/redis"
)

func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.BlogConfig.Redis.Addr,
		DB:       config.BlogConfig.Redis.DB,
		Password: config.BlogConfig.Redis.Password,
	})

	_, err := client.Ping().Result()

	if err != nil {
		log.Fatalf("Failed to connect to Redis, got error: %v", err)
	}

	return client
}
