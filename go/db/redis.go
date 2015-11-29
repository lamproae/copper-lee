package main

import (
	"github.com/astaxie/goredis"
	"fmt"
)

func main() {
	var client goredis.Client
	client.Set("a", []byte("hello"))
	val, _ := client.Get("a")
	fmt.Println(string(val))
	client.Del("a")

	vals := []string{"a", "b", "c", "d", "e"}
	for _, v := range vals {
		client.Rpush("I", []byte(v))
	}

	dbvals, _ := client.Lrange("I", 0, 4)
	for i, v := range dbvals {
		fmt.Println(i, ":", string(v))
	}

	client.Del("I")
}

