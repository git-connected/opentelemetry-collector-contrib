module github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/hostobserver

go 1.16

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer v0.0.0-00010101000000-000000000000
	github.com/shirou/gopsutil v3.21.7+incompatible
	github.com/stretchr/testify v1.7.0
	github.com/tklauser/go-sysconf v0.3.6 // indirect
	go.opentelemetry.io/collector v0.34.0
	go.uber.org/zap v1.19.0
	golang.org/x/tools v0.1.3 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer => ../
