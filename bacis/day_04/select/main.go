package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	go func(chan int) {
		for {
			// select 多路复用
			select {
			// 写入通道
			case ch <- 0:
			case ch <- 1:

			}
		}
	}(ch)
	for i := 0; i < 10; i++ {
		// 读取通道信息

		fmt.Println(<-ch)
	}
}
