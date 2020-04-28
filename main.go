package main

import (
	"net"
	"strconv"
)

const (
	addr    = ""
	port    = 8000
	apiKeyT = ""
)

func main() {
	conStat := servStat{0, 0}
	src := addr + ":" + strconv.Itoa(port)
	listener, err := net.Listen("tcp", src)
	handleServError(err)
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		handleServError(err)
		conStat.connect()

		go handleConnection(conn, conStat)
	}
}
