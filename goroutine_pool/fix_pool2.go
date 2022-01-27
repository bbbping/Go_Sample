package main

import (
	"fmt"
	"runtime"
	"sync"
)

//使用channel的缓冲队列的容量，限制启动的goroutine数
var wg2 = sync.WaitGroup{}

func busi2(ch chan int, i int) {

	for task := range ch {
		fmt.Println("go task ", task, " goroutine count = ", runtime.NumGoroutine())
		wg2.Done()
	}
}

//有goroutine取出任务才能塞进去
func sendTask(task int, ch chan int) {
	wg2.Add(1)
	ch <- task
}

func main() {

	ch := make(chan int) //无buffer的

	//启动工作池(go的数量固定)
	total_go := 3
	for i := 0; i < total_go; i++ {
		go busi2(ch, i)
	}

	task_total := 100

	for t := 0; t < task_total; t++ {
		//发任务
		sendTask(t, ch)
	}
	wg2.Wait()
}
