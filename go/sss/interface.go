package main

import (
	"fmt"
)

type intInterface struct {
}

type stringInterface struct {
}

func (number intInterface) Add(a int, b int) int {
	return a + b
}

func (text stringInterface) Add(a string, b string) string {
	return a + b
}

func main() {
	number := new(intInterface)
	fmt.Println(number.Add(1, 2))

	text := new(stringInterface)
	fmt.Println(text.Add("this old man", " he played one"))
}
