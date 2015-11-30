package main

import (
	"github.com/gorlla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleIndex)
	router.HandleFunc("/hello", handleHello)
	router.HandleFunc("/goodbye", handleGoodbye)

	router.HandleFunc("/things/{id}", handleThingsRead).Methods("GET")
	router.HandleFunc("/things/{id}", handleThingsUpdate).Methods("PUT")

	http.Handle("/", router)
}

func handleThingsRead(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//TODO: Load thing with ID vars["id"]
}
