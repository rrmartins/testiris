package main

import (
  "fmt"
  "net"
  "strconv"
)

const (
  SVR_HOST = "localhost"
  SVR_PORT = "9982"
  SVR_TYPE = "tcp"
)

func main(){
  fmt.Println("client is running")
  conn, err := net.Dial(SVR_TYPE, SVR_HOST + ":" + SVR_PORT)
  if err != nil {
    panic(err)
  }
  fmt.Println("send data")
  for i := 0; i < 10; i++ {
    _, err = conn.Write([]byte("message from client "+strconv.Itoa(i)+"\n"))
  }

  defer conn.Close()
}
