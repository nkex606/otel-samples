package metrics

import (
	"context"
	"log"
	"otel-world/config"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

var WorldMeter metric.Meter

var CommonLabels []attribute.KeyValue = []attribute.KeyValue{
	attribute.String("meter_label_key1", "meter_label_val1"),
	attribute.String("meter_label_key2", "meter_label_val2"),
}

func InitProvider(ctx context.Context, grpcEndpoint string, meterName string) func() {
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(config.GetSvcName()),
		),
	)
	if err != nil {
		log.Fatalf("failed to create resource: %v", err)
	}

	metricExp, err := otlpmetricgrpc.New(
		ctx,
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithEndpoint(grpcEndpoint),
	)
	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
	}

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(res),
		sdkmetric.WithReader(
			sdkmetric.NewPeriodicReader(
				metricExp,
				sdkmetric.WithInterval(2*time.Second),
			),
		),
	)
	otel.SetMeterProvider(meterProvider)

	WorldMeter = meterProvider.Meter(meterName)

	return func() {
		if err := meterProvider.Shutdown(ctx); err != nil {
			otel.Handle(err)
		}
	}
}
