// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelhttptrace // import "go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace"

import (
	"context"
	"crypto/tls"
	"net/http/httptrace"
	"net/textproto"
	"sync"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace/internal/semconv"
)

// ScopeName is the instrumentation scope name.
const ScopeName = "go.opentelemetry.io/otel/instrumentation/httptrace"

// HTTP attributes.
var (
	HTTPStatus                 = attribute.Key("http.status")
	HTTPHeaderMIME             = attribute.Key("http.mime")
	HTTPRemoteAddr             = attribute.Key("http.remote")
	HTTPLocalAddr              = attribute.Key("http.local")
	HTTPConnectionReused       = attribute.Key("http.conn.reused")
	HTTPConnectionWasIdle      = attribute.Key("http.conn.wasidle")
	HTTPConnectionIdleTime     = attribute.Key("http.conn.idletime")
	HTTPConnectionStartNetwork = attribute.Key("http.conn.start.network")
	HTTPConnectionDoneNetwork  = attribute.Key("http.conn.done.network")
	HTTPConnectionDoneAddr     = attribute.Key("http.conn.done.addr")
	HTTPDNSAddrs               = attribute.Key("http.dns.addrs")
)

var hookMap = map[string]string{
	"http.dns":     "http.getconn",
	"http.connect": "http.getconn",
	"http.tls":     "http.getconn",
}

func parentHook(hook string) string { _ = "STUB: not implemented"; return "" }

// ClientTraceOption allows customizations to how the httptrace.Client
// collects information.
type ClientTraceOption interface {
	apply(*clientTracer)
}

type clientTraceOptionFunc func(*clientTracer)

func (fn clientTraceOptionFunc) apply(c *clientTracer) {
	_ = "STUB: not implemented"

	// WithoutSubSpans will modify the httptrace.ClientTrace to only collect data
	// as Events and Attributes on a span found in the context.  By default
	// sub-spans will be generated.
	return
}

func WithoutSubSpans() ClientTraceOption { _ = "STUB: not implemented"; return *new(ClientTraceOption) }

// WithRedactedHeaders will be replaced by fixed '****' values for the header
// names provided.  These are in addition to the sensitive headers already
// redacted by default: Authorization, WWW-Authenticate, Proxy-Authenticate
// Proxy-Authorization, Cookie, Set-Cookie.
func WithRedactedHeaders(headers ...string) ClientTraceOption {
	_ = "STUB: not implemented"
	return *new(ClientTraceOption)
}

// WithoutHeaders will disable adding span attributes for the http headers
// and values.
func WithoutHeaders() ClientTraceOption { _ = "STUB: not implemented"; return *new(ClientTraceOption) }

// WithInsecureHeaders will add span attributes for all http headers *INCLUDING*
// the sensitive headers that are redacted by default.  The attribute values
// will include the raw un-redacted text.  This might be useful for
// debugging authentication related issues, but should not be used for
// production deployments.
func WithInsecureHeaders() ClientTraceOption {
	_ = "STUB: not implemented"
	return *new(ClientTraceOption)
}

// WithTracerProvider specifies a tracer provider for creating a tracer.
// The global provider is used if none is specified.
func WithTracerProvider(provider trace.TracerProvider) ClientTraceOption {
	_ = "STUB: not implemented"
	return *new(ClientTraceOption)
}

type clientTracer struct {
	context.Context

	tracerProvider trace.TracerProvider

	tr trace.Tracer

	activeHooks     map[string]context.Context
	root            trace.Span
	mtx             sync.Mutex
	redactedHeaders map[string]struct{}
	addHeaders      bool
	useSpans        bool
	semconv         semconv.HTTPClient
}

// NewClientTrace returns an httptrace.ClientTrace implementation that will
// record OpenTelemetry spans for requests made by an http.Client. By default
// several spans will be added to the trace for various stages of a request
// (dns, connection, tls, etc). Also by default, all HTTP headers will be
// added as attributes to spans, although several headers will be automatically
// redacted: Authorization, WWW-Authenticate, Proxy-Authenticate,
// Proxy-Authorization, Cookie, and Set-Cookie.
func NewClientTrace(ctx context.Context, opts ...ClientTraceOption) *httptrace.ClientTrace {
	_ = "STUB: not implemented"
	return nil
}

func (ct *clientTracer) start(hook, spanName string, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

// end was called before start finished, add the start attributes and end the span here

func (ct *clientTracer) end(hook string, err error, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"

	// sometimes end may be called without previous start
	return
}

// start is not finished before end is called.
// Start a span here with the ending attributes that will be finished when start finishes.
// Yes, it's backwards. v0v

func (ct *clientTracer) getParentContext(hook string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (ct *clientTracer) span(hook string) trace.Span {
	_ = "STUB: not implemented"
	return *new(trace.Span)
}

func (ct *clientTracer) getConn(host string) { _ = "STUB: not implemented"; return }

func (ct *clientTracer) gotConn(info httptrace.GotConnInfo) { _ = "STUB: not implemented"; return }

func (ct *clientTracer) putIdleConn(err error) { _ = "STUB: not implemented"; return }

func (ct *clientTracer) gotFirstResponseByte() { _ = "STUB: not implemented"; return }

func (ct *clientTracer) dnsStart(info httptrace.DNSStartInfo) { _ = "STUB: not implemented"; return }

func (ct *clientTracer) dnsDone(info httptrace.DNSDoneInfo) { _ = "STUB: not implemented"; return }

func (ct *clientTracer) connectStart(network, addr string) { _ = "STUB: not implemented"; return }

func (ct *clientTracer) connectDone(network, addr string, err error) {
	_ = "STUB: not implemented"
	return
}

func (ct *clientTracer) tlsHandshakeStart() { _ = "STUB: not implemented"; return }

func (ct *clientTracer) tlsHandshakeDone(_ tls.ConnectionState, err error) {
	_ = "STUB: not implemented"
	return
}

func (ct *clientTracer) wroteHeaderField(k string, v []string) { _ = "STUB: not implemented"; return }

func (ct *clientTracer) wroteHeaders() { _ = "STUB: not implemented"; return }

func (ct *clientTracer) wroteRequest(info httptrace.WroteRequestInfo) {
	_ = "STUB: not implemented"
	return
}

func (ct *clientTracer) got100Continue() { _ = "STUB: not implemented"; return }

// It's possible that Got100Continue is called before GotFirstResponseByte at which point span can be `nil`.

func (ct *clientTracer) wait100Continue() { _ = "STUB: not implemented"; return }

// It's possible that Wait100Continue is called before GotFirstResponseByte at which point span can be `nil`.

func (ct *clientTracer) got1xxResponse(code int, header textproto.MIMEHeader) error {
	_ = "STUB: not implemented"
	return nil
}

// It's possible that Got1xxResponse is called before GotFirstResponseByte at which point span can be `nil`.

func sliceToString(value []string) string { _ = "STUB: not implemented"; return "" }

func sm2s(value map[string][]string) string { _ = "STUB: not implemented"; return "" }
