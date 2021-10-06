package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

//通知退出机制

func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Lable
			}
		}
	}()
	return ch
}
func main() {
	done := make(chan struct{})
	ch := GenerateIntA(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// 发送通知，告诉生产者停止生产

	close(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("NumGoroutine", runtime.NumGoroutine())
}
