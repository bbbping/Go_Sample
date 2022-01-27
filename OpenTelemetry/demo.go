package main

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
	"go_sample/OpenTelemetry/trace_provider"
	"io"
	"log"
	"os"
	"os/signal"
	"strconv"
)

// name is the Tracer name used to identify this instrumentation library.
const name = "fib_mike_demo"


//OpenTelemetry例子

func main() {
	//l := log.New(os.Stdout, "", 0)
	//
	//// Write telemetry data to a file.
	//f, err := os.Create("/tmp/traces.txt")
	//if err != nil {
	//	l.Fatal(err)
	//}
	//defer f.Close()
	//
	////设置日志输出的流
	//exp, err := newExporter(f)
	//if err != nil {
	//	l.Fatal(err)
	//}
	//
	////创建追踪提供者，将输出的流，和原始元素设置好
	//tp := strace.NewTracerProvider(
	//	strace.WithBatcher(exp),
	//	strace.WithResource(newResource()),
	//)
	tp, err := trace_provider.InitJaegerLocalFile(name, "/tmp/traces.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()
	//设置一下输出
	otel.SetTracerProvider(tp)


	//下面的是业务逻辑
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	errCh := make(chan error)
	l := log.New(os.Stdout, "", 0)
	app := NewApp(os.Stdin, l)
	go func() {
		errCh <- app.Run(context.Background())
	}()

	select {
	case <-sigCh:
		log.Println("\ngoodbye")
		return
	case err := <-errCh:
		if err != nil {
			log.Fatal(err)
		}
	}
}

// newResource returns a resource describing this application.
func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("服务名称：fib"),
			semconv.ServiceVersionKey.String("版本：v0.1.0"),
			attribute.String("environment", "环境：demo"),
		),
	)
	return r
}

// newExporter returns a console exporter.
func newExporter(w io.Writer) (*stdouttrace.Exporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human-readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
}

//一个计算斐波拉契数列的样例函数
func Fibonacci(n uint) (uint64, error) {
	if n <= 1 {
		return uint64(n), nil
	}

	if n > 93 {
		return 0, fmt.Errorf("unsupported fibonacci number %d: too large", n)
	}


	var n2, n1 uint64 = 0, 1
	for i := uint(2); i < n; i++ {
		n2, n1 = n1, n1+n2
	}

	return n2 + n1, nil
}

// App is a Fibonacci computation application.
type App struct {
	r io.Reader
	l *log.Logger
}

// NewApp returns a new App.
func NewApp(r io.Reader, l *log.Logger) *App {
	return &App{r: r, l: l}
}

// 运行主模块Run starts polling users for Fibonacci number requests and writes results.
func (a *App) Run(ctx context.Context) error {
	for {
		var span trace.Span
		ctx, span = otel.Tracer(name).Start(ctx, "Run func")
		n, err := a.Poll(ctx)
		if err != nil {
			span.End()
			return err
		}

		a.Write(ctx, n)
		span.End()
	}
}

// 监听用户输入的方法Poll asks a user for input and returns the request.
func (a *App) Poll(ctx context.Context) (uint, error) {
	_, span := otel.Tracer(name).Start(ctx, "Poll func")
	defer span.End()

	a.l.Print("请输入你想进行计算的斐波拉契数列: ")

	var n uint
	_, err := fmt.Fscanf(a.r, "%d\n", &n)

	//用户输入负数也报错
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return 0, err
	}

	// Store n as a string to not overflow an int64.
	nStr := strconv.FormatUint(uint64(n), 10)
	span.SetAttributes(attribute.String("请求参数n的值：", nStr))
	return n, err
}

// 输出到用户Write writes the n-th Fibonacci number back to the user.
func (a *App) Write(ctx context.Context, n uint) {
	var span trace.Span
	ctx, span = otel.Tracer(name).Start(ctx, "Write func")
	defer span.End()

	//f, err := Fibonacci(n)

	//下面使用了一个方法，在父方法对子方法的内容进行追踪
	f, err := func(ctx context.Context) (uint64, error) {
		_, span := otel.Tracer(name).Start(ctx, "Fibonacci")
		defer span.End()
		f, err := Fibonacci(n)
		if err != nil {
			span.RecordError(err) //使用sdk的记录错误的方法
			span.SetStatus(codes.Error, err.Error())
		}
		return f, err
	}(ctx)

	if err != nil {
		a.l.Printf("Fibonacci(%d)的结果: %v\n", n, err)
	} else {
		a.l.Printf("Fibonacci(%d)的结果 = %d\n", n, f)
	}
}
