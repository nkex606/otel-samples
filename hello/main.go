package main

import (
	"context"
	"fmt"
	"net/http"
	"otel-hello/config"
	"otel-hello/metrics"
	"otel-hello/traces"
	"time"

	rest "otel-hello/delivery/http"
	_microSvc "otel-hello/microsvc"
	_helloSvc "otel-hello/service"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	// init app config
	config.InitConf()

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
	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}

	// microservice layer
	wmc := _microSvc.NewWorldMicroSvc(client)

	// service layer
	hsvc := _helloSvc.NewHelloService(wmc)

	// handler
	rest.NewHealthzHandler(g)
	rest.NewHelloHandler(g, hsvc)

	// start server
	g.Run(config.GetHttpConfig().Port)
}
