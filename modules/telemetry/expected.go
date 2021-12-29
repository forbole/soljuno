package telemetry

import "github.com/forbole/soljuno/types"

type Config interface {
	GetTelemetryConfig() types.TelemetryConfig
}
