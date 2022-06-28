module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter

go 1.16

require (
	github.com/cenkalti/backoff/v4 v4.1.3
	github.com/elastic/go-elasticsearch/v7 v7.13.1
	github.com/elastic/go-structform v0.0.9
	github.com/pelletier/go-toml v1.9.3 // indirect
	github.com/stretchr/testify v1.7.4
	go.opentelemetry.io/collector v0.54.0
	go.opentelemetry.io/collector/model v0.31.0
	go.uber.org/multierr v1.8.0
	go.uber.org/zap v1.21.0
)
