// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelhttp // import "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

import (
	"context"
	"io"
	"net/http"
	"net/http/httptrace"
	"sync/atomic"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp/internal/semconv"
)

// Transport implements the http.RoundTripper interface and wraps
// outbound HTTP(S) requests with a span and enriches it with metrics.
type Transport struct {
	rt http.RoundTripper

	tracer             trace.Tracer
	propagators        propagation.TextMapPropagator
	spanStartOptions   []trace.SpanStartOption
	filters            []Filter
	spanNameFormatter  func(string, *http.Request) string
	clientTrace        func(context.Context) *httptrace.ClientTrace
	metricAttributesFn func(*http.Request) []attribute.KeyValue

	semconv semconv.HTTPClient
}

var _ http.RoundTripper = &Transport{}

// NewTransport wraps the provided http.RoundTripper with one that
// starts a span, injects the span context into the outbound request headers,
// and enriches it with metrics.
//
// If the provided http.RoundTripper is nil, http.DefaultTransport will be used
// as the base http.RoundTripper.
func NewTransport(base http.RoundTripper, opts ...Option) *Transport {
	_ = "STUB: not implemented"
	return nil
}

func (t *Transport) applyConfig(c *config) { _ = "STUB: not implemented"; return }

func defaultTransportFormatter(_ string, r *http.Request) string {
	_ = "STUB: not implemented"
	return ""

	// RoundTrip creates a Span and propagates its context via the provided request's headers
	// before handing the request to the configured base RoundTripper. The created span will
	// end when the response body is closed or when a read from the body returns io.EOF.
}

func (t *Transport) RoundTrip(r *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Simply pass through to the base RoundTripper if a filter rejects the request

// According to RoundTripper spec, we shouldn't modify the origin request.

// Records the last body wrapper. Can be nil.

// The underlying transport will fail to make a retry request, hence, record no data.

// Record the metrics on error or no error.

// traces

func (t *Transport) metricAttributesFromRequest(r *http.Request) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

// newWrappedBody returns a new and appropriately scoped *wrappedBody as an
// io.ReadCloser. If the passed body implements io.Writer, the returned value
// will implement io.ReadWriteCloser.
func newWrappedBody(span trace.Span, record func(n int64), body io.ReadCloser) io.ReadCloser {
	_ = "STUB: not implemented"
	// The successful protocol switch responses will have a body that
	// implement an io.ReadWriteCloser. Ensure this interface type continues
	// to be satisfied if that is the case.
	return *new(io.ReadCloser)
}

// Remove the implementation of the io.ReadWriteCloser and only implement
// the io.ReadCloser.

// wrappedBody is the response body type returned by the transport
// instrumentation to complete a span. Errors encountered when using the
// response body are recorded in span tracking the response.
//
// The span tracking the response is ended when this body is closed.
//
// If the response body implements the io.Writer interface (i.e. for
// successful protocol switches), the wrapped body also will.
type wrappedBody struct {
	span     trace.Span
	recorded atomic.Bool
	record   func(n int64)
	body     io.ReadCloser
	read     atomic.Int64
}

var _ io.ReadWriteCloser = &wrappedBody{}

func (wb *wrappedBody) Write(p []byte) (int, error) {
	_ = "STUB: not implemented"
	// This will not panic given the guard in newWrappedBody.
	return 0, nil
}

func (wb *wrappedBody) Read(b []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Record the number of bytes read
		nil
}

// nothing to do here but fall through to the return

// recordBytesRead is a function that ensures the number of bytes read is recorded once and only once.
func (wb *wrappedBody) recordBytesRead() {
	_ = "STUB: not implemented"
	// note: it is more performant (and equally correct) to use atomic.Bool over sync.Once here. In the event that
	// two goroutines are racing to call this method, the number of bytes read will no longer increase. Using
	// CompareAndSwap allows later goroutines to return quickly and not block waiting for the race winner to finish
	// calling wb.record(wb.read.Load()).
	return
}

// Record the total number of bytes read

func (wb *wrappedBody) Close() error { _ = "STUB: not implemented"; return nil }
