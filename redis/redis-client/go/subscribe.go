package main

import (
  "log"
  // "sync/atomic"
  // "time"
  "encoding/json"

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

  // keyspace 是以操作指令相關
  // keyevent 是以事件相關
  // pubsub := cient.Subscribe("__key*@0__:*") // 訂閱 space&event (使用 key) 的所有操作事件(使用 *)
  // pubsub := client.Subscribe("__keyspace@0__:*")
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
      Get(client, msgi.Payload)
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

/*
Get ...
*/
func Get(client *redis.Client, key string) {
  statusCmd := client.Get(key)
  log.Println("key is : ", key, ", get : ", statusCmd)

  jsonEncode, _ := statusCmd.Bytes()

  jsonDecode := map[string]interface{}{}
  json.Unmarshal(jsonEncode, &jsonDecode)
  log.Println("get : ", jsonDecode)
  // if len(jsonDecode) == 0 {
  //   log.Println("No Find value at this key")
  // }
  return
}
