package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type StockInfo struct {
	Name   string  `json:"name"`
	Code   string  `json:"code"`
	Date   string  `json:"date"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume uint64  `json:"volume"`
	Adj    float64 `json:"adj"`
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	db, err := sql.Open("mysql", "kkkmmu:Lee123!!!@/yahoo?charset=utf8")
	CheckError(err)
	rows, err := db.Query("select * from ss600016")
	CheckError(err)
	fmt.Println(rows)
	var stock StockInfo
	for rows.Next() {
		err = rows.Scan(&stock.Date, &stock.Open, &stock.High, &stock.Low, &stock.Close, &stock.Volume, &stock.Adj)
		js, err := json.Marshal(stock)
		CheckError(err)
		fmt.Println(string(js))
	}

	defer db.Close()
}
