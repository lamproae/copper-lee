package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type msg struct {
	Message string `json:"msg"` // = {"msg":"Hi"}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode(&msg{Message: "Hello Golang"}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
