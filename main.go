package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"golang.org/x/time/rate"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go_kit/endpoints"
	"go_kit/instrument"
	"go_kit/loggings"
	"go_kit/service"
	"go_kit/transports"
)


func main() {

	ctx := context.Background()
	errChan := make(chan error)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var svc service.Service
	svc = service.ArithmeticService{}

	// 第二章加入
	svc = loggings.LoggingMiddleware(logger)(svc)

	endpoint := endpoints.MakeArithmeticEndpoint(svc)

	// 第三章加入
	//ratebucket := ratelimit.NewBucket(time.Second*1, 3)
	//endpoint = instrument.NewTokenBucketLimitterWithJuju(ratebucket)(endpoint)
	// 第三章 第二种方法 加入
	ratebucket := rate.NewLimiter(rate.Every(time.Second*1), 100)
	endpoint = instrument.NewTokenBucketLimitterWithBuildIn(ratebucket)(endpoint)

	// 第四章加入
	fieldKeys := []string{"method"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "raysonxin",
		Subsystem: "arithmetic_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)

	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "raysonxin",
		Subsystem: "arithemetic_service",
		Name:      "request_latency",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	svc = instrument.Metrics(requestCount, requestLatency)(svc)

	r := transports.MakeHttpHandler(ctx, endpoint, logger)

	go func() {
		fmt.Println("Http Server start at port:9000")
		handler := r
		errChan <- http.ListenAndServe(":9000", handler)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println(<-errChan)
}

