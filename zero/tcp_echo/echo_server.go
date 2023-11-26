package tcp_echo

import (
	"fmt"
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
			// handle error
		}
		go server.handleConnection(conn)
	}
}

func (server *Server) handleConnection(conn net.Conn) {
	log.Default().Println("Handling connection")
	for {
		input := make([]byte, 32768)
		read_size, err := conn.Read(input)
		if read_size == 0 {
			log.Default().Println("connection closed")
			break
		}
		if err != nil {
			log.Default().Println("error reading data")
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
