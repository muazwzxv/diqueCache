package main

import (
	"fmt"
	"log"
	"muazwzxv/distributedCache/cache"
	"net"
)

type ServerOptions struct {
	ListenAddr string

	// used for some consensus protocol?
	IsLeader bool
}

type Server struct {
	ServerOptions

	cache cache.ICacher
}

func NewServer(opts ServerOptions, cache cache.ICacher) *Server {
	return &Server{
		ServerOptions: opts,
		cache:         cache,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return fmt.Errorf("listen err: %s", err)
	}

	log.Printf("Server starting on port [%s] \n", s.ListenAddr)

	for {
		conn, err := ln.Accept()
		if err != nil {
      log.Printf("accept error: %s \n", err)
      continue
		}

    go s.handleFunc(conn)
	}
}

func (s *Server) handleFunc(conn net.Conn) {
  defer func() {
    conn.Close()
  }()

  buf := make([]byte, 2048) 
  for {
    n, err := conn.Read(buf)
    if err != nil {
      log.Printf("conn read error: %s \n", err)
      break
    }

    msg := buf[:n]
    fmt.Println(string(msg))
  }
}
