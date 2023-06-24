package main

import (
	"fmt"
	"log"
	"muazwzxv/distributedCache/cache"
	"net"
	"strconv"
	"strings"
	"time"
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
		fmt.Printf("Bytes: %x \n", msg)
		fmt.Printf("String: %s", string(msg))
	}
}

func (s *Server) HandleCommand(conn net.Conn, rawCmd []byte) {
	var (
		raw   = string(rawCmd)
		parts = strings.Split(raw, " ")
	)

	if len(parts) == 0 {
		// response
		log.Println("Invalid set command")
		return
	}

	cmd := Command(parts[0])
	if cmd == SetCmd {
		if len(parts) != 4 {
			// response
			return
		}

		key := []byte(parts[1])
		value := []byte(parts[2])

    ttl, err := strconv.Atoi(parts[3])
    if err != nil {
      log.Println("Invalid set command")
    }

		payload := SetMessage{
			Key:   key,
			Value: value,
			TTL:   time.Duration(ttl),
		}

    if err := s.HandleSetCmd(conn, payload); err != nil {
      // response
      return 
    }
	}
}

func (s *Server) HandleSetCmd(conn net.Conn, setMsg SetMessage) error {
  log.Println("Handling set command: ", setMsg)
	return nil
}
