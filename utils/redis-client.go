package utils

import "github.com/go-redis/redis/v7"

var RedisClient *redis.Client

func InitRedisClient() {
  client := redis.NewClient(&redis.Options{
    Addr:     "127.0.0.1:6379",
    Password: "", // no password set
    DB:       0,  // use default DB
  })

  _, err := client.Ping().Result()
  if err != nil {
  	panic(err)
  	return
  }
  RedisClient = client
}
