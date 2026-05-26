// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package b3 // import "go.opentelemetry.io/contrib/propagators/b3"

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

const (
	// Default B3 Header names.
	b3ContextHeader      = "b3"
	b3DebugFlagHeader    = "x-b3-flags"
	b3TraceIDHeader      = "x-b3-traceid"
	b3SpanIDHeader       = "x-b3-spanid"
	b3SampledHeader      = "x-b3-sampled"
	b3ParentSpanIDHeader = "x-b3-parentspanid"

	b3TraceIDPadding = "0000000000000000"

	// B3 Single Header encoding widths.
	separatorWidth      = 1       // Single "-" character.
	samplingWidth       = 1       // Single hex character.
	traceID64BitsWidth  = 64 / 4  // 16 hex character Trace ID.
	traceID128BitsWidth = 128 / 4 // 32 hex character Trace ID.
	spanIDWidth         = 16      // 16 hex character ID.
	parentSpanIDWidth   = 16      // 16 hex character ID.
)

var (
	empty = trace.SpanContext{}

	errInvalidSampledByte        = errors.New("invalid B3 Sampled found")
	errInvalidSampledHeader      = errors.New("invalid B3 Sampled header found")
	errInvalidTraceIDHeader      = errors.New("invalid B3 traceID header found")
	errInvalidSpanIDHeader       = errors.New("invalid B3 spanID header found")
	errInvalidParentSpanIDHeader = errors.New("invalid B3 ParentSpanID header found")
	errInvalidScope              = errors.New("require either both traceID and spanID or none")
	errInvalidScopeParent        = errors.New("traceID and spanID required for ParentSpanID")
	errInvalidScopeParentSingle  = errors.New("traceID, spanID and Sampled required for ParentSpanID")
	errEmptyContext              = errors.New("empty request context")
	errInvalidTraceIDValue       = errors.New("invalid B3 traceID value found")
	errInvalidSpanIDValue        = errors.New("invalid B3 spanID value found")
	errInvalidParentSpanIDValue  = errors.New("invalid B3 ParentSpanID value found")
)

type propagator struct {
	cfg config
}

var _ propagation.TextMapPropagator = propagator{}

// New creates a B3 implementation of propagation.TextMapPropagator.
// B3 propagator serializes SpanContext to/from B3 Headers.
// This propagator supports both versions of B3 headers,
//  1. Single Header:
//     b3: {TraceId}-{SpanId}-{SamplingState}-{ParentSpanId}
//  2. Multiple Headers:
//     x-b3-traceid: {TraceId}
//     x-b3-parentspanid: {ParentSpanId}
//     x-b3-spanid: {SpanId}
//     x-b3-sampled: {SamplingState}
//     x-b3-flags: {DebugFlag}
//
// The Single Header propagator is used by default.
func New(opts ...Option) propagation.TextMapPropagator {
	_ = "STUB: not implemented"
	return *new(propagation.TextMapPropagator)
}

// Inject injects a context into the carrier as B3 headers.
// The parent span ID is omitted because it is not tracked in the
// SpanContext.
func (b3 propagator) Inject(ctx context.Context, carrier propagation.TextMapCarrier) {
	_ = "STUB: not implemented"
	return
}

// Since Debug implies deferred, don't also send "X-B3-Sampled".

// Extract extracts a context from the carrier if it contains B3 headers.
func (propagator) Extract(ctx context.Context, carrier propagation.TextMapCarrier) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Default to Single Header if a valid value exists.

// The Single Header value was invalid, fallback to Multiple Header.

// clear the deferred flag if we don't have a valid SpanContext

func (b3 propagator) Fields() []string { _ = "STUB: not implemented"; return nil }

// extractMultiple reconstructs a SpanContext from header values based on B3
// Multiple header. It is based on the implementation found here:
// https://github.com/openzipkin/zipkin-go/blob/v0.2.2/propagation/b3/spancontext.go
// and adapted to support a SpanContext.
func extractMultiple(ctx context.Context, traceID, spanID, parentSpanID, sampled, flags string) (context.Context, trace.SpanContext, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(trace.SpanContext), nil
}

// correct values for an existing sampled header are "0" and "1".
// For legacy support and  being lenient to other tracing implementations we
// allow "true" and "false" as inputs for interop purposes.

// Zero value for TraceFlags sample bit is unset.

// The only accepted value for Flags is "1". This will set Debug bitmask and
// sampled bitmask to 1 since debug implicitly means sampled. All other
// values and omission of header will be ignored. According to the spec. User
// shouldn't send X-B3-Sampled header along with X-B3-Flags header. Thus we will
// ignore X-B3-Sampled header when X-B3-Flags header is sent and valid.

// Pad 64-bit trace IDs.

// Validate parent span ID but we do not use it so do not save it.

// extractSingle reconstructs a SpanContext from contextHeader based on a B3
// Single header. It is based on the implementation found here:
// https://github.com/openzipkin/zipkin-go/blob/v0.2.2/propagation/b3/spancontext.go
// and adapted to support a SpanContext.
func extractSingle(ctx context.Context, contextHeader string) (context.Context, trace.SpanContext, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(trace.SpanContext), nil
}

// Trace ID by itself is invalid.

// traceID must be 64 bits
// {traceID}

// traceID must be 128 bits
// {traceID}

// {traceID}-

// {traceID}-{spanID}

// {traceID}-{spanID}- is invalid.

// {traceID}-{spanID}-

// {traceID}-{spanID}-{parentSpanID} is invalid.

// {traceID}-{spanID}-{sampling}-

// Validate parent span ID but we do not use it so do not
// save it.

// Zero value for TraceFlags sample bit is unset.
