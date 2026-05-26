// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package opencensus provides an OpenCensus trace context propagator.
package opencensus // import "go.opentelemetry.io/contrib/propagators/opencensus"

import (
	"context"

	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type key uint

const binaryKey key = 0

// binaryHeader is the same as traceContextKey is in opencensus:
// https://github.com/census-instrumentation/opencensus-go/blob/3fb168f674736c026e623310bfccb0691e6dec8a/plugin/ocgrpc/trace_common.go#L30
const binaryHeader = "grpc-trace-bin"

// Binary is an OpenTelemetry implementation of the OpenCensus grpc binary format.
// Binary propagation was temporarily removed from opentelemetry.  See
// https://github.com/open-telemetry/opentelemetry-specification/issues/437
type Binary struct{}

var _ propagation.TextMapPropagator = Binary{}

// Inject injects context into the TextMapCarrier.
func (Binary) Inject(ctx context.Context, carrier propagation.TextMapCarrier) {
	_ = "STUB: not implemented"
	return
}

// Extract extracts the SpanContext from the TextMapCarrier.
func (b Binary) Extract(ctx context.Context, carrier propagation.TextMapCarrier) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (Binary) extract(carrier propagation.TextMapCarrier) trace.SpanContext {
	_ = "STUB: not implemented"
	return *new(trace.SpanContext)
}

// Fields returns the fields that this propagator modifies.
func (Binary) Fields() []string { _ = "STUB: not implemented"; return nil }
