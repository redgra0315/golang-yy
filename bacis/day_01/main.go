package main

import (
	"fmt"
)

func main() {
	done := make(chan int, 10)
	for i := 0; i < cap(done); i++ {
		go func() {
			fmt.Println("hello")
			done <- 1

		}()
	}
	for i := 0; i < cap(done); i++ {
		<-done
	}
}
