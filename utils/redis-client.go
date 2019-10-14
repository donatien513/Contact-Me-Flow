package utils

import "github.com/go-redis/redis/v7"
import "github.com/donatien513/Contact-Me-Flow/config"

var RedisClient *redis.Client

func InitRedisClient() {
  client := redis.NewClient(&redis.Options{
    Addr:     config.RedisAddr,
    Password: config.RedisPass, // no password set
    DB:       0,  // use default DB
  })

  _, err := client.Ping().Result()
  if err != nil {
  	panic(err)
  	return
  }
  RedisClient = client
}
