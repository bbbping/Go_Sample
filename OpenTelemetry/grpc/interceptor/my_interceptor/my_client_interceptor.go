package my_interceptor

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go_sample/OpenTelemetry/grpc/grpcutils"
	"google.golang.org/grpc"
	grpcCodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context, //上下文，元数据都在这里
		method string, //当前rpc方法名称
		req, reply interface{}, //请求/返回对象
		cc *grpc.ClientConn, //连接
		invoker grpc.UnaryInvoker, //调用器
		callOpts ...grpc.CallOption,  //调用option
	) error {
		//读取元数据
		requestMetadata, _ := metadata.FromOutgoingContext(ctx)
		//常规操作，复制一份
		metadataCopy := requestMetadata.Copy()

		tracer := otel.GetTracerProvider().Tracer(
			tracingInstrumentGrpcClient,
			trace.WithInstrumentationVersion("v0.1"),
		)

		//根据调用的rpc方法名称和目标，获取span的一些参数
		name, attr := spanInfo(method, cc.Target())
		//创建span
		var span trace.Span
		ctx, span = tracer.Start(
			ctx,
			name,
			trace.WithSpanKind(trace.SpanKindClient),
			trace.WithAttributes(attr...),
		)
		defer span.End()
		//将span关联的ctx注入到元数据中
		fmt.Println("metadataCopy 注入前:",metadataCopy)
		Inject(ctx, metadataCopy)
		fmt.Println("metadataCopy 注入后:",metadataCopy)
		//fmt.Println("发送请求的元数据：")
		//for k, v := range metadataCopy {
		//	fmt.Println("key:",k,"value:",v)
		//}
		//fmt.Println("==============")

		//然后再根据元数据和ctx再衍生出新的xtx
		fmt.Println("衍生前的ctx",ctx)
		ctx = metadata.NewOutgoingContext(ctx, metadataCopy)
		fmt.Println("衍生后的ctx",ctx)

		fmt.Println("发送请求的元数据：")
		for k, v := range metadataCopy {
			fmt.Println("key:",k,"value:",v)
		}
		fmt.Println("==============")

		//messageSent.Event(ctx, 1, req)

		span.AddEvent(tracingEventGrpcRequest, trace.WithAttributes(
			//attribute.String(tracingEventGrpcRequestBaggage, gconv.String(gtrace.GetBaggageMap(ctx))),
			//attribute.String(tracingEventGrpcMetadataOutgoing, gconv.String(grpcctx.Ctx.OutgoingMap(ctx))),
			attribute.String(
				tracingEventGrpcRequestMessage,
				grpcutils.MarshalMessageToJsonStringForTracing(
					req, "Request", tracingMaxContentLogSize,
				),
			),
		))

		//调用远程方法
		err := invoker(ctx, method, req, reply, cc, callOpts...)


		//记录返回的数据到span中
		span.AddEvent(tracingEventGrpcResponse, trace.WithAttributes(
			attribute.String(
				tracingEventGrpcResponseMessage,
				grpcutils.MarshalMessageToJsonStringForTracing(
					reply, "Response", tracingMaxContentLogSize,
				),
			),
		))

		//用追踪链记录错误
		if err != nil {
			s, _ := status.FromError(err)
			span.SetStatus(codes.Error, s.Message())
			span.SetAttributes(statusCodeAttr(s.Code()))
		} else {
			span.SetAttributes(statusCodeAttr(grpcCodes.OK))
		}

		return err
	}
}

