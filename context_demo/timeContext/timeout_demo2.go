package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

// 实现超时的方法之二，context使用timeout设置超时时间
func main() {
	uri := "https://httpbin.org/delay/3"
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Fatalf("http.NewRequest() failed with '%s'\n", err)
	}
	//这里设置了超时时间后100ms就返回了
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*100)
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("http.DefaultClient.Do() failed with:\n'%s'\n", err)
	}
	defer resp.Body.Close()
}
