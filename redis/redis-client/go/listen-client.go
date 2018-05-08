package main

import (
  "log"
  "sync/atomic"
  "time"

  "github.com/go-redis/redis"
)

func f(from string) {
  for i := 0; i < 3; i++ {
    log.Println(from, ":", i)
  }
}

func main() {
  ring := RedisClient()

  ring.ForEachShard(func(client *redis.Client) error {
    wrapRedisProcess(client)
    return nil
  })
  for {
    ring.Ping()

    for range time.Tick(1 * time.Second) {
      f("direct")
      go f("goroutine")
    }
  }
}

/*
RedisClient ...
*/
func RedisClient() (client *redis.Ring) {

  client = redis.NewRing(&redis.RingOptions{
    Addrs: map[string]string{
      "shard1": "localhost:6379",
    },
  })
  return
}

/*
Subscribe ...
*/
// func Subscribe(client *redis.Client) {

//   pubsub := client.Subscribe("__keyevent@0__:expired")
//   defer pubsub.Close()

//   for {
//     msgi, err := pubsub.ReceiveTimeout(time.Second)
//     if err != nil {
//       log.Println(err)
//     }
//     switch msg := msgi.(type) {
//     case *redis.Subscription:
//       log.Println("subscribed to", msg.Channel)
//     case *redis.Message:
//       log.Println("received", msg.Payload, "from", msg.Channel)
//     default:
//       log.Println("ppp")
//     }
//   }
// }

func wrapRedisProcess(client *redis.Client) {
  const precision = time.Microsecond
  var count, avgDur uint32

  // use Goroutines
  go func() {
    for range time.Tick(3 * time.Second) {
      n := atomic.LoadUint32(&count)
      dur := time.Duration(atomic.LoadUint32(&avgDur)) * precision
      log.Printf("%s: processed=%d avg_dur=%s\n", client, n, dur)
    }
  }()

  client.WrapProcess(func(oldProcess func(redis.Cmder) error) func(redis.Cmder) error {
    return func(cmd redis.Cmder) error {
      start := time.Now()
      err := oldProcess(cmd)
      dur := time.Since(start)

      const decay = float64(1) / 100
      ms := float64(dur / precision)
      for {
        avg := atomic.LoadUint32(&avgDur)
        newAvg := uint32((1-decay)*float64(avg) + decay*ms)
        if atomic.CompareAndSwapUint32(&avgDur, avg, newAvg) {
          break
        }
      }
      atomic.AddUint32(&count, 1)

      return err
    }
  })
}
