package main

import "fmt"

func main() {
	var ch = make(chan int)

	go func () {
		ch <- 1
	} ()

	fmt.Println("Hello world\n")
	fmt.Printf("%d\n", <-ch)
}
