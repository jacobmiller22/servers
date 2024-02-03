package tcpserver

import (
	"log"
	"syscall"
)

type TCPServer struct{}

func NewTCPServer() *TCPServer {
	return &TCPServer{}
}

func (s *TCPServer) StartTCPServer(port uint16, handleConnection func(int) error) error {
	log.Printf("Starting TCP Server on port %d\n", port)

	sockaddr := &syscall.SockaddrInet4{
		Port: int(port),
		Addr: [4]byte{127, 0, 0, 1},
	}

	// IPv4 (Didn't find PF_INET), TCP, 0 for protocol for the given type
	sockfd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)

	defer syscall.Close(sockfd)

	if err != nil {
		log.Printf("Error creating socket: %v", err)
		return err
	}
	err = syscall.Bind(sockfd, sockaddr)
	if err != nil {
		log.Printf("Error binding to socket: %v", err)
		return err
	}

	err = syscall.Listen(sockfd, syscall.SOMAXCONN)
	if err != nil {
		log.Printf("Error listening to socket: %v", err)
		return err
	}

	for {
		connfd, _, err := syscall.Accept(sockfd)
		if err != nil {
			log.Printf("Error accepting incoming connection: %v", err)
			return err
		}

		go handleConnection(connfd)

	}
	return nil
}
