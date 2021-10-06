package main

import (
	"fmt"
	"runtime"
)

// 不要通过共享内存来通信，而是通过通信来共享内存。

func main() {

	//	创建无缓冲的通道
	ch1 := make(chan struct{})
	go func(i chan struct{}) {
		sum := 0
		for i := 0; i < 100000; i++ {
			sum += i
		}
		fmt.Println(sum)
		//	 写通道
		ch1 <- struct{}{}
	}(ch1)
	//NumGoroutine返回当前存在的goroutine数量
	fmt.Println("NumberGoroutine=", runtime.NumGoroutine())
	<-ch1
	//	创建有缓存的通道
	//ch2:=make(chan int,10)

}
