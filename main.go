package main

import (
	"log"
	"muazwzxv/distributedCache/cache"
	"net"
	"time"
)


func main() {

  /**
    use http library for this?
    - easier parsing mechanism for payloads 
    - easier response handler
    - use REST format to talk to cache?
  */
  opts := ServerOptions{
    ListenAddr: ":3000",
    IsLeader: true,
  }

  svr := NewServer(opts, cache.NewInMemoryCache())

  go Pinger()

  if err := svr.Start(); err != nil {
    log.Printf("Failed to start appliation: %s \n", err)
  }
}

func Pinger() {
  time.Sleep(time.Second * 2)

  conn, err := net.Dial("tcp", ":3000")
  if err != nil {
    log.Fatal(err)
  }

  _, err = conn.Write([]byte("SET Sum Shit 1000"))
  if err != nil {
    log.Printf("Failed to write to client: %s", err)
  }
}

