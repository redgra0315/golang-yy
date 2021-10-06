package main

import "fmt"

const (
	a = 1 << iota
	b = 1 << iota
	c = 3
	d = 1 << iota
	e = 1 << iota
)

var A = 11

type User struct {
	Name string
	Age  int
}

func main() {

	fmt.Printf("Hello,world,NIHAO ")
	fmt.Println(a, b, c, d, e)
	p := &A
	fmt.Println(p)

	andes := User{
		Name: "hangman",
		Age:  18,
	}
	p1 := &andes
	fmt.Println(p1.Name)
	fmt.Println(p1.Age)
}
