// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ot // import "go.opentelemetry.io/contrib/propagators/ot"

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

const (
	// Default OT Header names.
	traceIDHeader       = "ot-tracer-traceid"
	spanIDHeader        = "ot-tracer-spanid"
	sampledHeader       = "ot-tracer-sampled"
	baggageHeaderPrefix = "ot-baggage-"

	otTraceIDPadding = "0000000000000000"

	traceID64BitsWidth = 64 / 4 // 16 hex character Trace ID.
)

var (
	empty = trace.SpanContext{}

	errInvalidSampledHeader = errors.New("invalid OT Sampled header found")
	errInvalidTraceIDHeader = errors.New("invalid OT traceID header found")
	errInvalidSpanIDHeader  = errors.New("invalid OT spanID header found")
	errInvalidScope         = errors.New("require either both traceID and spanID or none")
)

// OT propagator serializes SpanContext to/from ot-trace-* headers.
type OT struct{}

var _ propagation.TextMapPropagator = OT{}

// Inject injects a context into the carrier as OT headers.
// NOTE: In order to interop with systems that use the OT header format, trace ids MUST be 64-bits.
func (OT) Inject(ctx context.Context, carrier propagation.TextMapCarrier) {
	_ = "STUB: not implemented"
	return
}

// don't bother injecting anything if either trace or span IDs are not valid

// Extract extracts a context from the carrier if it contains OT headers.
func (OT) Extract(ctx context.Context, carrier propagation.TextMapCarrier) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Fields returns the OT header keys whose values are set with Inject.
func (OT) Fields() []string { _ = "STUB: not implemented"; return nil }

// extractBags extracts OpenTracing baggage information from carrier.
func extractBags(carrier propagation.TextMapCarrier) (baggage.Baggage, error) {
	_ = "STUB: not implemented"
	return *new(baggage.Baggage), nil
}

// extract reconstructs a SpanContext from header values based on OT
// headers.
func extract(traceID, spanID, sampled string) (trace.SpanContext, error) {
	_ = "STUB: not implemented"
	return *new(trace.SpanContext), nil
}

// Zero value for TraceFlags sample bit is unset.

// Zero value for TraceFlags sample bit is unset.

// Pad 64-bit trace IDs.
