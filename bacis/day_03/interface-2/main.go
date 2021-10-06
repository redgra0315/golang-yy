package main

import "fmt"

// 类型断言

type Inter interface {
	Ping()
	Pang()
}
type Anter struct {
	Inter
	//String()
}
type St struct {
	Name string
}

func (St) Ping() {
	fmt.Println("Ping")
}
func (*St) Pang() {
	fmt.Println("Pang")
}

func main() {
	st := &St{"anders"}
	var i interface{} = st
	if o, ok := i.(Inter); ok {
		o.Ping()
		o.Pang()
	}
	//if p, ok := i.(Anter); ok {
	//	p.String
	//}
	if s, ok := i.(*St); ok {
		s.Ping()
		s.Pang()

		//s := i.(*St)
		fmt.Printf("%s\n", s.Name)
	}
}
