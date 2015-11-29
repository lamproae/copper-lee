package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":1202"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp4", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn *net.TCPConn) {
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(100000000000))
	for {
		var buf [512]byte
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}

		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
