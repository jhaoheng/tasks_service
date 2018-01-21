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

  Set("Person", jsonData, 0)
  Get("Person")
}

/*
Set ...
*/
func Set(key string, value interface{}, expiration time.Duration) {
  statusCmd := redisClient.Set(key, value, expiration)
  log.Println("set : ", statusCmd)
}

/*
Get ...
*/
func Get(key string) (val interface{}) {
  statusCmd := redisClient.Get(key)
  log.Println("get : ", statusCmd)

  return
}

/*
flushAll ...
*/
func flushAll() {
  redisClient.FlushAll()
}
