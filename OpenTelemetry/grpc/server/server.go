package main

import (
	"context"
	"fmt"
	"go_sample/OpenTelemetry/grpc/interceptor/trace_interceptor"
	"go_sample/OpenTelemetry/grpc/user"
	"go_sample/OpenTelemetry/trace_provider"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

const (
	ServiceName       = "tracing-grpc-server"
	//JaegerUdpEndpoint = "localhost:6831"
	JaegerUdpEndpoint = "101.34.2.244:6831"
)

type server struct{}

func (s *server) Insert(ctx context.Context, in *user.InsertReq) (*user.InsertRes, error) {
	ts := time.Now().Unix()
	//name := in.Name
	return &user.InsertRes{Id: int32(ts)}, nil
}

func (s *server) Query(ctx context.Context, in *user.QueryReq) (*user.QueryRes, error) {
	return &user.QueryRes{Id: 123, Name: "小米"}, nil
}
func (s *server) Delete(ctx context.Context, in *user.DeleteReq) (*user.DeleteRes, error) {
	return &user.DeleteRes{}, nil
}

func main() {
	ctx := context.Background()
	tp, err := trace_provider.InitJaeger(ServiceName,JaegerUdpEndpoint)
	if err != nil {
		log.Panic(err)
	}
	defer tp.Shutdown(ctx)

	if err != nil {
		log.Fatal("初始化jaeger失败")
	}

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("监听端口失败:", err)
		return
	}

	//创建grpc
	s := grpc.NewServer(
		//加载trace拦截器
		//grpc.UnaryInterceptor(my_interceptor.UnaryServerInterceptor()),
		grpc.UnaryInterceptor(trace_interceptor.NewGRPUnaryServerInterceptor()),
	)
	//注册服务
	user.RegisterUserServer(s, &server{})
	reflection.Register(s)
	err = s.Serve(lis)
	if err != nil {
		fmt.Println("开启服务失败:", err)
		return
	}

}
