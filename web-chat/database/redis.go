package database

import (
	"github.com/go-redis/redis/v7"
	"log"
)

var rdb *redis.Client

func InitClient() (rdb *redis.Client) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatal(err)
		return
	}
	return rdb
}