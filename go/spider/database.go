package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	db, err := sql.Open("mysql", "root:leeweop@/sinastock?charset=utf8")
	CheckError(err)
	rows, err := db.Query("select * from perDayStat")
	CheckError(err)
	fmt.Println(rows)
	for rows.Next() {
	}

	defer db.Close()
}
