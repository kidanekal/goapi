package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

func TestPrometheusMiddleware(t *testing.T) {
	// Reset the Prometheus metric
	httpRequestsTotal.Reset()

	// Create a new Prometheus registry and register the httpRequestsTotal metric
	registry := prometheus.NewRegistry()
	registry.MustRegister(httpRequestsTotal)

	// Define the next handler
	next := func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot) // Use a unique status code for testing
	}

	// Create a new request and response recorder
	req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	// Wrap the next handler with the middleware
	handler := prometheusMiddleware("/foo", next)

	// Call the handler
	handler(context.Background(), w, req)

	// Collect the metrics
	metricFamilies, err := registry.Gather()
	assert.NoError(t, err)

	// Find the http_requests_total metric
	var found bool
	for _, mf := range metricFamilies {
		if mf.GetName() == "http_requests_total" {
			for _, m := range mf.GetMetric() {
				var path, method, status string
				for _, label := range m.GetLabel() {
					if label.GetName() == "path" {
						path = label.GetValue()
					} else if label.GetName() == "method" {
						method = label.GetValue()
					} else if label.GetName() == "status" {
						status = label.GetValue()
					}
				}
				if path == "/foo" && method == http.MethodGet && status == http.StatusText(http.StatusTeapot) {
					assert.Equal(t, float64(1), m.GetCounter().GetValue())
					found = true
				}
			}
		}
	}

	assert.True(t, found, "Expected metric not found")
}
