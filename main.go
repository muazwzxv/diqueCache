package main

import (
	"log"
	"muazwzxv/distributedCache/cache"
)


func main() {

  opts := ServerOptions{
    ListenAddr: ":3000",
    IsLeader: true,
  }

  svr := NewServer(opts, cache.NewInMemoryCache())

  if err := svr.Start(); err != nil {
    log.Printf("Failed to start appliation: %s \n", err)
  }
}

