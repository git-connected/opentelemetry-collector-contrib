module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/honeycombexporter

go 1.16

require (
	github.com/census-instrumentation/opencensus-proto v0.3.0
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/go-cmp v0.5.6
	github.com/honeycombio/libhoney-go v1.15.3
	github.com/klauspost/compress v1.13.1
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.36.0
	go.opentelemetry.io/collector/model v0.36.0
	go.uber.org/zap v1.19.1
	google.golang.org/protobuf v1.27.1
)
