package main

import (
	"github.com/jacobmiller22/servers/go/tcp/tcpserver"
)

func main() {
	server := tcpserver.NewTCPServer()

	server.StartTCPServer(3000)
}
