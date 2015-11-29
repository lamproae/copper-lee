package main

import (
	"fmt"
	"time"
	"database/sql"
	"github.com/astaxie/beedb"
	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Userinfo struct {
	Uid int `PK`
	Username string
	Departname string
	Created time.Time
}

func main() {
	db, err := sql.Open("mysql", "kkkmmu:leeweop@/go?charset=utf8")
	checkErr(err)

	orm := beedb.New(db)
	beedb.OnDebug = true

	var saveone Userinfo
	saveone.Username = "Test Add User"
	saveone.Departname = "Test Add Departname"
	saveone.Created = time.Now()
	orm.Save(&saveone)

	fmt.Println(saveone.Uid)

	add := make(map[string]interface{})
	add["username"] = "beedb"
	add["departname"] = "google"
	add["created"] = "2015-12-11"
	orm.SetTable("userinfo").Insert(add)

	/* This part test failed.
	addslice := make([]map[string]interface{}, 10)
	add2 := make(map[string]interface{})
	add3 := make(map[string]interface{})
	add2["username"] = "yang"
	add2["departname"] = "kaiyi"
	add2["created"] = "2012-11-1"
	add3["username"] = "rong"
	add3["departname"] = "kaiyi"
	add3["created"] =  "2011-11-02"
	addslice = append(addslice, add, add2)

	fmt.Println(addslice)
	//orm.SetTable("userinfo").Insert(addslice)
	*/

	saveone.Username = "Update Username"
	saveone.Departname = "Update Departname"
	saveone.Created = time.Now()
	orm.Save(&saveone)

	t := make(map[string]interface{})
	t["username"] =  "Update Username"
	orm.SetTable("userinfo").SetPK("uid").Where(2).Update(t)

	var user Userinfo
	orm.Where("uid=?", 10).Find(&user)
	fmt.Println(user)

	var user2 Userinfo
	orm.Where(3).Find(&user2)
	fmt.Println(user2)

	var user3 Userinfo
	orm.Where("username == ?", "Update Username").Find(&user3)
	fmt.Println(&user3)

	var allusers []Userinfo
	orm.Where("uid > ?", "3").FindAll(&allusers)
	for v, _ := range allusers {
		fmt.Println(v)
	}

	var allusers2 []Userinfo
	orm.FindAll(&allusers2)
	for v, _ := range allusers2 {
		fmt.Println(v)
	}
	orm.Delete(&saveone)
	orm.Delete(&allusers)
}

