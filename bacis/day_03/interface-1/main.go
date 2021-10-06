package main

import "fmt"

type Printer interface {
	Print()
}
type S struct{}

func (s S) Print() {
	fmt.Println("hello world!")
}
func main() {
	var i Printer
	// 必须初始化，才能使用
	i = S{}
	i.Print()
}
