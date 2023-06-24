package main

import "muazwzxv/distributedCache/cache"


func main() {

  opts := ServerOptions{
    ListenAddr: "3000",
    IsLeader: true,
  }

  _ = NewServer(opts, cache.NewInMemoryCache())
}

