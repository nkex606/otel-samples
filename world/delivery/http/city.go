package rest

import (
	"fmt"
	"net/http"
	"otel-world/metrics"
	"otel-world/model"
	"otel-world/traces"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type cityHandler struct {
	cityservice model.CityService
}

func NewCityHandler(g *gin.Engine, cs model.CityService) {
	handler := cityHandler{
		cityservice: cs,
	}

	g.GET("/city/:id", handler.GetCityName)
	g.GET("/city/capital/:name", handler.GetCapitalByCity)
}

func (cityHandler *cityHandler) GetCityName(c *gin.Context) {
	// counter metric
	counter, err := metrics.WorldMeter.Int64Counter(
		"GetCityName.Counter",
		metric.WithDescription("Number of GetCityName calls."),
		metric.WithUnit("{call}"),
	)
	if err != nil {
		fmt.Printf("GetCityName counter err: %v", err)
	}
	counter.Add(c, 1, metric.WithAttributes(metrics.CommonLabels...))

	cid := c.Param("id")
	id, _ := strconv.Atoi(cid)
	city, err := cityHandler.cityservice.CityNameById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Get city fail: %s", err.Error()))
	}

	c.JSON(http.StatusOK, gin.H{"name": city})
}

func (cityHandler *cityHandler) GetCapitalByCity(c *gin.Context) {
	// counter metric
	counter, err := metrics.WorldMeter.Int64Counter(
		"GetCapitalByCity.Counter",
		metric.WithDescription("Number of GetCapitalByCity calls."),
		metric.WithUnit("{call}"),
	)
	if err != nil {
		fmt.Printf("GetCapitalByCity counter err: %v", err)
	}
	counter.Add(c, 1, metric.WithAttributes(metrics.CommonLabels...))

	// trace
	span := trace.SpanFromContext(c.Request.Context())
	bag := baggage.FromContext(c.Request.Context())

	var baggageAttributes []attribute.KeyValue
	baggageAttributes = append(baggageAttributes, traces.CommonAttrs...)
	for _, member := range bag.Members() {
		baggageAttributes = append(baggageAttributes, attribute.String("baggage key:"+member.Key(), member.Value()))
	}
	span.SetAttributes(baggageAttributes...)

	cityName := c.Param("name")
	capital, err := cityHandler.cityservice.CapitalNameByCity(c, cityName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Get capital fail: %s", err.Error()))
	}
	c.JSON(http.StatusOK, gin.H{"name": capital})
}
