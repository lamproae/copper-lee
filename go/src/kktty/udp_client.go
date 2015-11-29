package main

import (
	//	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid input")
		panic("Invalid input")
	}

	addr, err := net.ResolveUDPAddr("udp", ":"+os.Args[2])
	if err != nil {
		fmt.Println(err)
		panic("errrrrrrrrrrrrrrrrrrrrr")
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println(err)
		panic("errrrrrrrrrrrrrrrrrrrrr")
	}

	defer conn.Close()
	for {
		conn.Write([]byte("hello world"))
		// Since that no connection has been established, we cannot use WriteToUDP here. */
		//conn.WriteToUDP([]byte("hello world"), addr)
		data := make([]byte, 1024)
		_, addr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(addr.String())
		fmt.Println(string(data))
		conn.WriteToUDP(data, addr)
	}
}
