package cmd

import (
	"net"
	"sync"
)

type Client struct {
	Name     string
	Conn     net.Conn
	Message  chan string
	Messages []string
}

var (
	mutex       sync.Mutex
	Users       map[string]*Client
	allMessages string
)
