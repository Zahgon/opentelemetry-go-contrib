// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelaws // import "go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws"

import (
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type config struct {
	TracerProvider    trace.TracerProvider
	TextMapPropagator propagation.TextMapPropagator
	AttributeBuilders []AttributeBuilder
}

// Option applies an option value.
type Option interface {
	apply(*config)
}

// optionFunc provides a convenience wrapper for simple Options
// that can be represented as functions.
type optionFunc func(*config)

func (o optionFunc) apply(c *config) {
	_ = "STUB: not implemented"

	// WithTracerProvider specifies a tracer provider to use for creating a tracer.
	// If none is specified, the global TracerProvider is used.
	return
}

func WithTracerProvider(provider trace.TracerProvider) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithTextMapPropagator specifies a Text Map Propagator to use when propagating context.
// If none is specified, the global TextMapPropagator is used.
func WithTextMapPropagator(propagator propagation.TextMapPropagator) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithAttributeBuilder specifies an attribute setter function for setting service specific attributes.
// If none is specified, the service will be determined by the DefaultAttributeBuilder function and the corresponding attributes will be included.
func WithAttributeBuilder(attributeBuilders ...AttributeBuilder) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
