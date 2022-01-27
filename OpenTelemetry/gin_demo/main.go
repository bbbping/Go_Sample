package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	strace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/attribute"
)

func InitRouter() (*gin.Engine, error) {
	r := gin.Default()
	gin.SetMode("debug")
	ro := r.Group("/trace").Use(otelgin.Middleware("mike_gin_trace_demo"))
	ro.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "pong",
		})
	})

	ro.GET("/age", func(c *gin.Context) {
		age := c.Query("age")

		c.JSON(200, gin.H{
			"msg": "your age="+age,
		})
	})

	return r, nil
}

func main() {

	l := log.New(os.Stdout, "", 0)
	// Write telemetry data to a file.
	f, err := os.Create("/tmp/gin_traces.txt")
	if err != nil {
		l.Fatal(err)
	}
	defer f.Close()

	//设置日志输出的流
	exp, err := newExporter(f)
	if err != nil {
		l.Fatal(err)
	}

	//创建追踪提供者，将输出的流，和原始元素设置好
	tp := strace.NewTracerProvider(
		strace.WithBatcher(exp),
		strace.WithResource(newResource()),
	)

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			l.Fatal(err)
		}
	}()
	//设置一下输出
	otel.SetTracerProvider(tp)

	router, _ := InitRouter()
	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    time.Second * 60,
		WriteTimeout:   time.Second * 60,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}
	if err := server.ListenAndServe(); err != nil {
		panic("启动服务失败")
	}
}

// newResource returns a resource describing this application.
func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("服务名称：mike-gin-demo"),
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
