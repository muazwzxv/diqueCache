package main

import "time"

type Command string

const (
  SetCmd Command = "SET" 
  GetCmd Command = "GET"
)

type SetMessage struct {
  Key []byte
  Value []byte
  TTL time.Duration
}

