package db

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

var Redis *redis.Client

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("can't load .env file! not debug mod!")
	}

	Redis = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err = Redis.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
}
