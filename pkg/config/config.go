package config

import (
	"os"

	"github.com/sekthor/tracy/pkg/telemetry"
)

type Config struct {
	Telemetry telemetry.TelemetryConfig
}

func ReadEnv() Config {
	var conf Config
	var option string

	if option = os.Getenv("OTEL_ENDPOINT"); option != "" {
		conf.Telemetry.Metrics.Endpoint = option
	}

	if option = os.Getenv("OTEL_INSECURE"); option == "true" {
		conf.Telemetry.Metrics.Insecure = true
	}

	if option = os.Getenv("OTEL_PROTOCOL"); option != "" {
		conf.Telemetry.Metrics.Protocol = option
	}

	if option = os.Getenv("OTEL_METRICS_ENABLED"); option == "true" {
		conf.Telemetry.Metrics.Enabled = true
	}

	return conf
}
