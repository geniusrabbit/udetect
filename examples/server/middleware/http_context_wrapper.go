package middleware

import (
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/sspserver/udetect/examples/server/context/ctxlogger"
)

// HTTPContextWrapper global middelware of the HTTP rounter
func HTTPContextWrapper(h http.Handler, ctxWrapper func(ctx context.Context) context.Context) http.Handler {
	var (
		buckets      = []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10}
		metricsCount = promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "http_request_count",
			Help: "Count of HTTP requests",
		}, []string{"method", "path"})
		metricTiming = prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "server",
			Name:      "http_request_duration_seconds",
			Help:      "Histogram of response time for handler in seconds",
			Buckets:   buckets,
		}, []string{"method", "path"})
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		defer func() {
			duration := time.Since(startTime)
			metricsCount.WithLabelValues(r.Method, r.URL.Path).Inc()
			metricTiming.WithLabelValues(r.Method, r.URL.Path).Observe(duration.Seconds())
		}()

		if ctxWrapper != nil {
			newCtx := ctxWrapper(r.Context())
			r = r.WithContext(newCtx)
		}
		ctxlogger.Get(r.Context()).Info("HTTP",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path))
		h.ServeHTTP(w, r)
	})
}
