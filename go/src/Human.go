package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string 
}

type Employee struct {
	Human
	company string
}

type Sicentist struct {
	Human
	area string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (s *Sicentist) SayHi() {
	s.Human.SayHi()
	fmt.Printf("I am a %s scientist!\n", s.area)
}

func main() {
	mark := Student{Human{"liwei", 20, "13111111111"}, "XiDian"}
	sam := Employee{Human{"Sam", 39, "13211112222"}, "Dasan"}
	copper := Sicentist{Human{"Copper", 40, "18211224433"}, "Physicist"}

	mark.SayHi()
	sam.SayHi()
	copper.SayHi()
}
