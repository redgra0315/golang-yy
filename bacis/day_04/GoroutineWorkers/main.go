package main

import (
	"fmt"
)

// Number 设置工作池的goroutine数目
const Number = 10

// Task 工作任务
type Task struct {
	begin  int
	end    int
	result chan<- int
}

// Do 任务处理：计算begin 到end 的和
// Do 执行结果写入chan result
func (t *Task) Do() {
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += 1
	}
	t.result <- sum
}

// InitTask 初始化待处理task chan

func InitTask(tackchan chan<- Task, r chan int, p int) {
	qu := p / 10
	mod := p % 10
	high := qu * 10
	for j := 0; j < qu; j++ {
		b := 10*j + 1
		e := 10 * (j + 1)
		tak := Task{
			begin:  b,
			end:    e,
			result: r,
		}
		tackchan <- tak
	}
	if mod != 0 {
		tsk := Task{
			begin:  high + 1,
			end:    p,
			result: r,
		}
		tackchan <- tsk
	}
	close(tackchan)
}

// 读取task chan 并发写到worker goroutine 处理，总的数量是哦workers

func DistributeTask(taskchan <-chan Task, workers int, done chan struct{}) {
	for i := 0; i < workers; i++ {
		go ProcessTask(taskchan, done)
	}
}
func ProcessTask(taskchan <-chan Task, done chan struct{}) {
	for t := range taskchan {
		t.Do()
	}
	done <- struct{}{}
}

// CloseResult 通过done channel 同步等待所有工作goroutine的结束，然后关闭结果chan
func CloseResult(done chan struct{}, resultchan chan int, workers int) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(done)
	close(resultchan)
}

// ProcessResult 读取结果通道，汇总结果
func ProcessResult(resultchan chan int) int {
	sum := 0
	for r := range resultchan {
		sum += r
	}
	return sum
}
func main() {
	workers := Number
	// 工作通道
	taskchan := make(chan Task, 10)
	//结果通道
	resultchan := make(chan int, 10)

	// worker 信号通道
	done := make(chan struct{}, 10)
	// 初始化task的共routine,计算100个自然数之和
	go InitTask(taskchan, resultchan, 1000)
	//分发任务到NUMBER个goroutine池
	DistributeTask(taskchan, workers, done)
	//获取各个goroutine处理完成任务的通知，并关闭结果通道
	go CloseResult(done, resultchan, workers)
	//通过结果通道获取结果并汇总
	sum := ProcessResult(resultchan)
	fmt.Println("sum=", sum)
}
