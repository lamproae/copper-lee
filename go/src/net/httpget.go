package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usuage: ", os.Args[0], "host:port")
		os.Exit(1)
	}

	url := os.Args[1]
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	b, _ := httputil.DumpResponse(response, false)
	fmt.Print(string(b))

	contentTypes := response.Header["Content-Type"]
	if !acceptableCharset(contentTypes) {
		fmt.Println("Cannot handle", contentTypes)
		os.Exit(4)
	}

	reader := response.Body
	for {
		var buf = make([]byte, 512)
		n, err := reader.Read(buf[0:])
		if err != nil {
			os.Exit(0)
		}

		fmt.Print(string(buf[0:n]))
	}
	os.Exit(0)
}

func acceptableCharset(contentTypes []string) bool {
	for _, cType := range contentTypes {
		if strings.Index(cType, "UTF-8") != -1 {
			return true
		} else if strings.Index(cType, "GBK") != -1 {
			return true
		} else if strings.Index(cType, "utf-8") != -1 {
			return true
		}
	}
	return false
}
