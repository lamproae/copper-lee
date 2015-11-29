package main

import (
	"bytes"
	"encoding/asn1"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usuage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	conn, err := net.Dial("tcp", service)
	checkError(err)
	result, err := readFully(conn)
	checkError(err)

	var daytime time.Time
	_, err1 := asn1.Unmarshal(result, &daytime)
	checkError(err1)

	fmt.Println("After marshal/unmarshal: ", daytime.String())

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	for {
		var buf [512]byte
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
