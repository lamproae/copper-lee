package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usuage: %s hostmane\n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]

	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}

	fmt.Println("Resovled addres is ", addr.String())

	os.Exit(0)
}
