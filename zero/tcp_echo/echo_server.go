package tcp_echo

import (
	"fmt"
	"io"
	"log"
	"net"
)

type Server struct {
	Host string
	Port string
}

func (server *Server) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server.Host, server.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Default().Println("Listening on " + server.Host + ":" + server.Port)

	for {
		log.Default().Println("Waiting")
		conn, err := listener.Accept()
		if err != nil {
			log.Default().Println("Error accepting connection")
		}
		go server.handleConnection(conn)
	}
}

func (server *Server) handleConnection(conn net.Conn) {
	log.Default().Println("Handling connection")
	for {
		input := make([]byte, 32768)
		_, err := conn.Read(input)

		if err != nil && err == io.EOF {
			log.Default().Println("Closing Connection EOF")
			break
		}

		log.Default().Println("Received: " + string(input))

		_, err = conn.Write(input)
		if err != nil {
			log.Default().Println("error writing data")
			break
		}

		log.Default().Println("Writing: " + string(input))
	}

	return
}
