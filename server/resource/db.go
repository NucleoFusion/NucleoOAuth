package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func ConnectToPG() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", os.Getenv("PSQL_URI"))
	if err != nil {
		return db, err
	}

	fmt.Println("CONNECTED")

	return db, nil

}

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
