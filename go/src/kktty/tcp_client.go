package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":"+os.Args[2])
	if err != nil {
		fmt.Println(err)
		panic("Resolve error")
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println(err)
		panic("Dial err")
	}

	bio := bufio.NewReader(os.Stdin)
	conn.Write([]byte("Hello TCP"))
	for {
		data := make([]byte, 1024)
		if _, err := conn.Read(data); err != nil {
			fmt.Println(string(data))
			continue
		}
		fmt.Println(string(data))
	again:
		line, _, err := bio.ReadLine()
		if err != nil {
			fmt.Println("Read error")
			panic("read error")
		}
		fmt.Println(string(line))
		if len(line) != 0 {
			conn.Write(line)
		} else {
			goto again
		}
	}
}
