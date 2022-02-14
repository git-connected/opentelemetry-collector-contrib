module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/simpleprometheusreceiver/examples/federation/prom-counter

go 1.16

require (
	go.opentelemetry.io/otel v1.4.0
	go.opentelemetry.io/otel/exporters/prometheus v0.27.0
	go.opentelemetry.io/otel/metric v0.27.0
	go.opentelemetry.io/otel/sdk/export/metric v0.22.0
	go.opentelemetry.io/otel/sdk/metric v0.27.0
	go.uber.org/zap v1.18.1
	google.golang.org/protobuf v1.27.1 // indirect
)
