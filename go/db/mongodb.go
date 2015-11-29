package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Person struct {
	Name string
	Phone string
}

func main() {
	session, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Ale", "+5555555522"},
	&Person{"Cla", "+23234444443"})

	if err != nil {
		panic(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name":"Ale"}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println("Phone:", result.Phone)
}

