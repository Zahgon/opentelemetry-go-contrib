// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelmux // import "go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"

import (
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	oteltrace "go.opentelemetry.io/otel/trace"
)

// config is used to configure the mux middleware.
type config struct {
	TracerProvider     oteltrace.TracerProvider
	Propagators        propagation.TextMapPropagator
	spanNameFormatter  func(string, *http.Request) string
	PublicEndpoint     bool
	PublicEndpointFn   func(*http.Request) bool
	Filters            []Filter
	MeterProvider      metric.MeterProvider
	MetricAttributesFn func(*http.Request) []attribute.KeyValue
}

// Option specifies instrumentation configuration options.
type Option interface {
	apply(*config)
}

type optionFunc func(*config)

func (o optionFunc) apply(c *config) {
	_ = "STUB: not implemented"

	// Filter is a predicate used to determine whether a given http.request should
	// be traced. A Filter must return true if the request should be traced.
	return
}

type Filter func(*http.Request) bool

// WithPublicEndpoint configures the Handler to link the span with an incoming
// span context. If this option is not provided, then the association is a child
// association instead of a link.
func WithPublicEndpoint() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPublicEndpointFn runs with every request, and allows conditionally
// configuring the Handler to link the span with an incoming span context. If
// this option is not provided or returns false, then the association is a
// child association instead of a link.
// Note: WithPublicEndpoint takes precedence over WithPublicEndpointFn.
func WithPublicEndpointFn(fn func(*http.Request) bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithPropagators specifies propagators to use for extracting
// information from the HTTP requests. If none are specified, global
// ones will be used.
func WithPropagators(propagators propagation.TextMapPropagator) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithTracerProvider specifies a tracer provider to use for creating a tracer.
// If none is specified, the global provider is used.
func WithTracerProvider(provider oteltrace.TracerProvider) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSpanNameFormatter specifies a function to use for generating a custom span
// name. By default, the route name (path template or regexp) is used. The route
// name is provided so you can use it in the span name without needing to
// duplicate the logic for extracting it from the request.
func WithSpanNameFormatter(fn func(routeName string, r *http.Request) string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithFilter adds a filter to the list of filters used by the handler.
// If any filter indicates to exclude a request then the request will not be
// traced. All filters must allow a request to be traced for a Span to be created.
// If no filters are provided then all requests are traced.
// Filters will be invoked for each processed request, it is advised to make them
// simple and fast.
func WithFilter(f Filter) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMeterProvider specifies a meter provider to use for creating a metric.
// If none is specified, the global provider is used.
func WithMeterProvider(provider metric.MeterProvider) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithMetricAttributesFn returns an Option to set a function that maps an HTTP request to a slice of attribute.KeyValue.
// These attributes will be included in metrics for every request.
func WithMetricAttributesFn(metricAttributesFn func(r *http.Request) []attribute.KeyValue) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
