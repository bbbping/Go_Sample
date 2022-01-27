package main

import (
	"context"
	"fmt"
	"sync"
	"github.com/pkg/errors"
	"time"
)

//一个使用context进行取消的例子
func Rpc(ctx context.Context, url string,isSuccess bool,second int64) error {
	result := make(chan int)
	err := make(chan error)

	go func() {
		// 进行RPC调用，并且返回是否成功，成功通过result传递成功信息，错误通过error传递错误信息
		//isSuccess := true
		time.Sleep(time.Duration(second)*time.Second)
		if isSuccess {
			result <- 1
		} else {
			err <- errors.New("some error happen")
		}
	}()

	select {
	case <- ctx.Done():
		// 其他RPC调用调用失败
		fmt.Println(url,"被取消了")
		return ctx.Err()
	case e := <- err:
		// 本RPC调用失败，返回错误信息
		fmt.Println(url,"执行失败")
		return e
	case <- result:
		// 本RPC调用成功，不返回错误信息
		fmt.Println(url,"执行成功")
		return nil
	}
}


func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// RPC1调用
	err := Rpc(ctx, "http://rpc_1_url",true,1)
	if err != nil {
		return
	}

	wg := sync.WaitGroup{}

	// RPC2调用
	wg.Add(1)
	go func(){
		defer wg.Done()
		err := Rpc(ctx, "http://rpc_2_url",false,1)
		if err != nil {
			cancel()
		}
	}()

	// RPC3调用
	wg.Add(1)
	go func(){
		defer wg.Done()
		err := Rpc(ctx, "http://rpc_3_url",true,5)
		if err != nil {
			cancel()
		}
	}()

	// RPC4调用
	wg.Add(1)
	go func(){
		defer wg.Done()
		err := Rpc(ctx, "http://rpc_4_url",true,5)
		if err != nil {
			cancel()
		}
	}()

	wg.Wait()
}
