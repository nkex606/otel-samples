package http

import (
	"context"
	"fmt"
	"net/http"
	"otel-hello/domain"
	"otel-hello/metrics"
	"otel-hello/traces"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type helloHandler struct {
	helloSvc domain.HelloService
}

func NewHelloHandler(g *gin.Engine, hsvc domain.HelloService) {
	handler := &helloHandler{
		helloSvc: hsvc,
	}

	c := g.Group("/helloworld")
	c.GET("/:id", handler.CallWorld)
	c.GET("/capital/:cityname", handler.CallWorldWithCapital)
}

func (h *helloHandler) CallWorld(c *gin.Context) {
	// counter metric
	counter, err := metrics.HelloMeter.Int64Counter(
		"CallWorld.Counter",
		metric.WithDescription("Number of CallWorld calls."),
		metric.WithUnit("{call}"),
	)
	if err != nil {
		fmt.Printf("CallWorld counter err: %v", err)
	}
	counter.Add(c, 1, metric.WithAttributes(metrics.CommonLabels...))

	// histogram metric
	histogram, err := metrics.HelloMeter.Float64Histogram(
		"CallWorld.duration",
		metric.WithDescription("The duration of handler CallWorld execution."),
		metric.WithUnit("s"),
	)
	if err != nil {
		fmt.Printf("CallWorld histogram err: %v", err)
	}

	start := time.Now()
	cityId := c.Param("id")
	city, err := h.helloSvc.CallWorld(c, cityId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Get city fail.")
		return
	}
	duration := time.Since(start)
	histogram.Record(c, duration.Seconds())
	c.JSON(http.StatusOK, city)
}

func (h *helloHandler) CallWorldWithCapital(c *gin.Context) {
	// counter metric
	counter, err := metrics.HelloMeter.Int64Counter(
		"CallWorldWithCapital.Counter",
		metric.WithDescription("Number of CallWorldWithCapital calls."),
		metric.WithUnit("{call}"),
	)
	if err != nil {
		fmt.Printf("CallWorldWithCapital counter err: %v", err)
	}
	counter.Add(c, 1, metric.WithAttributes(metrics.CommonLabels...))

	// trace
	member1, _ := baggage.NewMember("hello", "world")
	member2, _ := baggage.NewMember("asdf", "qwer")
	bag, _ := baggage.New(member1, member2)

	ctxWithBag := baggage.ContextWithBaggage(context.Background(), bag)
	ctx, span := traces.HelloTracer.Start(ctxWithBag, "Request_To_World", trace.WithAttributes(traces.CommonAttrs...))
	defer span.End()
	cityName := c.Param("cityname")
	capital, err := h.helloSvc.CallWorldWithCapital(ctx, cityName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Get capital fail.")
		return
	}

	c.JSON(http.StatusOK, capital)
}
