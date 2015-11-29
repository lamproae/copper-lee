package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Current pid is: ", os.Getpid())
	fmt.Println("Current ppid is: ", os.Getppid())
	fmt.Println("Currnet app name is: ", os.Args[0])
	fmt.Println("Current eniron is: ", os.Environ())
	wd, _ := os.Getwd()
	fmt.Println("Current working directory is: ", wd)
	hostname, _ := os.Hostname()
	fmt.Println("Current hostname is: ", hostname)
	fmt.Println("Current PATH is: ", os.Getenv("PATH"))
	fmt.Println("Current uid is: ", os.Getuid())
	fmt.Println("Current gid is: ", os.Getgid())
	groups, _ := os.Getgroups()
	fmt.Println("Current Group is: ", groups)

	file, err := os.Open(os.Args[0] + ".go")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	data := make([]byte, 1024)
	for {
		_, err := file.Read(data)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		//fmt.Println(count, data)
		fmt.Println(string(data))
	}

}
