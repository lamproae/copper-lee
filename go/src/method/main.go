package main

import (
	"./speaker"
	"fmt"
)

// Cannot define different package under a same directory.
// ---> Package is identified by directory name.
func main() {
	fmt.Println("Hello world")
	speaker.SaySomeThing("Hello speaker")
}
