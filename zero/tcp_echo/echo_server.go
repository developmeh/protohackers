package tcp_echo

import (
	"fmt"
	"io"
	"log"
	"net"
)

type Server struct {
	TCP *net.TCPAddr
}

func (server *Server) Run() {
	listener, err := net.ListenTCP("tcp", server.TCP)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Default().Println(fmt.Sprintf("Listening on %s:%d", server.TCP.IP.String(), server.TCP.Port))

	for {
		log.Default().Println("Waiting")
		conn, err := listener.Accept()
		if err != nil {
			log.Default().Println("Error accepting connection")
		}
		go server.echo(conn)
	}
}

// echo is a handler that echos received data
// While this is more verbose than it needs to be to help me understand the
// mechanics of the streams. io.Copy() could be used to simplify this.
// It's optimized for direct copy of the readable and writable streams
// The reality is I will probably need to interact with the data so learning how to create
// a compatibility io.Reader and io.Writer is important.
// Some items of note:
//   - io.EOF is not an error, it is a signal that the stream has ended
//   - the buffer is refreshed on each read
//   - conn.Write(input[:length]) is essential to ensure only the data read is written
//     - in my first version this was omitted, and it failed to meet the RFC 862
//          https://www.rfc-editor.org/rfc/rfc862.html
func (server *Server) echo(conn net.Conn) {
	defer conn.Close()
	log.Default().Println("Handling connection")
	for {
		input := make([]byte, 32768)
		length, err := conn.Read(input)

		if err != nil && err == io.EOF {
			log.Default().Println("Closing Connection EOF")
			break
		}

		log.Default().Println("Received: " + string(input))

		_, err = conn.Write(input[:length])
		if err != nil {
			log.Default().Println("error writing data")
			break
		}

		log.Default().Println("Writing: " + string(input))
	}

	return
}
