package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	/* In  go, When the main functions exits, the program stops. */
	go doSomething()
	fmt.Println("end")
	time.Sleep(20000)
}

func doSomething() {
	fmt.Println("Doing Somthing")
}
