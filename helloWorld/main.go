package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net"
)

func main() {
  conn, err := Dial.Listen("tcp", "localhost:8080")
  if err != nil {
    log.Panic(err)
  }
  defer conn.Close()

  for {
    conn, err := li.Accept()
    if err != nil {
      log.Println(err)
    }
    io.WriteString(conn, "\nHello from TCP server\n")
    fmt.Fprintln(conn, "How is your day?")
    fmt.Fprintf(conn, "%v", "Well, I hope!")

    conn.Close()
  }

}
