package api

import (
	"net/http"
	"net/http/pprof"

	"github.com/julienschmidt/httprouter"
	"github.com/kidanekal/goapi/constants"
	"github.com/kidanekal/goapi/logger"
	"github.com/kidanekal/goapi/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/net/context"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "method", "status"},
	)
)

func NewRouter() *httprouter.Router {

	router := httprouter.New()

	//
	// Health
	//
	createRoute(router.GET, constants.HealthPath, HealthHandler, middleware.VersionHeader)

	//
	// Testing
	//
	router.HandlerFunc("POST", "/panic", PanicHandler)

	// Metrics endpoint
	router.Handler("GET", "/metrics", promhttp.Handler())

	addProfiling(router)

	return router

}

func createRoute(method func(path string, handle httprouter.Handle),
	path string,
	handler middleware.Handle,
	middlewares ...func(middleware.Handle) middleware.Handle) {

	for _, m := range middlewares {
		handler = m(handler)
	}

	log := logger.CLI("package", "api")

	// // Wrap the handler with metrics middleware
	routeHandle := middleware.Context(path, log, middleware.PrometheusMiddleware(path, handler))
	method(path, routeHandle)
}

func HealthHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// 200 response is the default
}

func addProfiling(router *httprouter.Router) {
	router.HandlerFunc("GET", "/debug/pprof", pprof.Index)
	router.HandlerFunc("GET", "/debug/pprof/cmdline", pprof.Cmdline)
	router.HandlerFunc("GET", "/debug/pprof/profile", pprof.Profile)
	router.HandlerFunc("GET", "/debug/pprof/symbol", pprof.Symbol)
	router.HandlerFunc("POST", "/debug/pprof/symbol", pprof.Symbol)
	router.Handler("GET", "/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handler("GET", "/debug/pprof/heap", pprof.Handler("heap"))
	router.Handler("GET", "/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	router.Handler("GET", "/debug/pprof/block", pprof.Handler("block"))
}
