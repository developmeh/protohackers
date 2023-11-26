package main

import (
	"net"
	"zero/tcp_echo"
)

func main() {
	server := tcp_echo.Server{TCP: &net.TCPAddr{IP: net.ParseIP("0.0.0.0"), Port: 9999}}
	server.Run()
}
