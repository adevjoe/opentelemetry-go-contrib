module github.com/adevjoe/opentelemetry-go-contrib/instrumentation/k8s.io/client-go

go 1.17

replace go.opentelemetry.io/otel => go.opentelemetry.io/otel v1.7.0

replace go.opentelemetry.io/otel/trace => go.opentelemetry.io/otel/trace v1.7.0

require (
	github.com/signalfx/splunk-otel-go v1.4.0
	go.opentelemetry.io/otel v1.14.0
	go.opentelemetry.io/otel/trace v1.14.0
	k8s.io/client-go v0.26.2
)

require (
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/oauth2 v0.0.0-20220223155221-ee480838109b // indirect
	golang.org/x/text v0.7.0 // indirect
	golang.org/x/time v0.0.0-20220210224613-90d013bbcef8 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	k8s.io/apimachinery v0.26.2 // indirect
	k8s.io/klog/v2 v2.80.1 // indirect
	k8s.io/utils v0.0.0-20221107191617-1a15be271d1d // indirect
)
