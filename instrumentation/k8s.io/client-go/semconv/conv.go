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

package semconv

import (
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

var (
	nc = &NetConv{
		NetHostNameKey:     semconv.NetHostNameKey,
		NetHostPortKey:     semconv.NetHostPortKey,
		NetPeerNameKey:     semconv.NetPeerNameKey,
		NetPeerPortKey:     semconv.NetPeerPortKey,
		NetTransportOther:  semconv.NetTransportOther,
		NetTransportTCP:    semconv.NetTransportTCP,
		NetTransportUDP:    semconv.NetTransportUDP,
		NetTransportInProc: semconv.NetTransportInProc,
	}

	hc = &HTTPConv{
		NetConv: nc,

		EnduserIDKey:                 semconv.EnduserIDKey,
		HTTPClientIPKey:              semconv.HTTPClientIPKey,
		HTTPFlavorKey:                semconv.HTTPFlavorKey,
		HTTPMethodKey:                semconv.HTTPMethodKey,
		HTTPRequestContentLengthKey:  semconv.HTTPRequestContentLengthKey,
		HTTPResponseContentLengthKey: semconv.HTTPResponseContentLengthKey,
		HTTPRouteKey:                 semconv.HTTPRouteKey,
		HTTPSchemeHTTP:               semconv.HTTPSchemeHTTP,
		HTTPSchemeHTTPS:              semconv.HTTPSchemeHTTPS,
		HTTPStatusCodeKey:            semconv.HTTPStatusCodeKey,
		HTTPTargetKey:                semconv.HTTPTargetKey,
		HTTPURLKey:                   semconv.HTTPURLKey,
		HTTPUserAgentKey:             semconv.HTTPUserAgentKey,
	}
)

// ClientResponse returns attributes for an HTTP response received by a client
// from a server. It will return the following attributes if the related values
// are defined in resp: "http.status.code", "http.response_content_length".
//
// This does not add all OpenTelemetry required attributes for an HTTP event,
// it assumes ClientRequest was used to create the span with a complete set of
// attributes. If a complete set of attributes can be generated using the
// request contained in resp. For example:
//
//	append(ClientResponse(resp), ClientRequest(resp.Request)...)
func ClientResponse(resp *http.Response) []attribute.KeyValue {
	return hc.ClientResponse(resp)
}

// ClientRequest returns attributes for an HTTP request made by a client. The
// following attributes are always returned: "http.url", "http.flavor",
// "http.method", "net.peer.name". The following attributes are returned if the
// related values are defined in req: "net.peer.port", "http.user_agent",
// "http.request_content_length", "enduser.id".
func ClientRequest(req *http.Request) []attribute.KeyValue {
	return hc.ClientRequest(req)
}

// ClientStatus returns a span status code and message for an HTTP status code
// value received by a client.
func ClientStatus(code int) (codes.Code, string) {
	return hc.ClientStatus(code)
}
