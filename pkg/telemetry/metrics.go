package telemetry

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/sdk/metric"
)

func newMeterProvider(ctx context.Context, conf Config) (*metric.MeterProvider, error) {

	exporter, err := newMetricExporter(ctx, conf)

	if err != nil {
		return nil, err
	}

	// set default if unset
	if conf.Interval == 0 {
		conf.Interval = 3
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(exporter,
			metric.WithInterval(conf.Interval*time.Second))),
	)

	return meterProvider, nil
}

func newMetricExporter(ctx context.Context, conf Config) (metric.Exporter, error) {
	var exporter metric.Exporter
	var err error

	switch conf.Type {

	case "http", "https":
		var options []otlpmetrichttp.Option
		if conf.Endpoint != "" {
			options = append(options, otlpmetrichttp.WithEndpoint(conf.Endpoint))
		}
		if conf.Insecure {
			options = append(options, otlpmetrichttp.WithInsecure())
		}
		exporter, err = otlpmetrichttp.New(ctx, options...)

	case "grpc":
		var options []otlpmetricgrpc.Option
		if conf.Endpoint != "" {
			options = append(options, otlpmetricgrpc.WithEndpoint(conf.Endpoint))
		}
		if conf.Insecure {
			options = append(options, otlpmetricgrpc.WithInsecure())
		}
		exporter, err = otlpmetricgrpc.New(ctx, options...)

	default:
		exporter, err = stdoutmetric.New()
	}

	return exporter, err
}
