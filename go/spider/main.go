package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func httpDo() {
	v := url.Values{}
	v.Set("theStockCode", "002269")
	body := ioutil.NopCloser(strings.NewReader(v.Encode()))
	client := http.Client{}

	req, err := http.NewRequest("POST", "http://www.webxml.com.cn/WebServices/ChinaStockWebService.asmx/getStockInfoByCode", body)
	if err != nil {
		fmt.Println("Create new request error!")
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/soap+xml; charset=utf8")
	req.Header.Set("Content-Length: ", "19")

	resp, err := client.Do(req)
	fmt.Println(resp)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read inforamtion errror!")
	}

	fmt.Println(string(data))

	defer resp.Body.Close()
}

func httpGet() {
	//	req, err := http.Get("http://www.webxml.com.cn/WebServices/ChinaStockWebService.asmx/getStockInfoByCode?theStockCode=002269")
	//	req, err := http.Get("http://hq.sinajs.cn/list=002269")
	//Yahoo finance data for history.
	req, err := http.Get("http://table.finance.yahoo.com/table.csv?s=600000.SS")
	//Yahoo finance data for today
	//req, err := http.Get("http://finance.yahoo.com/d/quotes.csv?s=600000.SS")
	//req, err := http.Get("http://table.finance.yahoo.com/table.csv?s=002269.ss")	/* History data. */
	//	req, err := http.Get("http://table.finance.yahoo.com/quotes.csv?s=002269.ss") /* current data. */
	//	req, err := http.Get("http://ichart.yahoo.com/quotes.csv?s=002269.ss") /* current data. */
	//req, err := http.Get("http://ichart.yahoo.com/table.csv?s=002269.ss") /* current data. */
	if err != nil {
		fmt.Println(err.Error())
	}

	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer req.Body.Close()
	fmt.Println(string(result))
}

func main() {
	httpGet()
}
