package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

var RedisClient *redis.Client

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6479",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
		return
	}

	fmt.Println("Connection status:", pong)
}
