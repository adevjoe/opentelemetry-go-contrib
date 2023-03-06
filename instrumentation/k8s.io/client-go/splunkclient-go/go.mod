module github.com/adevjoe/opentelemetry-go-contrib/instrumentation/k8s.io/client-go/splunkclient-go

go 1.17

require (
	github.com/signalfx/splunk-otel-go v1.4.0
	go.opentelemetry.io/otel v1.14.0
	go.opentelemetry.io/otel/trace v1.14.0
	k8s.io/client-go v0.23.2
)

require (
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/oauth2 v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	k8s.io/apimachinery v0.26.2 // indirect
	k8s.io/klog/v2 v2.90.1 // indirect
	k8s.io/utils v0.0.0-20230220204549-a5ecb0141aa5 // indirect
)

replace (
	github.com/signalfx/splunk-otel-go => ../../../../
	github.com/signalfx/splunk-otel-go/instrumentation/internal => ../../../internal/
)
