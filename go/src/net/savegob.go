package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Person struct {
	Name  Name
	Email []Email
}

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string
	Address string
}

func main() {
	person := Person{
		Name: Name{Family: "Newwarch", Personal: "Jan"},
		Email: []Email{Email{Kind: "home", Address: "liwei@163.com"},
			Email{Kind: "work", Address: "liwei@gmail.com"}}}
	saveGob("Person.gob", person)
}

func saveGob(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}
