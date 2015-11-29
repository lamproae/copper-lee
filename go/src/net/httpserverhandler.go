package main

import (
	"net/http"
)

func main() {
	myHandler := http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		rw.WriteHeader(http.StatusNoContent)
	})

	http.ListenAndServe(":8000", myHandler)
}
