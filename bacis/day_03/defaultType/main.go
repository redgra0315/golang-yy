package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {

	a := Person{
		name: "zhangsan",
		age:  16,
	}
	fmt.Println(a)
	p := new(Person)
	//p.name = "李四"
	//p.age = 19
	fmt.Print(p.name, p.age)
}
