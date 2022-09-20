module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/honeycombexporter

go 1.16

require (
	github.com/census-instrumentation/opencensus-proto v0.3.0
	github.com/google/go-cmp v0.5.8
	github.com/honeycombio/libhoney-go v1.15.3
	github.com/klauspost/compress v1.15.9
	github.com/pelletier/go-toml v1.9.3 // indirect
	github.com/stretchr/testify v1.8.0
	go.opentelemetry.io/collector v0.60.0
	go.opentelemetry.io/collector/model v0.31.0
	go.uber.org/zap v1.23.0
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.1
)
