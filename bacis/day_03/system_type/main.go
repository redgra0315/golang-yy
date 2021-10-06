package main

import "fmt"

type Map map[string]string

func (m Map) Print() {
	for _, key := range m {
		fmt.Println(key)
	}
}

type iMap Map

func (m iMap) Print1() {
	for _, key := range m {
		fmt.Println(key)
	}
}

type slice []int

func (s slice) Print2() {
	for _, key := range s {
		fmt.Println(key)
	}
}

func main() {
	mp := make(map[string]string, 10)
	mp["hi"] = "data"

	var ma Map = mp
	//var im iMap = ma
	ma.Print()
	var i interface {
		Print()
	} = ma
	i.Print()
	s1 := []int{1, 2, 3, 4}
	var s2 slice
	s2 = s1
	s2.Print2()

}
