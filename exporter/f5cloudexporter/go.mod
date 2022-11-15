module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/f5cloudexporter

go 1.16

require (
	cloud.google.com/go/compute/metadata v0.2.1 // indirect
	github.com/pelletier/go-toml v1.9.3 // indirect
	github.com/stretchr/testify v1.8.1
	go.opentelemetry.io/collector v0.64.1
	go.opentelemetry.io/collector/exporter/otlphttpexporter v0.64.1
	go.uber.org/zap v1.23.0
	golang.org/x/oauth2 v0.0.0-20221014153046-6fdb5e3db783
	google.golang.org/api v0.100.0
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)
