package telemetry

import "time"

type Config struct {
	Enabled  bool
	Type     string // stdout, http, https, grpc
	Insecure bool
	Interval time.Duration // seconds
	Endpoint string
}
