// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelrestful // import "go.opentelemetry.io/contrib/instrumentation/github.com/emicklei/go-restful/otelrestful"

import (
	"net/http"

	"go.opentelemetry.io/otel/propagation"
	oteltrace "go.opentelemetry.io/otel/trace"
)

// config is used to configure the go-restful middleware.
type config struct {
	TracerProvider   oteltrace.TracerProvider
	Propagators      propagation.TextMapPropagator
	PublicEndpoint   bool
	PublicEndpointFn func(*http.Request) bool
}

// Option applies a configuration value.
type Option interface {
	apply(*config)
}

type optionFunc func(*config)

func (o optionFunc) apply(c *config) {
	_ = "STUB: not implemented"

	// WithPublicEndpoint configures the Handler to link the span with an incoming
	// span context. If this option is not provided, then the association is a child
	// association instead of a link.
	return
}

func WithPublicEndpoint() Option { _ = "STUB: not implemented"; return *new(Option) }

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

// WithPublicEndpointFn runs with every request, and allows conditionally
// configuring the Handler to link the span with an incoming span context. If
// this option is not provided or returns false, then the association is a
// child association instead of a link.
// Note: [WithPublicEndpoint] takes precedence over WithPublicEndpointFn.
func WithPublicEndpointFn(fn func(*http.Request) bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
