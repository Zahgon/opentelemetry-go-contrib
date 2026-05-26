// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelmux // import "go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux/internal/semconv"
)

const (
	// ScopeName is the instrumentation scope name.
	ScopeName = "go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
)

// Middleware sets up a handler to start tracing the incoming
// requests.  The service parameter should describe the name of the
// (virtual) server handling the request.
func Middleware(service string, opts ...Option) mux.MiddlewareFunc {
	_ = "STUB: not implemented"
	return *new(mux.MiddlewareFunc)
}

type traceware struct {
	service            string
	tracer             trace.Tracer
	propagators        propagation.TextMapPropagator
	handler            http.Handler
	spanNameFormatter  func(string, *http.Request) string
	publicEndpoint     bool
	publicEndpointFn   func(*http.Request) bool
	filters            []Filter
	meter              metric.Meter
	semconv            semconv.HTTPServer
	metricAttributesFn func(*http.Request) []attribute.KeyValue
}

// validMethods are all the OTel recognized HTTP methods.
var validMethods = map[string]struct{}{
	http.MethodGet:     {},
	http.MethodHead:    {},
	http.MethodPost:    {},
	http.MethodPut:     {},
	http.MethodPatch:   {},
	http.MethodDelete:  {},
	http.MethodConnect: {},
	http.MethodOptions: {},
	http.MethodTrace:   {},
}

// defaultSpanNameFunc returns the semconv based default span name.
func defaultSpanNameFunc(routeName string, r *http.Request) string {
	_ = "STUB: not implemented"
	return ""
}

// ServeHTTP implements the http.Handler interface. It does the actual
// tracing of the request.
func (tw traceware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Simply pass through to the handler if a filter rejects the request

// Linking incoming span context if any for public endpoint.

// if request body is nil or NoBody, we don't want to mutate the body as it
// will affect the identity of it in an unforeseeable way because we assert
// ReadCloser fulfills a certain interface and it is indeed nil or NoBody.

// Wrap w to use our ResponseWriter methods while also exposing
// other interfaces that w may implement (http.CloseNotifier,
// http.Flusher, http.Hijacker, http.Pusher, io.ReaderFrom).

func (tw traceware) metricAttributesFromRequest(r *http.Request) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

func extractRoute(r *http.Request) string { _ = "STUB: not implemented"; return "" }
