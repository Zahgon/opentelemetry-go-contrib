// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelmongo // import "go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/v2/mongo/otelmongo"

import (
	"go.mongodb.org/mongo-driver/v2/event"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// ScopeName is the instrumentation scope name.
const ScopeName = "go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/v2/mongo/otelmongo"

// config is used to configure the mongo tracer.
type config struct {
	MeterProvider  metric.MeterProvider
	TracerProvider trace.TracerProvider

	Meter  metric.Meter
	Tracer trace.Tracer

	CommandAttributeDisabled bool

	SpanNameFormatter SpanNameFormatterFunc
}

// newConfig returns a config with all Options set.
func newConfig(opts ...Option) config { _ = "STUB: not implemented"; return *new(config) }

// Option specifies instrumentation configuration options.
type Option interface {
	apply(*config)
}

type optionFunc func(*config)

func (o optionFunc) apply(c *config) {
	_ = "STUB: not implemented"

	// WithMeterProvider specifies a [metric.MeterProvider] to use for creating a Meter.
	// If none is specified, the global MeterProvider is used.
	return
}

func WithMeterProvider(provider metric.MeterProvider) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// SpanNameFormatterFunc is a function that resolves the span name given an
// *event.CommandStartedEvent.
type SpanNameFormatterFunc func(e *event.CommandStartedEvent) string

// WithSpanNameFormatter specifies a function that resolves the span name given an
// *event.CommandStartedEvent. If none is specified, the default resolver is used,
// which returns "<collection>.<command>" if the collection is non-empty,
// and just "<command>" otherwise.
func WithSpanNameFormatter(resolver SpanNameFormatterFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithTracerProvider specifies a tracer provider to use for creating a tracer.
// If none is specified, the global provider is used.
func WithTracerProvider(provider trace.TracerProvider) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithCommandAttributeDisabled specifies if the MongoDB command is added as an attribute to Spans or not.
// This is disabled by default and the MongoDB command will not be added as an attribute
// to Spans if this option is not provided.
func WithCommandAttributeDisabled(disabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
