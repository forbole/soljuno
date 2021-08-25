package logging

import (
	"github.com/prometheus/client_golang/prometheus"
)

// StartSlot represents the Telemetry counter used to set the start slot of the parsing
var StartSlot = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "soljuno_initial_slot",
		Help: "Initial parsing slot.",
	},
)

// WorkerCount represents the Telemetry counter used to track the worker count
var WorkerCount = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "soljuno_worker_count",
		Help: "Number of active workers.",
	},
)

// WorkerHeight represents the Telemetry counter used to track the last indexed slot for each worker
var WorkerHeight = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "soljuno_last_indexed_slot",
		Help: "Height of the last indexed block.",
	},
	[]string{"worker_index"},
)

// ErrorCount represents the Telemetry counter used to track the number of errors emitted
var ErrorCount = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "soljuno_error_count",
		Help: "Total number of errors emitted.",
	},
)

func init() {
	err := prometheus.Register(StartSlot)
	if err != nil {
		panic(err)
	}

	err = prometheus.Register(WorkerCount)
	if err != nil {
		panic(err)
	}

	err = prometheus.Register(WorkerHeight)
	if err != nil {
		panic(err)
	}

	err = prometheus.Register(ErrorCount)
	if err != nil {
		panic(err)
	}
}
