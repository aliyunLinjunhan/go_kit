package instrument

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics"
	"github.com/juju/ratelimit"
	"golang.org/x/time/rate"
	"time"

	"go_kit/service"
)

var ErrLimitExceed = errors.New("Rate limit exceed!")

func NewTokenBucketLimitterWithJuju(bkt *ratelimit.Bucket) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if bkt.TakeAvailable(1) == 0 {
				return nil, ErrLimitExceed
			}
			return next(ctx, request)
		}
	}
}

// NewTokenBucketLimitterWithBuildIn 使用x/time/rate创建限流中间件
func NewTokenBucketLimitterWithBuildIn(bkt *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !bkt.Allow() {
				return nil, ErrLimitExceed
			}
			return next(ctx, request)
		}
	}
}

// metricMiddleware 定义监控中间件，嵌入Service
// 新增监控指标项：requestCount和requestLatency
type metricMiddleware struct {
	service.Service
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
}

// Metrics 指标采集方法
func Metrics(requestCount metrics.Counter, requestLatency metrics.Histogram) service.ServiceMiddleware {
	return func(next service.Service) service.Service {
		return metricMiddleware{
			next,
			requestCount,
			requestLatency}
	}
}

func (mw metricMiddleware) Add(a, b int) (ret int) {

	defer func(beign time.Time) {
		lvs := []string{"method", "Add"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(beign).Seconds())
	}(time.Now())

	ret = mw.Service.Add(a, b)
	return ret
}

func (mw metricMiddleware) Subtract(a, b int) (ret int) {

	defer func(beign time.Time) {
		lvs := []string{"method", "Subtract"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(beign).Seconds())
	}(time.Now())

	ret = mw.Service.Subtract(a, b)
	return ret
}

func (mw metricMiddleware) Multiply(a, b int) (ret int) {

	defer func(beign time.Time) {
		lvs := []string{"method", "Multiply"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(beign).Seconds())
	}(time.Now())

	ret = mw.Service.Multiply(a, b)
	return ret
}

func (mw metricMiddleware) Divide(a, b int) (int, error) {

	defer func(beign time.Time) {
		lvs := []string{"method", "Divide"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(beign).Seconds())
	}(time.Now())

	ret, err := mw.Service.Divide(a, b)
	return ret, err
}
