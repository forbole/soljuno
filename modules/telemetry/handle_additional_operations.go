package telemetry

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	_ "github.com/mkevac/debugcharts"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// RunAdditionalOperations runs the module additional operations
func RunAdditionalOperations(cfg Config) error {
	err := checkConfig(cfg)
	if err != nil {
		return err
	}

	go startPrometheus(cfg)

	return nil
}

// checkConfig checks if the given config is valid
func checkConfig(cfg Config) error {
	if cfg == nil {
		return fmt.Errorf("no telemetry config found")
	}

	return nil
}

// startPrometheus starts a Prometheus server using the given configuration
func startPrometheus(cfg Config) {
	router := mux.NewRouter()
	router.Handle("/metrics", promhttp.Handler())
	router.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)
	router.PathPrefix("/debug/charts").Handler(http.DefaultServeMux)

	port := cfg.GetTelemetryConfig().GetPort()
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	if err != nil {
		panic(err)
	}
}
