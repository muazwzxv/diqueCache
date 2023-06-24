package main

import "muazwzxv/distributedCache/cache"

type ServerOptions struct {
  ListenAddr string

  // used for some consensus protocol?
  IsLeader bool
}

type Server struct {
  opts ServerOptions 
  
  cache cache.ICacher
}

func NewServer(opts ServerOptions, cache cache.ICacher) *Server {
  return &Server{
    opts: opts,
    cache: cache,
  }
}

