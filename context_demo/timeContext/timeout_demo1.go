package main

import (
	"context"
	"fmt"
	"time"
)

// 实现超时的方法之一，在外层使用一个定时器，到点就执行关闭，使用defer进行cancel操作
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
}
