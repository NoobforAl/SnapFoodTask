package db

import (
	"FoodAdder/models"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Redis *redis.Client
var Sql *gorm.DB

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

	if dsn := os.Getenv("URL_MYSQL"); dsn != "" {
		Sql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else {
		Sql, err = gorm.Open(sqlite.Open("testdv.db"), &gorm.Config{})
	}

	if err != nil {
		log.Fatal(err)
	}

	Sql.AutoMigrate(&models.Food{})
}
