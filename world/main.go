package main

import (
	"context"
	"fmt"
	"otel-world/config"
	"otel-world/database"
	rest "otel-world/delivery/http"
	"otel-world/metrics"
	"otel-world/traces"
	"time"

	_cityRepo "otel-world/repo"
	_citySvc "otel-world/service"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {
	// init app config
	config.InitConf()
	database.InitMysql()

	pctx := context.Background()
	ctx, cancel := context.WithTimeout(pctx, time.Second)
	defer cancel()

	otelCollectorEp := config.GetOTELConfig().Endpoint

	// init tracer provider and tracer
	shutdownTP := traces.InitProvider(ctx, otelCollectorEp, fmt.Sprintf("%s-tracer", config.GetSvcName()))
	defer shutdownTP()

	// init meter provider and meter
	shutdownMP := metrics.InitProvider(ctx, otelCollectorEp, fmt.Sprintf("%s-meter", config.GetSvcName()))
	defer shutdownMP()

	g := gin.Default()
	g.Use(otelgin.Middleware(config.GetSvcName()))
	// repo layer
	cityRepo := _cityRepo.NewCityRepo(database.GetDB())

	// service layer
	citySvc := _citySvc.NewCityService(cityRepo)

	// handler
	rest.NewCityHandler(g, citySvc)
	rest.NewHealthzHandler(g)

	// start server
	g.Run(config.GetHttpConfig().Port)
}
