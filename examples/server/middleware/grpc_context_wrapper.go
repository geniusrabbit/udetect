package middleware

import (
	"context"
	"sync"
	"time"

	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type tagsType struct {
	s []string
}

var (
	metricsCount *prometheus.CounterVec
	metricTiming *prometheus.HistogramVec
	tagsBuffer   = sync.Pool{New: func() interface{} {
		return &tagsType{s: make([]string, 0, 2)}
	}}
)

func init() {
	buckets := []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10}
	// NOTE: Register once to avoid double registration
	metricsCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "grpc_request_count",
		Help: "Count of GRPC requests",
	}, []string{"method", "error"})
	metricTiming = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "server",
		Name:      "grpc_request_duration_seconds",
		Help:      "Histogram of response time for handler in seconds",
		Buckets:   buckets,
	}, []string{"method", "error"})
}

// GRPCContextUnaryWrapper implements wrapper of unary handler with context overriding
func GRPCContextUnaryWrapper(ctxWrapper func(ctx context.Context) (context.Context, error)) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		var (
			startTime = time.Now()
			newCtx    context.Context
			tags      = tagsBuffer.Get().(*tagsType)
		)
		tags.s = append(tags.s[:0], info.FullMethod, "0")
		defer func() {
			if err != nil {
				tags.s[1] = "1"
			}
			duration := time.Since(startTime)
			metricsCount.WithLabelValues(tags.s...).Inc()
			metricTiming.WithLabelValues(tags.s...).Observe(duration.Seconds())
			tagsBuffer.Put(tags)
		}()
		if newCtx, err = ctxWrapper(ctx); err != nil {
			return nil, err
		}
		return handler(newCtx, req)
	}
}

// GRPCContextStreamWrapper implements wrapper of stream handler with context overriding
func GRPCContextStreamWrapper(ctxWrapper func(ctx context.Context) (context.Context, error)) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		var (
			startTime = time.Now()
			tags      = tagsBuffer.Get().(*tagsType)
		)
		tags.s = append(tags.s[:0], info.FullMethod, "0")
		defer func() {
			if err != nil {
				tags.s[1] = "1"
			}
			duration := time.Since(startTime)
			metricsCount.WithLabelValues(tags.s...).Inc()
			metricTiming.WithLabelValues(tags.s...).Observe(duration.Seconds())
			tagsBuffer.Put(tags)
		}()
		newStream := grpc_middleware.WrapServerStream(ss)
		newStream.WrappedContext, err = ctxWrapper(ss.Context())
		if err != nil {
			return err
		}
		return handler(srv, ss)
	}
}
