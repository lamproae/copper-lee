package main

import (
	"fmt"
	"os"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

type Motion struct {
	Happy bool
	Sad   bool
	Scale []int
}

func main() {
	name := "liwei\n"
	temp := "Hello , {{/* I am the special one! */}}, Value is commented\n"
	temps := "This is {{.Count}} sweaters of {{.Material}}\n"

	sweaters := Inventory{"wool", 10}

	tpl1, _ := template.New("test").Parse("Hello, {{.}}")
	fmt.Println(tpl1.Name())
	tpl2, _ := template.New("test1").Parse(temp)
	fmt.Println(tpl2.Name())
	tpl3, _ := template.New("test2").Parse(temps)
	fmt.Println(tpl3.Name())

	tpl1.Execute(os.Stdout, name)
	tpl2.Execute(os.Stdout, name)
	tpl3.Execute(os.Stdout, sweaters)
	tpl4, _ := template.ParseFiles("./tpl1.txt")
	tpl4.Execute(os.Stdout, sweaters)
	tpl5, _ := template.ParseGlob("./*.txt")
	tpl5.Execute(os.Stdout, sweaters)

	tpl6, _ := template.New("test6").Parse("This is the 6 template test {{if .}} I'm happy {{end}}\n")
	tpl6.Execute(os.Stdout, "1")
	tpl6.Execute(os.Stdout, nil)
	tpl6.Execute(os.Stdout, 1)
	tpl6.Execute(os.Stdout, 0)

	tpl7, _ := template.New("test7").Parse("This is the 7 template test {{if .Happy}} I'm very Happy {{else if .Sad}} I'm very sad {{end}}\n")
	tpl7.Execute(os.Stdout, &Motion{true, true, nil})
	tpl7.Execute(os.Stdout, Motion{false, true, nil})
	tpl7.Execute(os.Stdout, Motion{false, false, nil})

	var motion Motion
	motion.Happy = true
	motion.Sad = true
	motion.Scale = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(motion)
	tpl8, _ := template.New("test8").Parse("This is the 8 template test {{range .Scale}} <p> {{.}} </p> {{end}}\n")
	tpl8.Execute(os.Stdout, &motion)

	motion2 := Motion{Happy: true, Sad: false, Scale: []int{}}
	tpl9, _ := template.New("test9").Parse("This is the 9 template test {{range .Scale}} <p> {{.}} </p> {{else}} NO element exist !!{{end}}\n")
	tpl9.Execute(os.Stdout, &motion2)

	tpl10, _ := template.New("test10").Parse("This is the 10 template test {{with .Count}} {{.}} {{end}}\n")
	tpl10.Execute(os.Stdout, Inventory{Count: 0, Material: "Gold"})
	tpl10.Execute(os.Stdout, Inventory{Count: 100, Material: "Gold"})

	tpl11, _ := template.New("test11").Parse("This is the 11 template test {{with .Count}} {{.}} {{else}} Empty valuse !!{{end}}\n")
	tpl11.Execute(os.Stdout, Inventory{Count: 0, Material: "Gold"})
	tpl11.Execute(os.Stdout, Inventory{Count: 100, Material: "Gold"})

	tplf, _ := template.ParseFiles("tpl3.txt")
	tplf.Execute(os.Stdout, Motion{Happy: true, Sad: false, Scale: []int{4, 5, 8, 19, 17}})
	tplf.Execute(os.Stdout, Motion{Happy: false, Sad: true, Scale: []int{4, 5, 8, 19, 17}})
	tplf.Execute(os.Stdout, Motion{Happy: false, Sad: false, Scale: []int{4, 5, 8, 19, 17}})
	tplf.Execute(os.Stdout, Motion{Happy: false, Sad: false, Scale: []int{}})
}
