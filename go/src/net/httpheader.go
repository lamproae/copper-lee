package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usuage: ", os.Args[0], "host")
		os.Exit(1)
	}

	url := os.Args[1]
	response, err := http.Head(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	fmt.Println(response.Status)
	for k, v := range response.Header {
		fmt.Println(k+": ", v)
	}

	os.Exit(0)
}
