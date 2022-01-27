package main

import (
	"context"
	protobuf "go_sample/grpc_stream_unary_demo/proto-gens"
	"io"
	"log"
	"google.golang.org/grpc"
)
//模拟client是unary，server是stream
func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cli := protobuf.NewPost06Client(conn)
	facbCli, err := cli.Facb(context.Background(), &protobuf.FacbRequest{Max: int64(100)})
	if err != nil {
		panic(err)
	}

	for {
		if resp, err := facbCli.Recv(); err != nil {
			if err == io.EOF {
				return
			} else {
				panic(err)
			}
		} else {
			log.Printf("[D] index: %d, facb: %d", resp.Index, resp.Curr)
		}
	}
}
