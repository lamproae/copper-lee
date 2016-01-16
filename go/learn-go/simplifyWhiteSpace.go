package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

func SimplifyWhitespace(s string) string {
	var buffer bytes.Buffer
	skip := true

	for _, char := range s {
		if unicode.IsSpace(char) {
			if !skip {
				buffer.WriteRune(' ')
				skip = true
			}
		} else {
			buffer.WriteRune(char)
			skip = false
		}
	}

	s = buffer.String()
	if skip && len(s) > 0 {
		s = s[:len(s)-1]
	}

	return s
}

func SimpleSimplifyWhitespace(s string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(s)), " ")
}

func main() {
	s := "            	liwei is a \n hao ren    "
	fmt.Println(s)
	fmt.Println(SimpleSimplifyWhitespace(s))
	fmt.Println(SimplifyWhitespace(s))
}
