package handlers

import (
	"fmt"
	"github.com/autom8ter/util"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"google.golang.org/grpc/grpclog"
	"net/http"
	"net/http/pprof"
	"os"
)

type RouterMiddleware func(r *mux.Router) *mux.Router

func DebugFunc() RouterMiddleware {
	return func(r *mux.Router) *mux.Router {
		r.HandleFunc("/debug", func(w http.ResponseWriter, request *http.Request) {
			fmt.Println("registered handler: ", "/debug")
			w.Write([]byte(fmt.Sprintln("Status: 💡 API is up and running 💡 ")))
			w.Write([]byte(fmt.Sprintln("---------------------------------------------------------------------")))
			w.Write([]byte(fmt.Sprintln("Configuration Settings:")))
			w.Write([]byte(fmt.Sprintln(util.ToPrettyJsonString(viper.AllSettings()))))
			w.Write([]byte(fmt.Sprintln("---------------------------------------------------------------------")))
			w.Write([]byte(fmt.Sprintln("Environmental Variables:")))
			w.Write([]byte(fmt.Sprintln(os.Environ())))
			w.Write([]byte(fmt.Sprintln("---------------------------------------------------------------------")))
			w.Write([]byte(fmt.Sprintln("Registered Router Paths:")))
			err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
				type routeLog struct {
					Name         string
					PathRegExp   string
					PathTemplate string
					HostTemplate string
					Methods      []string
				}
				meth, _ := route.GetMethods()
				host, _ := route.GetHostTemplate()
				pathreg, _ := route.GetPathRegexp()
				pathtemp, _ := route.GetPathTemplate()
				rout := &routeLog{
					Name:         route.GetName(),
					PathRegExp:   pathreg,
					PathTemplate: pathtemp,
					HostTemplate: host,
					Methods:      meth,
				}
				w.Write([]byte(util.ToPrettyJson(rout)))
				return nil
			})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			w.Write([]byte(fmt.Sprintln("---------------------------------------------------------------------")))
		})
		fmt.Println("registered handler: ", "/debug/pprof/")
		r.Handle("/debug/pprof", http.HandlerFunc(pprof.Index))
		fmt.Println("registered handler: ", "/debug/pprof/cmdline")
		r.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
		fmt.Println("registered handler: ", "/debug/pprof/profile")
		r.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
		fmt.Println("registered handler: ", "/debug/pprof/symbol")
		r.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
		fmt.Println("registered handler: ", "/debug/pprof/trace")
		r.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
		return r
	}
}

func MetricsFunc() RouterMiddleware {
	return func(r *mux.Router) *mux.Router {

		var (
			inFlightGauge = prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "in_flight_requests",
				Help: "A gauge of requests currently being served by the wrapped handler.",
			})

			counter = prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Name: "api_requests_total",
					Help: "A counter for requests to the wrapped handler.",
				},
				[]string{"code", "method"},
			)

			// duration is partitioned by the HTTP method and handler. It uses custom
			// buckets based on the expected request duration.
			duration = prometheus.NewHistogramVec(
				prometheus.HistogramOpts{
					Name:    "request_duration_seconds",
					Help:    "A histogram of latencies for requests.",
					Buckets: []float64{.025, .05, .1, .25, .5, 1},
				},
				[]string{"handler", "method"},
			)

			// responseSize has no labels, making it a zero-dimensional
			// ObserverVec.
			responseSize = prometheus.NewHistogramVec(
				prometheus.HistogramOpts{
					Name:    "response_size_bytes",
					Help:    "A histogram of response sizes for requests.",
					Buckets: []float64{200, 500, 900, 1500},
				},
				[]string{},
			)
		)

		// Register all of the metrics in the standard registry.
		prometheus.MustRegister(inFlightGauge, counter, duration, responseSize)
		var chain http.Handler
		if err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			pth, _ := route.GetPathTemplate()
			chain = promhttp.InstrumentHandlerInFlight(inFlightGauge,
				promhttp.InstrumentHandlerDuration(duration.MustCurryWith(prometheus.Labels{"handler": pth}),
					promhttp.InstrumentHandlerCounter(counter,
						promhttp.InstrumentHandlerResponseSize(responseSize, route.GetHandler()),
					),
				),
			)
			route = route.Handler(chain)
			return nil
		}); err != nil {
			grpclog.Infof(err.Error())
		}
		fmt.Println("registered handler: ", "/metrics")
		r.Handle("/metrics", promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{}))
		return r
	}
}
