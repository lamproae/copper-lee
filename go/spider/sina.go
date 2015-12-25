package main

import (
        "database/sql"
        "encoding/json"
        "fmt"
        _ "github.com/go-sql-driver/mysql"
        "io/ioutil"
        "net/http"
        "os"
        "strconv"
        "strings"
       )

type StockInfo struct {
    Name           string
        Code           string
        TOpeningPrice  float64
        YClosingPrice  float64
        CurrentPrice   float64
        THigh          float64
        TLow           float64
        BBidPrice      float64
        SBidPrice      float64
        TotalCount     uint64
        TotalValue     float64
        BuyOneCount    uint64
        BuyOnePrice    float64
        BuyTwoCount    uint64
        BuyTwoPrice    float64
        BuyThreeCount  uint64
        BuyThreePrice  float64
        BuyFourCount   uint64
        BuyFourPrice   float64
        BuyFiveCount   uint64
        BuyFivePrice   float64
        SellOneCount   uint64
        SellOnePrice   float64
        SellTwoCount   uint64
        SellTwoPrice   float64
        SellThreeCount uint64
        SellThreePrice float64
        SellFourCount  uint64
        SellFourPrice  float64
        SellFiveCount  uint64
        SellFivePrice  float64
        Date           string
        Time           string
}

func httpGet() {
    req, err := http.Get("http://hq.sinajs.cn/list=sh600589")
        if err != nil {
            fmt.Println(err.Error())
        }

    result, err := ioutil.ReadAll(req.Body)
        if err != nil {
            fmt.Println(err.Error())
        }
    defer req.Body.Close()

        fmt.Println(string(result))
        var stock StockInfo

        str := strings.Split(string(result), ",")
        str2 := strings.Split(string(str[0]), "=")
        str3 := str2[0][len(str2[0])-6 : len(str2[0])]
        stock.Code = str3
        stock.Name = str2[1]

        stock.TOpeningPrice, _ = strconv.ParseFloat(string(str[1]), 64)
        stock.YClosingPrice, _ = strconv.ParseFloat(string(str[2]), 64)
        stock.CurrentPrice, _ = strconv.ParseFloat(string(str[3]), 64)
        stock.THigh, _ = strconv.ParseFloat(string(str[4]), 64)
        stock.TLow, _ = strconv.ParseFloat(string(str[5]), 64)
        stock.BBidPrice, _ = strconv.ParseFloat(string(str[6]), 64)
        stock.SBidPrice, _ = strconv.ParseFloat(string(str[7]), 64)
        stock.TotalCount, _ = strconv.ParseUint(string(str[8]), 10, 64)
        stock.TotalValue, _ = strconv.ParseFloat(string(str[9]), 64)
        stock.BuyOneCount, _ = strconv.ParseUint(string(str[10]), 10, 64)
        stock.BuyTwoCount, _ = strconv.ParseUint(string(str[12]), 10, 64)
        stock.BuyThreeCount, _ = strconv.ParseUint(string(str[14]), 10, 64)
        stock.BuyFourCount, _ = strconv.ParseUint(string(str[16]), 10, 64)
        stock.BuyFiveCount, _ = strconv.ParseUint(string(str[18]), 10, 64)
        stock.BuyOnePrice, _ = strconv.ParseFloat(string(str[11]), 64)
        stock.BuyTwoPrice, _ = strconv.ParseFloat(string(str[13]), 64)
        stock.BuyThreePrice, _ = strconv.ParseFloat(string(str[15]), 64)
        stock.BuyFourPrice, _ = strconv.ParseFloat(string(str[17]), 64)
        stock.BuyFivePrice, _ = strconv.ParseFloat(string(str[19]), 64)
        stock.SellOneCount, _ = strconv.ParseUint(string(str[20]), 10, 64)
        stock.SellTwoCount, _ = strconv.ParseUint(string(str[22]), 10, 64)
        stock.SellThreeCount, _ = strconv.ParseUint(string(str[24]), 10, 64)
        stock.SellFourCount, _ = strconv.ParseUint(string(str[26]), 10, 64)
        stock.SellFiveCount, _ = strconv.ParseUint(string(str[28]), 10, 64)
        stock.SellOnePrice, _ = strconv.ParseFloat(string(str[21]), 64)
        stock.SellTwoPrice, _ = strconv.ParseFloat(string(str[23]), 64)
        stock.SellThreePrice, _ = strconv.ParseFloat(string(str[25]), 64)
        stock.SellFourPrice, _ = strconv.ParseFloat(string(str[27]), 64)
        stock.SellFivePrice, _ = strconv.ParseFloat(string(str[29]), 64)
        stock.Date = str[30]
        stock.Time = str[31]

        db, err := sql.Open("mysql", "root:leeweop@/sinastock?charset=utf8")
        CheckError(err)

        fmt.Println(stock)
        action, err := db.Prepare("INSERT stat SET Name=?, Code=?, TOpeningPrice=?, YClosingPrice=?, CurrentPrice=?, THigh=?, TLow =?, BBidPrice=?, SBidPrice=?, TotalCount=?, TotalValue=?, BuyOneCount=?, BuyOnePrice=?, BuyTwoCount=?, BuyTwoPrice=?, BuyThreeCount=?, BuyThreePrice=?, BuyFourCount=?, BuyFourPrice=?, BuyFiveCount=?, BuyFivePrice=?, SellOneCount=?, SellOnePrice=?, SellTwoCount=?, SellTwoPrice=?, SellThreeCount=?, SellThreePrice=?, SellFourCount=?, SellFourPrice=?, SellFiveCount=?, SellFivePrice=?, Date=?, Time=? ")
        CheckError(err)
        _, err = action.Exec(stock.Name, stock.Code, stock.TOpeningPrice, stock.YClosingPrice, stock.CurrentPrice, stock.THigh, stock.TLow, stock.BBidPrice, stock.SBidPrice, stock.TotalCount, stock.TotalValue, stock.BuyOneCount, stock.BuyOnePrice, stock.BuyTwoCount, stock.BuyTwoPrice, stock.BuyThreeCount, stock.BuyThreePrice, stock.BuyFourCount, stock.BuyFourPrice, stock.BuyFiveCount, stock.BuyFivePrice, stock.SellOneCount, stock.SellOnePrice, stock.SellTwoCount, stock.SellTwoPrice, stock.SellThreeCount, stock.SellThreePrice, stock.SellFourCount, stock.SellFourPrice, stock.SellFiveCount, stock.SellFivePrice, stock.Date, stock.Time)
        CheckError(err)

        var ds StockInfo
        rows, err := db.Query("select * from stat")
        CheckError(err)

        for rows.Next() {
            err = rows.Scan(&ds.Name, &ds.Code, &ds.TOpeningPrice, &ds.YClosingPrice, &ds.CurrentPrice, &ds.THigh, &ds.TLow, &ds.BBidPrice, &ds.SBidPrice, &ds.TotalCount, &ds.TotalValue, &ds.BuyOneCount, &ds.BuyOnePrice, &ds.BuyTwoCount, &ds.BuyTwoPrice, &ds.BuyThreeCount, &ds.BuyThreePrice, &ds.BuyFourCount, &ds.BuyFourPrice, &ds.BuyFiveCount, &ds.BuyFivePrice, &ds.SellOneCount, &ds.SellOnePrice, &ds.SellTwoCount, &ds.SellTwoPrice, &ds.SellThreeCount, &ds.SellThreePrice, &ds.SellFourCount, &ds.SellFourPrice, &ds.SellFiveCount, &ds.SellFivePrice, &ds.Date, &ds.Time)
                CheckError(err)
                fmt.Println(ds)
        }

        fmt.Printf("%s", result)
            fmt.Println(stock)
            b, _ := json.Marshal(stock)
            fmt.Println(string(b))

            var tmp StockInfo
            json.Unmarshal(b, &tmp)
            fmt.Println(tmp)
}

func main() {
    httpGet()
}

func CheckError(err error) {
    if err != nil {
        fmt.Println(err.Error())
            os.Exit(-1)
    }
}

//Yahoo finance data for history.
//req, err := http.Get("http://table.finance.yahoo.com/table.csv?s=600000.SS")
//Yahoo finance data for today
//req, err := http.Get("http://finance.yahoo.com/d/quotes.csv?s=600000.SS")
