package main

import (
	"fmt"
	"log"
	"net"
	Cmd "net-cat-pac/cmd"
	"os"
)

func main() {
	args := os.Args[1:]
	port := ""
	switch {
	case len(args) > 1:
		{
			fmt.Println("[USAGE]: ./net-cat $port")
			return
		}
	case len(args) == 1:
		{
			port = args[0]
		}
	case len(args) == 0:
		{
			port = "8989"
		}
	}
	server, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer server.Close()
	fmt.Printf("Listening on port %s\n", port)
	fmt.Printf("Command: \"nc localhost %s\"\n", port)

	go Cmd.Broadcast()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		if len(Cmd.Users) == 10 {
			fmt.Fprintln(conn, "Chat is already full")
			conn.Close()
			continue
		}
		user := &Cmd.Client{
			Conn:    conn,
			Message: make(chan string),
		}
		go user.Handle()
		go user.ReadMessages()
	}
}
