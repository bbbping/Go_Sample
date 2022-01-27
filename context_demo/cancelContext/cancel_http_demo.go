package main

import (
	"context"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)
//一个可中止、可控制的http请求
func httpRequest(
	ctx context.Context,
	client *http.Client,
	req *http.Request,
	respChan chan []byte,
	errChan chan error,
) {
	//关键这里，使用request创建一个上下文，以便可以关闭
	req = req.WithContext(ctx)
	tr := &http.Transport{}
	client.Transport = tr
	go func() {
		resp, err := client.Do(req)
		if err != nil {
			errChan <- err
		}
		if resp != nil {
			defer resp.Body.Close()
			respData, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				errChan <- err
			}
			respChan <- respData
		} else {
			errChan <- errors.New("HTTP request failed")
		}
	}()
	for {
		select {
		case <-ctx.Done():
			tr.CancelRequest(req)
			errChan <- errors.New("HTTP request cancelled")
			return
		case <-errChan:
			tr.CancelRequest(req)
			return
		}
	}
}
