package main

import "zero/tcp_echo"

func main() {
	server := tcp_echo.Server{Host: "0.0.0.0", Port: "9999"}
	server.Run()
}
