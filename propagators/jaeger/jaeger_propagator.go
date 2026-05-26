// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package jaeger // import "go.opentelemetry.io/contrib/propagators/jaeger"

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

const (
	jaegerHeader        = "uber-trace-id"
	separator           = ":"
	traceID128bitsWidth = 128 / 4
	spanIDWidth         = 64 / 4

	idPaddingChar = "0"

	flagsDebug      = 0x02
	flagsSampled    = 0x01
	flagsNotSampled = 0x00

	deprecatedParentSpanID = "0"
)

var (
	empty = trace.SpanContext{}

	errMalformedTraceContextVal = errors.New("header value of uber-trace-id should contain four different part separated by : ")
	errInvalidTraceIDLength     = errors.New("invalid trace id length, must be either 16 or 32")
	errMalformedTraceID         = errors.New("cannot decode trace id from header, should be a string of hex, lowercase trace id can't be all zero")
	errInvalidSpanIDLength      = errors.New("invalid span id length, must be 16")
	errMalformedSpanID          = errors.New("cannot decode span id from header, should be a string of hex, lowercase span id can't be all zero")
	errMalformedFlag            = errors.New("cannot decode flag")
)

// Jaeger propagator serializes SpanContext to/from Jaeger Headers
//
// Jaeger format:
//
// uber-trace-id: {trace-id}:{span-id}:{parent-span-id}:{flags}.
type Jaeger struct{}

var _ propagation.TextMapPropagator = &Jaeger{}

// Inject injects a context to the carrier following jaeger format.
// The parent span ID is set to an dummy parent span id as the most implementations do.
func (Jaeger) Inject(ctx context.Context, carrier propagation.TextMapCarrier) {
	_ = "STUB: not implemented"
	return
}

// Extract extracts a context from the carrier if it contains Jaeger headers.
func (Jaeger) Extract(ctx context.Context, carrier propagation.TextMapCarrier) context.Context {
	_ = "STUB: not implemented"
	// extract tracing information
	return *new(context.Context)
}

func extract(ctx context.Context, headerVal string) (context.Context, trace.SpanContext, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(trace.SpanContext), nil
}

// extract trace ID

// padding when length is less than 32

// extract span ID

// padding when length is less than 16

// skip third part as it is deprecated

// extract flag

// if sample bit is set, we check if debug bit is also set

// ignore other bit, including firehose since we don't have corresponding flag in trace context.

// Fields returns the Jaeger header key whose value is set with Inject.
func (Jaeger) Fields() []string { _ = "STUB: not implemented"; return nil }
