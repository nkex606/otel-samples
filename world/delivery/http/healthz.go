package rest

import (
	"fmt"
	"net/http"

	"otel-world/metrics"
	"otel-world/traces"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type healthzHandler struct {
}

func NewHealthzHandler(g *gin.Engine) {
	handler := &healthzHandler{}
	g.GET("/healthz", handler.Check)
}

func (h *healthzHandler) Check(c *gin.Context) {
	// counter metric
	counter, err := metrics.WorldMeter.Int64Counter(
		"healthz.Counter",
		metric.WithDescription("Number of healthz calls."),
		metric.WithUnit("{call}"),
	)
	if err != nil {
		fmt.Printf("healthz counter err: %v", err)
	}
	counter.Add(c, 1, metric.WithAttributes(metrics.CommonLabels...))

	// trace
	_, span := traces.WorldTracer.Start(c, "healthz request")
	span.SetAttributes(traces.CommonAttrs...)
	span.SetAttributes(attribute.String("handler", "healthzHandler"))
	span.SetAttributes(attribute.String("handler func", "Check"))
	span.AddEvent("Check event", trace.WithAttributes(attribute.Int("intVal", 1234), attribute.String("stringVal", "xoxo")))
	span.End()

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
