// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelmongo // import "go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"

import (
	"go.opentelemetry.io/otel/trace"
)

// ScopeName is the instrumentation scope name.
const ScopeName = "go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"

// config is used to configure the mongo tracer.
type config struct {
	TracerProvider trace.TracerProvider

	Tracer trace.Tracer

	CommandAttributeDisabled bool
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

	// WithTracerProvider specifies a tracer provider to use for creating a tracer.
	// If none is specified, the global provider is used.
	return
}

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
