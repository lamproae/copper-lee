package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Your are who")
}
