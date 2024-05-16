package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sekthor/tracy/pkg/telemetry"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

var (
	meter    = otel.Meter("service")
	reqCount metric.Int64Counter
)

func init() {
	var err error
	reqCount, err = meter.Int64Counter("request.total",
		metric.WithDescription("total amount of requests"),
		metric.WithUnit("{request}"),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ctx := context.Background()

	conf := telemetry.TelemetryConfig{
		Metrics: telemetry.Config{
			Enabled:  true,
			Type:     "http",
			Insecure: true,
			Endpoint: "localhost:4318",
		},
	}

	shutdown, err := telemetry.SetupOTelSDK(ctx, conf)

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.GET("greet", greet)
	err = router.Run()
	if err != nil {
		log.Fatal(err)
	}

	defer shutdown(ctx)
}

func greet(c *gin.Context) {
	reqCount.Add(c, 1)
	c.JSON(http.StatusOK, gin.H{
		"msg": "Hello World",
	})
}
