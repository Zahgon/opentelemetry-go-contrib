// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package otelhttptrace provides instrumentation for the [net/http/httptrace]
// package.
package otelhttptrace // import "go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace"

import (
	"context"
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// Option allows configuration of the httptrace Extract()
// and Inject() functions.
type Option interface {
	apply(*config)
}

type optionFunc func(*config)

func (o optionFunc) apply(c *config) { _ = "STUB: not implemented"; return }

type config struct {
	propagators propagation.TextMapPropagator
}

func newConfig(opts []Option) *config { _ = "STUB: not implemented"; return nil }

// WithPropagators sets the propagators to use for Extraction and Injection.
func WithPropagators(props propagation.TextMapPropagator) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Extract returns the Attributes, Context Entries, and SpanContext that were encoded by Inject.
func Extract(ctx context.Context, req *http.Request, opts ...Option) ([]attribute.KeyValue, baggage.Baggage, trace.SpanContext) {
	_ = "STUB: not implemented"
	return nil, *new(baggage.Baggage), *new(trace.SpanContext)
}

// Inject sets attributes, context entries, and span context from ctx into
// the request.
func Inject(ctx context.Context, req *http.Request, opts ...Option) {
	_ = "STUB: not implemented"
	return
}
