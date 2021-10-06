package main

import "fmt"

type Inter interface {
	Qing()
	Qiang()
}

type St struct{}

func (St) Qing() {
	fmt.Println("Qing")
}
func (*St) Qiang() {
	fmt.Println("Qiang")
}

func main() {
	var st *St = nil
	var it Inter = st

	fmt.Printf("%p\n", st)
	fmt.Printf("%p\n", it)

	if it != nil {
		it.Qiang()
		it.Qing()
	}
}
