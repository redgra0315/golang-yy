package main

import "fmt"

type SliceInt []int

func (s SliceInt) Sum() int {
	sum := 0
	for _, i := range s {
		sum += i
	}
	return sum
}

type Map map[string]string

func (m Map) Print() {
	for _, key := range m {
		fmt.Println(key)
	}
}

type MyInt int

func main() {

	s := SliceInt{1, 2, 3, 4, 54, 65}
	b := s.Sum()
	fmt.Println(b)

	a := MyInt(10)
	e := MyInt(20)
	c := a + e
	d := a * e
	fmt.Printf("%d\n", c)
	fmt.Printf("%d\n", d)
}
