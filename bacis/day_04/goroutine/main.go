package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	go func() {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
		}
		fmt.Println(sum)
		fmt.Println(1 * time.Second)
	}()
	runtime.GOMAXPROCS(4)
	fmt.Println("NumGoroutine", runtime.NumGoroutine())
	//time.Sleep(5 * time.Second)
}
