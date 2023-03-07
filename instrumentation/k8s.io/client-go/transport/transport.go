/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package transport

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"k8s.io/client-go/transport"

	"github.com/adevjoe/opentelemetry-go-contrib/instrumentation/k8s.io/client-go/internal"
	"github.com/adevjoe/opentelemetry-go-contrib/instrumentation/k8s.io/client-go/option"
	"github.com/adevjoe/opentelemetry-go-contrib/instrumentation/k8s.io/client-go/semconv"
)

// instrumentationName is the instrumentation library identifier for a Tracer.
const instrumentationName = "client-go"

// NewWrapperFunc returns a Kubernetes WrapperFunc that can be used with a
// client configuration to trace all communication the client makes.
func NewWrapperFunc(opts ...option.Option) transport.WrapperFunc {
	return func(rt http.RoundTripper) http.RoundTripper {
		if rt == nil {
			rt = http.DefaultTransport
		}

		wrapped := roundTripper{
			RoundTripper: rt,
			cfg: internal.NewConfig(
				instrumentationName,
				localToInternal(opts)...,
			),
		}

		return &wrapped
	}
}

func localToInternal(opts []option.Option) []internal.Option {
	out := make([]internal.Option, len(opts))
	for i, o := range opts {
		out[i] = internal.Option(o)
	}
	return out
}

// roundTripper wraps an http.RoundTripper's requests with a span.
type roundTripper struct {
	http.RoundTripper

	cfg *internal.Config
}

var _ http.RoundTripper = (*roundTripper)(nil)

func (rt *roundTripper) RoundTrip(r *http.Request) (resp *http.Response, err error) {
	const nLocalOpts = 2
	opts := make([]trace.SpanStartOption, len(rt.cfg.DefaultStartOpts), len(rt.cfg.DefaultStartOpts)+nLocalOpts)
	copy(opts, rt.cfg.DefaultStartOpts)
	opts = append(
		opts,
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(semconv.ClientRequest(r)...),
	)

	tracer := rt.cfg.ResolveTracer(r.Context())
	ctx, span := tracer.Start(r.Context(), name(r), opts...)

	// Ensure anything downstream knows about the started span.
	r = r.WithContext(ctx)
	rt.cfg.Propagator.Inject(ctx, propagation.HeaderCarrier(r.Header))

	resp, err = rt.RoundTripper.RoundTrip(r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.End()
		return
	}

	if resp.StatusCode < http.StatusMultipleChoices {
		span.SetAttributes(semconv.ClientResponse(resp)...)
		span.SetStatus(codes.Ok, codes.Ok.String())
	} else {
		span.SetAttributes(semconv.ClientResponse(resp)...)
		span.SetStatus(codes.Error, fmt.Sprintf("Invalid HTTP status code %d", resp.StatusCode))
	}

	resp.Body = &wrappedBody{ctx: ctx, span: span, body: resp.Body}

	return
}

const (
	prefixAPI   = "/api/v1/"
	prefixWatch = "watch/"
)

// name returns an appropriate span name based on the client request.
// OpenTelemetry semantic conventions require this name to be low cardinality,
// but since the Kubernetes API is somewhat predictable we can usually return
// more than just "HTTP {METHOD}".
func name(r *http.Request) string {
	path := r.URL.Path
	method := r.Method

	if !strings.HasPrefix(path, prefixAPI) {
		return "HTTP " + method
	}

	var out strings.Builder
	out.WriteString("HTTP " + method + " ")

	path = strings.TrimPrefix(path, prefixAPI)

	if strings.HasPrefix(path, prefixWatch) {
		path = strings.TrimPrefix(path, prefixWatch)
		out.WriteString(prefixWatch)
	}

	// For each {type}/{name}, tokenize the {name} portion.
	var previous string
	for i, part := range strings.Split(path, "/") {
		if i > 0 {
			out.WriteRune('/')
		}

		if i%2 == 0 {
			out.WriteString(part)
			previous = part
		} else {
			out.WriteString(tokenize(previous))
		}
	}

	return out.String()
}

func tokenize(k8Type string) string {
	switch k8Type {
	case "namespaces":
		return "{namespace}"
	case "proxy":
		return "{path}"
	default:
		return "{name}"
	}
}

type wrappedBody struct {
	ctx  context.Context
	span trace.Span
	body io.ReadCloser
}

var _ io.ReadCloser = (*wrappedBody)(nil)

func (wb *wrappedBody) Read(b []byte) (int, error) {
	n, err := wb.body.Read(b)
	switch err {
	case nil:
		// nothing to do here but fall through to the return
	case io.EOF:
		wb.span.End()
	default:
		wb.span.RecordError(err)
		wb.span.SetStatus(codes.Error, err.Error())
	}

	return n, err
}

func (wb *wrappedBody) Close() error {
	wb.span.End()
	return wb.body.Close()
}
