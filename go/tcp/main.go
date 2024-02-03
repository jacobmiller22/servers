package main

import (
	"log"
	"syscall"

	"github.com/jacobmiller22/servers/go/tcp/tcpserver"
)

func handleConnection(connfd int) error {
	defer syscall.Close(connfd)
	log.Printf("Accepted connection client")

	syscall.Write(connfd, []byte("Happy response from TCP Server\n"))

	log.Printf("Sent a response to client")
	return nil
}

func main() {
	server := tcpserver.NewTCPServer()

	server.StartTCPServer(3000, handleConnection)
}
