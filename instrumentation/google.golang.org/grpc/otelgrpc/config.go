// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelgrpc // import "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/stats"
)

// ScopeName is the instrumentation scope name.
const ScopeName = "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

// InterceptorFilter is a predicate used to determine whether a given request in
// interceptor info should be instrumented. A InterceptorFilter must return true if
// the request should be traced.
//
// Deprecated: Use stats handlers instead.
type InterceptorFilter func(*InterceptorInfo) bool

// Filter is a predicate used to determine whether a given request in
// should be instrumented by the attached RPC tag info.
// A Filter must return true if the request should be instrumented.
type Filter func(*stats.RPCTagInfo) bool

type semconvMode int

const (
	semconvModeNew semconvMode = iota // Default
	semconvModeOld
	semconvModeDup
)

// config is a group of options for this instrumentation.
type config struct {
	Filter             Filter
	InterceptorFilter  InterceptorFilter
	Propagators        propagation.TextMapPropagator
	TracerProvider     trace.TracerProvider
	MeterProvider      metric.MeterProvider
	SpanKind           trace.SpanKind
	SpanStartOptions   []trace.SpanStartOption
	SpanAttributes     []attribute.KeyValue
	MetricAttributes   []attribute.KeyValue
	MetricAttributesFn func(ctx context.Context) []attribute.KeyValue

	PublicEndpoint   bool
	PublicEndpointFn func(ctx context.Context, info *stats.RPCTagInfo) bool

	ReceivedEvent bool
	SentEvent     bool

	semconvMode semconvMode
}

// Option applies an option value for a config.
type Option interface {
	apply(*config)
}

type optionFunc func(*config)

func (f optionFunc) apply(c *config) {
	_ = "STUB: not implemented"

	// newConfig returns a config configured with all the passed Options.
	return
}

func newConfig(opts []Option) *config { _ = "STUB: not implemented"; return nil }

func parseSemconvMode() semconvMode { _ = "STUB: not implemented"; return *new(semconvMode) }

// WithPublicEndpoint configures the Handler to link the span with an incoming
// span context. If this option is not provided, then the association is a child
// association instead of a link.
func WithPublicEndpoint() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPublicEndpointFn runs with every request, and allows conditionally
// configuring the Handler to link the span with an incoming span context. If
// this option is not provided or returns false, then the association is a
// child association instead of a link.
// Note: WithPublicEndpoint takes precedence over WithPublicEndpointFn.
func WithPublicEndpointFn(fn func(context.Context, *stats.RPCTagInfo) bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithPropagators returns an Option to use the Propagators when extracting
// and injecting trace context from requests.
func WithPropagators(p propagation.TextMapPropagator) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithInterceptorFilter returns an Option to use the request filter.
//
// Deprecated: Use stats handlers instead.
func WithInterceptorFilter(f InterceptorFilter) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithFilter returns an Option to use the request filter.
func WithFilter(f Filter) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTracerProvider returns an Option to use the TracerProvider when
// creating a Tracer.
func WithTracerProvider(tp trace.TracerProvider) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithMeterProvider returns an Option to use the MeterProvider when
// creating a Meter. If this option is not provide the global MeterProvider will be used.
func WithMeterProvider(mp metric.MeterProvider) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Event type that can be recorded, see WithMessageEvents.
type Event int

// Different types of events that can be recorded, see WithMessageEvents.
const (
	ReceivedEvents Event = iota
	SentEvents
)

// WithMessageEvents configures the Handler to record the specified events
// (span.AddEvent) on spans. By default only summary attributes are added at the
// end of the request.
//
// Valid events are:
//   - ReceivedEvents: Record the number of bytes read after every gRPC read operation.
//   - SentEvents: Record the number of bytes written after every gRPC write operation.
func WithMessageEvents(events ...Event) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSpanOptions configures an additional set of
// trace.SpanOptions, which are applied to each new span.
//
// Deprecated: It is only used by the deprecated interceptor, and is unused by [NewClientHandler] and [NewServerHandler].
func WithSpanOptions(opts ...trace.SpanStartOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSpanKind returns an Option to set the span kind for spans created by
// the handler.
//
// By default, [NewServerHandler] creates spans with
// [trace.SpanKindServer] and [NewClientHandler] creates spans with
// [trace.SpanKindClient].
func WithSpanKind(sk trace.SpanKind) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSpanAttributes returns an Option to add custom attributes to the spans.
func WithSpanAttributes(a ...attribute.KeyValue) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithMetricAttributes returns an Option to add custom attributes to the metrics.
func WithMetricAttributes(a ...attribute.KeyValue) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithMetricAttributesFn returns an Option to add dynamic custom attributes to the handler's metrics.
// The function is called once per RPC and the returned attributes are applied to all metrics recorded by this handler.
//
// The context parameter is the standard gRPC request context and provides access to request-scoped data.
func WithMetricAttributesFn(fn func(ctx context.Context) []attribute.KeyValue) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
