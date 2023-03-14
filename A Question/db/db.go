package db

import (
	"A-Question/models"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Redis *redis.Client
var Sql *gorm.DB

// init redis database
func init() {

	// first load .env file for setting database
	// this only use for debug mode in production
	// not needed .env file
	// use environment in docker-compose
	err := godotenv.Load("../.env")
	if err != nil {
		log.Warn("can't load ../.env file! not debug mod!")
		log.Warn("try find with this path ./.env")

		err := godotenv.Load("./.env")
		if err != nil {
			log.Warn("can't load ./.env file! not debug mod!")
		}
	}

	// setup new redis client
	Redis = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	// check database is ok or not
	_, err = Redis.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
}

// init sql server
func init() {
	var err error

	// check we have url mysql if we have not
	// run program in sqlite with name db file is testdv.db
	if dsn := os.Getenv("URL_MYSQL"); dsn != "" {
		Sql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else {
		Sql, err = gorm.Open(sqlite.Open("testdv.db"), &gorm.Config{})
	}

	if err != nil {
		log.Fatal(err)
	}

	// migrate all model need for program
	Sql.AutoMigrate(&models.FoodModel{})
}
