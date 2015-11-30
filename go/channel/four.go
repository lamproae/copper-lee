package main

import (
	"fmt"
	"time"
)

func main() {
	for n := 2; n <= 12; n++ {
		timeStable(n)
	}
}

func timeStable(x int) {
	for i := 1; i <= 12; i++ {
		fmt.Printf("%d x %d = %d\n", i, x, x*i)
		time.Sleep(100 * time.Millisecond)
	}
}
