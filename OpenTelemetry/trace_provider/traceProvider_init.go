package trace_provider

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"io"
	"log"
	"os"
	"strings"
)

func InitJaeger(serviceName, endpoint string) (tp *trace.TracerProvider, err error) {
	var endpointOption jaeger.EndpointOption
	if strings.HasPrefix(endpoint, "http") {
		// HTTP.
		endpointOption = jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint))
	} else {
		// UDP.
		agentConfig := strings.SplitN(endpoint, ":", 2)
		if len(agentConfig) == 2 {
			endpointOption = jaeger.WithAgentEndpoint(jaeger.WithAgentHost(agentConfig[0]), jaeger.WithAgentPort(agentConfig[1]))
		} else {
			endpointOption = jaeger.WithAgentEndpoint(jaeger.WithAgentHost(agentConfig[0]))
		}
	}

	// Create the Jaeger exporter
	exp, err := jaeger.New(endpointOption)
	if err != nil {
		return nil, err
	}
	tp = trace.NewTracerProvider(
		// Always be sure to batch in production.
		trace.WithBatcher(exp),
		// Record information about this application in an Resource.
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		)),
	)
	otel.SetTracerProvider(tp)
	return tp, nil
}

func InitJaegerLocalFile(serviceName, filePath string) (tp *trace.TracerProvider, err error) {

	l := log.New(os.Stdout, "", 0)

	// Write telemetry data to a file.
	f, err := os.Create(filePath)
	if err != nil {
		l.Fatal(err)
	}
	//fixme 实验阶段，不关闭日志文件了
	//defer f.Close()

	//设置日志输出的流
	exp, err := newExporter(f)
	if err != nil {
		l.Fatal(err)
	}

	//创建追踪提供者，将输出的流，和原始元素设置好
	tp = trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(newResource(serviceName)),
	)

	//设置一下输出
	otel.SetTracerProvider(tp)

	return tp, nil
}

// newResource returns a resource describing this application.
func newResource(serviceName string) *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
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
