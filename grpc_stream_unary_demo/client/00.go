package main

import (
	"context"
	protobuf "go_sample/grpc_stream_unary_demo/proto-gens"
	"log"

	"google.golang.org/grpc"
)
//模拟client是stream，server是unary
func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cli := protobuf.NewPost06Client(conn)
	sumCli, err := cli.Sum(context.Background())
	if err != nil {
		panic(err)
	}
	sumCli.Send(&protobuf.SumRequest{Num: int64(1)})
	sumCli.Send(&protobuf.SumRequest{Num: int64(2)})
	sumCli.Send(&protobuf.SumRequest{Num: int64(3)})
	if resp, err := sumCli.CloseAndRecv(); err != nil {
		panic(err)
	} else {
		log.Printf("[D] resp: %s", resp.Result)
	}
}
