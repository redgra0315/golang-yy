package main

import (
	"fmt"
	"math/rand"
)

func GenerateInt(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Label:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Label

			}
		}
		close(ch)
	}()
	return ch
}
func main() {

	done := make(chan struct{})
	ch := GenerateInt(done)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
}
