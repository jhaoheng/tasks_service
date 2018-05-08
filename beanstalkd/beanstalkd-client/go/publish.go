package main

import (
  "fmt"
  "github.com/kr/beanstalk"
  "time"
)

func main() {

  conn, _ := beanstalk.Dial("tcp", "10.0.1.7:11300")

  id, err := conn.Put([]byte("myjob"), 1, 0, time.Minute)
  if err != nil {
    panic(err)
  }
  fmt.Println("job", id)
}
