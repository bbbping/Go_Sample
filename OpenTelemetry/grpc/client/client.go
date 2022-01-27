package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go_sample/OpenTelemetry/grpc/interceptor/trace_interceptor"
	"go_sample/OpenTelemetry/trace_provider"
	"log"

	"go_sample/OpenTelemetry/grpc/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	//"log"
	"time"
)

const (
	ServiceName       = "tracing-grpc-client"
	JaegerUdpEndpoint = "101.34.2.244:6831"
	//JaegerUdpEndpoint = "localhost:6831"
)

func main(){
	//追踪链的东西
	ctx := context.Background()
	//tp, err := trace_provider.InitJaegerLocalFile(ServiceName,"/tmp/grpc_trace.txt")
	tp, err := trace_provider.InitJaeger(ServiceName,JaegerUdpEndpoint)
	if err != nil {
		log.Panic(err)
	}
	defer tp.Shutdown(ctx)

	grpcOptions := []grpc.DialOption{
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                5 * 60 * time.Second,
			PermitWithoutStream: true,
		}),
	}
	//是否跳过加密验证
	rpcWithCredential:=true

	if rpcWithCredential {
		grpcOptions = append(grpcOptions, grpc.WithInsecure()) //grpc.WithInsecure()是暂时跳过了安全传输功能
	} else {
		//https需要用这个方法
		grpcOptions = append(grpcOptions, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})))
	}

	//加载trace拦截器
	grpcOptions= append(grpcOptions,grpc.WithUnaryInterceptor(trace_interceptor.NewGRPUnaryClientInterceptor()))
	//grpcOptions= append(grpcOptions,grpc.WithUnaryInterceptor(my_interceptor.UnaryClientInterceptor()))

	conn, err := grpc.Dial("127.0.0.1:8080", grpcOptions...)
	if err != nil{
		fmt.Println("连接失败:",err)
		return
	}
	defer conn.Close()


	//创建client
	c := user.NewUserClient(conn)

	var span trace.Span
	ctx, span = otel.Tracer(ServiceName).Start(ctx, "client start request")
	defer span.End()
	res, err := c.Query(ctx, &user.QueryReq{Id: 123})
	if err != nil {
		fmt.Println("调用服务失败:",err)
		return
	}
	fmt.Println("调用成功:",res.Name)

}