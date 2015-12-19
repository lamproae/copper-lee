package main

import (
	//	"database/sql"
	//	"encoding/json"
	"fmt"
	//	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	//"os"
	"strconv"
	"strings"
)

type StockInfo struct {
	Name   string
	Code   string
	Date   string
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume uint64
	Adj    float64
}

func httpGet() {
	req, err := http.Get("http://table.finance.yahoo.com/table.csv?s=600000.SS")
	checkError(err)

	result, err := ioutil.ReadAll(req.Body)
	checkError(err)

	defer req.Body.Close()
	//fmt.Println(string(result))

	str := strings.Split(string(result), "\n")
	for i, r := range str {
		//fmt.Println(r)
		if i == 0 {
			continue
		}

		if len(r) == 0 {
			continue
		}

		token := strings.Split(r, ",")
		var stock StockInfo
		stock.Date = token[0]
		stock.Open, _ = strconv.ParseFloat(token[1], 64)
		stock.High, _ = strconv.ParseFloat(token[2], 64)
		stock.Low, _ = strconv.ParseFloat(token[3], 64)
		stock.Close, _ = strconv.ParseFloat(token[4], 64)
		stock.Volume, _ = strconv.ParseUint(token[5], 10, 64)
		stock.Adj, _ = strconv.ParseFloat(token[6], 64)
		fmt.Println(stock)
	}
}

func main() {
	httpGet()
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
