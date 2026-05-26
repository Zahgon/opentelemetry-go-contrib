// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelhttp // import "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

import (
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp/internal/semconv"
)

// middleware is an http middleware which wraps the next handler in a span.
type middleware struct {
	operation string
	server    string

	tracer             trace.Tracer
	propagators        propagation.TextMapPropagator
	spanStartOptions   []trace.SpanStartOption
	readEvent          bool
	writeEvent         bool
	filters            []Filter
	spanNameFormatter  func(string, *http.Request) string
	publicEndpointFn   func(*http.Request) bool
	metricAttributesFn func(*http.Request) []attribute.KeyValue

	semconv semconv.HTTPServer
}

// NewHandler wraps the passed handler in a span named after the operation and
// enriches it with metrics.
func NewHandler(handler http.Handler, operation string, opts ...Option) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// NewMiddleware returns a tracing and metrics instrumentation middleware.
// The handler returned by the middleware wraps a handler
// in a span named after the operation and enriches it with metrics.
func NewMiddleware(operation string, opts ...Option) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func (h *middleware) configure(c *config) { _ = "STUB: not implemented"; return }

// serveHTTP sets up tracing and calls the given next http.Handler with the span
// context injected into the request context.
func (h *middleware) serveHTTP(w http.ResponseWriter, r *http.Request, next http.Handler) {
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

func (h *middleware) metricAttributesFromRequest(r *http.Request) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}
