package main

import (
	"database/sql"
	//	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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

var ShangHaiSuffix string = "ss"

func httpGet() {
	db, err := sql.Open("mysql", "root:leeweop@/yahoo?charset=utf8")
	checkError(err)

	req, err := http.Get("http://table.finance.yahoo.com/table.csv?s=600000.SS")
	checkError(err)

	result, err := ioutil.ReadAll(req.Body)
	checkError(err)

	defer req.Body.Close()
	//fmt.Println(string(result))

	_, err = db.Query("select * from ss600000")
	if err != nil {
		_, err := db.Query("CREATE TABLE ss600000 (Date char(64) not null primary key, Open double not null, High double not null, Low double not null, Close double not null, Volume bigint(64) not null, Adj double not null)")
		checkError(err)
	}

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
		insert, err := db.Prepare("INSERT ss600000 SET Date=?, Open=?, High=?, Low=?, Close=?, Volume=?, Adj=?")
		checkError(err)

		_, err = insert.Exec(stock.Date, stock.Open, stock.High, stock.Low, stock.Close, stock.Volume, stock.Adj)
		checkError(err)
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
