package main

import (
  "log"
  // "sync/atomic"
  // "time"

  "github.com/go-redis/redis"
)

func main() {

  redisClient := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "", // no password set
    DB:       0,  // use default DB
  })

  Subscribe(redisClient)
}

/*
Subscribe ...
*/
func Subscribe(client *redis.Client) {

  pubsub := client.Subscribe("__keyevent@0__:expired")
  defer pubsub.Close()

  for {
    // msgi, err := pubsub.ReceiveTimeout(time.Second)
    // if err != nil {
    //   log.Println("error : ", err)
    // } else {
    //   log.Println(msgi)
    // }

    msgi, err := pubsub.ReceiveMessage()
    if err != nil {
      log.Println("error : ", err)
    } else {
      log.Println(msgi)
    }

    // switch msg := msgi.(type) {
    // case *redis.Subscription:
    //   log.Println("subscribed to", msg.Channel)
    // case *redis.Message:
    //   log.Println("received", msg.Payload, "from", msg.Channel)
    //   log.Println(client.Get(msg.Payload))
    // default:
    //   log.Println("waiting...")
    // }
  }
}
