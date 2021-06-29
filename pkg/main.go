package main

import (
	"os"

	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

var telemetryChan chan TelemetryFrame
var telemetryErrorChan chan error

func main() {
	telemetryChan = make(chan TelemetryFrame)
	telemetryErrorChan = make(chan error)
	go RunTelemetryServer(telemetryChan, telemetryErrorChan)

	// Start listening to requests sent from Grafana. This call is blocking so
	// it won't finish until Grafana shuts down the process or the plugin choose
	// to exit by itself using os.Exit. Manage automatically manages life cycle
	// of datasource instances. It accepts datasource instance factory as first
	// argument. This factory will be automatically called on incoming request
	// from Grafana to create different instances of SimracingTelemetryDatasource (per datasource
	// ID). When datasource configuration changed Dispose method will be called and
	// new datasource instance created using NewSimracingTelemetryDatasource factory.
	if err := datasource.Manage("simracing-telemetry-datasource", NewSimracingTelemetryDatasource, datasource.ManageOpts{}); err != nil {
		log.DefaultLogger.Error(err.Error())
		os.Exit(1)
	}
}
