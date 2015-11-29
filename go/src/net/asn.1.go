package main

import (
	"encoding/asn1"
	"fmt"
	"os"
)

func main() {
	mdata, err := asn1.Marshal(11113)
	checkError(err)

	fmt.Println("After marshal: ", mdata)
	var n int
	_, err1 := asn1.Unmarshal(mdata, &n)
	checkError(err1)

	fmt.Println("After unmarshar1: ", n)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal errorr: %s", err.Error())
		os.Exit(1)
	}
}
