package main

import (
  "fmt"
  "github.com/kr/beanstalk"
  "time"
)

func main() {

  conn, _ := beanstalk.Dial("tcp", "10.0.1.7:11300")

  id, body, err := conn.Reserve(5 * time.Second)
  if err != nil {
    panic(err)
  }
  fmt.Println("job", id)
  fmt.Println(string(body))

  conn.Delete(id)
}
