package middleware

import (
	"context"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

// responseWriter wraps the standard http.ResponseWriter to capture the status code.
type responseWriter struct {
	http.ResponseWriter
	status int
}

// WriteHeader captures the status code for the response.
func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.status = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

// PrometheusMiddleware is an HTTP middleware that collects Prometheus metrics.
func PrometheusMiddleware(path string, next Handle) Handle {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}
		next(ctx, rw, r)
		status := rw.status
		httpRequestsTotal.WithLabelValues(path, r.Method, http.StatusText(status)).Inc()
	}
}

// httpRequestsTotal is a Prometheus counter vector to track HTTP requests.
var httpRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	},
	[]string{"path", "method", "status"},
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
}
