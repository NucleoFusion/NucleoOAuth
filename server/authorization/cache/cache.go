package cache

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func ConnectToRedis() *redis.Client {
	godotenv.Load(".env")

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Username: "default",
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	return rdb
}
