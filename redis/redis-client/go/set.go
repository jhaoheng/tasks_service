package main

import (
  "encoding/json"
  "github.com/go-redis/redis"
  "log"
  "time"
)

/*
redisClient ...
*/
var redisClient *redis.Client

/*
RedisNewClient ...
*/
func init() {
  redisClient = redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "", // no password set
    DB:       0,  // use default DB
  })

  pong, err := redisClient.Ping().Result()
  if err != nil {
    log.Fatal("Redis connect fail. ", err)
  } else {
    log.Println("Redis connection success. ", pong)
  }
  return
}

func main() {
  status := redisClient.FlushAll()
  log.Println(status)

  p := map[string]interface{}{
    "name": "max",
    "age":  1,
  }
  jsonData, _ := json.Marshal(p)

  var expiration time.Duration
  expiration = 10 * 1000000000        // convert to second, this is 10s
  Set("Person", jsonData, expiration) // 0 is keep
}

/*
Set ...
*/
func Set(key string, value interface{}, expiration time.Duration) {
  statusCmd := redisClient.Set(key, value, expiration)
  log.Println("set : ", statusCmd)
}

/*
flushAll ...
*/
func flushAll() {
  redisClient.FlushAll()
}
