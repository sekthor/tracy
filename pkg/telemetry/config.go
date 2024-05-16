package telemetry

import "time"

type TelemetryConfig struct {
	Metrics Config
}

type Config struct {
	Enabled  bool
	Protocol string // stdout, http, https, grpc
	Insecure bool
	Interval time.Duration // seconds
	Endpoint string
}
