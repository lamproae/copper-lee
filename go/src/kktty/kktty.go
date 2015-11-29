package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var ProtocolMap map[string]string = make(map[string]string)

func show_usuage() {
	fmt.Println("Usuage:")
	fmt.Println("kktty PROTOCOL IP")
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid Input")
		show_usuage()
		os.Exit(-1)
	}

	if _, ok := ProtocolMap[os.Args[1]]; !ok {
		fmt.Println("Unknown Protocol")
	}
	ipstr := os.Args[2]
	conn, err := net.Dial("tcp", ipstr+":"+ProtocolMap[os.Args[1]])
	if err != nil {
		fmt.Println(err)
	}

	rd := bufio.NewReader(conn)
	for {
		var data []byte = make([]byte, 1024)
		_, err := rd.Read(data)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		fmt.Println(string(data))
	}
}

func init() {
	ProtocolMap["ssh"] = "22"
	ProtocolMap["telnet"] = "23"
}
