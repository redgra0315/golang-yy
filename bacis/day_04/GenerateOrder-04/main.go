package main

import (
	"fmt"
	"sync"
)

// 工作任务
type task struct {
	begin  int
	end    int
	result chan<- int
}

// 执行结果写入result

func (t *task) do() {
	sum := 0
	for i := t.begin; i < t.end; i++ {
		sum += i
	}
	t.result <- sum
}

// 构建task并写入task通道

func InitTask(taskchan chan<- task, r chan int, p int) {

	qu := p / 10
	mod := p % 10
	high := qu * 10
	for j := 0; j < qu; j++ {
		b := 10*j + 1
		e := 10 * (j + 1)
		task := task{
			begin:  b,
			end:    e,
			result: r,
		}
		taskchan <- task
	}
	if mod != 0 {
		tak := task{
			begin:  high + 1,
			end:    p,
			result: r,
		}
		taskchan <- tak

	}
	close(taskchan)
}

// 等待每个task运行完，关闭结果通道

func DistributeTask(taskchan <-chan task, wait *sync.WaitGroup, result chan int) {
	for v := range taskchan {
		wait.Add(1)
		go ProcessTask(v, wait)
	}
	wait.Wait()
	close(result)
}
func ProcessTask(t task, wait *sync.WaitGroup) {
	t.do()
	wait.Done()
}
func ProcessResult(resultchan chan int) int {
	sum := 0
	for r := range resultchan {
		sum += r
	}
	return sum
}
func main() {
	taskchan := make(chan task, 10)
	resulrchan := make(chan int, 10)
	wait := &sync.WaitGroup{}
	go InitTask(taskchan, resulrchan, 200)
	go DistributeTask(taskchan, wait, resulrchan)
	sum := ProcessResult(resulrchan)
	fmt.Println("sum=", sum)
}
