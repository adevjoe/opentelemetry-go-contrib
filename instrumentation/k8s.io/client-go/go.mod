module github.com/adevjoe/opentelemetry-go-contrib/instrumentation/k8s.io/client-go

go 1.17

replace go.opentelemetry.io/otel => go.opentelemetry.io/otel v1.7.0

replace go.opentelemetry.io/otel/trace => go.opentelemetry.io/otel/trace v1.7.0

require (
	github.com/signalfx/splunk-otel-go v1.4.0
	go.opentelemetry.io/otel v1.7.0
	go.opentelemetry.io/otel/trace v1.7.0
	k8s.io/client-go v0.23.1
)

require (
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20211209124913-491a49abca63 // indirect
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	k8s.io/apimachinery v0.23.1 // indirect
	k8s.io/klog/v2 v2.30.0 // indirect
	k8s.io/utils v0.0.0-20210930125809-cb0fa318a74b // indirect
)
