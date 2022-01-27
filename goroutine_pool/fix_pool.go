package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)
//使用channel的缓冲队列的容量，限制启动的goroutine数
var wg = sync.WaitGroup{}

func busi(ch chan bool, i int) {
	fmt.Println("go func ", i, " goroutine count = ", runtime.NumGoroutine())
	time.Sleep(2*time.Second)
	<-ch
	wg.Done()
}

func main() { //用户需求的业务数量
	//task cnt := math.MaxInt64
	task_cnt := 10
	ch := make(chan bool, 3)
	for i := 0; i < task_cnt; i++ {
		wg.Add(1)
		ch <- true //如果channel满，就会阻塞
		go busi(ch, i)
	}
	wg.Wait()
}
