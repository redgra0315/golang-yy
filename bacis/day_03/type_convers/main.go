package main

import "fmt"

// 类型转换

type Map map[string]string
type iMap Map

func (m Map) Print() {
	for _, key := range m {
		fmt.Println(key)
	}
}

func (m iMap) Print1() {
	for _, key := range m {
		fmt.Println(key)
	}
}

// StringToByte 字符串与字节之间的转换
func StringToByte() {
	s := "hello,世界！"
	var a []byte
	a = []byte(s)
	var b string
	b = string(a)
	var c []rune
	c = []rune(b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

}
func main() {
	mp := make(map[string]string, 10)
	mp["hi"] = "data"

	var ma Map = mp
	//var im iMap = ma
	var im iMap = (iMap)(ma)
	ma.Print()
	var i interface {
		Print()
	} = ma
	i.Print()
	im.Print1()
	StringToByte()
}
