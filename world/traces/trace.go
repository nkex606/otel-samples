package traces

import (
	"context"
	"log"
	"otel-world/config"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"go.opentelemetry.io/otel/trace"
)

var WorldTracer trace.Tracer

var CommonAttrs []attribute.KeyValue = []attribute.KeyValue{
	attribute.String("app", "world"),
	attribute.String("baby", "shark"),
}

func InitProvider(ctx context.Context, grpcEndpoint string, tracerName string) func() {
	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(config.GetSvcName()),
		),
	)
	if err != nil {
		log.Fatalf("failed to create resource: %v", err)
	}

	traceExp, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(grpcEndpoint),
	)
	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
	}

	bsp := sdktrace.NewBatchSpanProcessor(traceExp)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	otel.SetTracerProvider(tracerProvider)

	WorldTracer = tracerProvider.Tracer(tracerName)

	return func() {
		if err := tracerProvider.Shutdown(ctx); err != nil {
			otel.Handle(err)
		}
	}
}
