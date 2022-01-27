package my_interceptor

import (
	"context"
	"fmt"
	"go_sample/OpenTelemetry/grpc/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		requestMetadata, _ := metadata.FromIncomingContext(ctx)
		metadataCopy := requestMetadata.Copy()
		fmt.Println("发送请求的元数据：")
		for k, v := range metadataCopy {
			fmt.Println("key:",k,"value:",v)
		}
		fmt.Println("==============")


		//tracer := otel.GetTracerProvider().Tracer(
		//	tracingInstrumentGrpcServer,
		//	trace.WithInstrumentationVersion("v0.12"),
		//)
		//
		////从元数据中获取client传送过来的bags
		//bags, spanCtx := Extract(ctx, metadataCopy)
		//ctx = baggage.ContextWithBaggage(ctx, bags)
		//ctx = trace.ContextWithRemoteSpanContext(ctx, spanCtx)

		//bags, spanCtx := Extract(ctx, &metadataCopy, opts...)
		//ctx = baggage.ContextWithBaggage(ctx, bags)
		//
		//tracer := newConfig(opts).TracerProvider.Tracer(
		//	instrumentationName,
		//	trace.WithInstrumentationVersion(SemVersion()),
		//)

		//name, attr := spanInfo(info.FullMethod, peerFromCtx(ctx))
		//ctx, span := tracer.Start(
		//	trace.ContextWithRemoteSpanContext(ctx, spanCtx),
		//	name,
		//	trace.WithSpanKind(trace.SpanKindServer),
		//	trace.WithAttributes(attr...),
		//)
		//defer span.End()

		//messageReceived.Event(ctx, 1, req)
		fmt.Println("开始调用rpc方法")
		resp, err := handler(ctx, req)
		fmt.Println("调用rpc方法结束，resp",resp)

		//fmt.Println(reflect.TypeOf(resp))
		res,ok := resp.(*user.QueryRes)
		if ok{
			fmt.Println("在server拦截器中修改返回参数")
			res.Name="小明"
			resp = res
		}
		//if err != nil {
		//	s, _ := status.FromError(err)
		//	span.SetStatus(codes.Error, s.Message())
		//	span.SetAttributes(statusCodeAttr(s.Code()))
		//	messageSent.Event(ctx, 1, s.Proto())
		//} else {
		//	span.SetAttributes(statusCodeAttr(grpc_codes.OK))
		//	messageSent.Event(ctx, 1, resp)
		//}

		return resp, err
	}
}

