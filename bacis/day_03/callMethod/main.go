package main

import "fmt"

type T struct {
	a int
}

func (t *T) Get() int {
	return t.a
}
func (t *T) Post() int {
	return t.a

}
func (t *T) Delete() int {
	return t.a

}
func (t *T) Put() int {
	return t.a

}
func main() {
	var t = &T{1}

	//a := t.a
	fmt.Println(t.Post())
	fmt.Println(t.Get())
	fmt.Println(t.Post())
	fmt.Println(t.Put())
}
