package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//	"html/template"
	"text/template"
	//"log"
	//"encoding/json"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
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

func mainPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		db, err := sql.Open("mysql", "kkkmmu:Lee123!!!@/yahoo?charset=utf8")
		checkErorr(err)
		defer db.Close()
		rows, err := db.Query("select * from ss600016")
		var count int = 0
		for rows.Next() {
			count++
		}

		data := make([]StockInfo, count)
		var i int = 0

		rows, err = db.Query("select * from ss600016")
		for rows.Next() {
			var stock StockInfo
			err = rows.Scan(&stock.Date, &stock.Open, &stock.High, &stock.Low, &stock.Close, &stock.Volume, &stock.Adj)
			data[i] = stock
			i++
		}

		//js, err := json.Marshal(data)
		//fmt.Println(string(js))

		fmt.Println("path", r.URL.Path)
		t, _ := template.ParseFiles("./html/index.html")
		t.Execute(w, data)
	} else {

	}
}

func assetsLoad(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == "GET" {
		//	w.Header().Set("Content-type", "text/html; charset=utf-8")
		w.Header().Set("Content-Type", "application/javascript")
		fmt.Println("path: %s\n", r.URL.Path)
		if m, _ := regexp.MatchString("^/assets", r.URL.Path); m {
			//file := path.Base(r.URL.Path)
			file := "." + r.URL.Path
			ext := path.Ext(r.URL.Path)
			fmt.Println(ext)
			if ext == ".svg" {
				fmt.Println(file)
				f, err := os.Open(file)
				checkErorr(err)
				io.Copy(w, f)
			} else {
				//file = "./assets/js/amstockcharts/amcharts/" + file
				//file = "./assets/js/amcharts/amcharts/" + file
				fmt.Println(file)
				t, _ := template.ParseFiles(file)
				t.Execute(w, nil)
			}
		}
	}
}

func main() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/assets/", assetsLoad)

	err := http.ListenAndServe("192.168.1.102:8880", nil)
	//err := http.ListenAndServe("10.71.1.77:8880", nil)
	checkErorr(err)
}

func checkErorr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
