// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package xray provides an OpenTelemetry propagator for the AWS XRAY
// propagation format.
package xray // import "go.opentelemetry.io/contrib/propagators/aws/xray"

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

const (
	traceHeaderKey       = "X-Amzn-Trace-Id"
	traceHeaderDelimiter = ";"
	kvDelimiter          = "="
	traceIDKey           = "Root"
	sampleFlagKey        = "Sampled"
	parentIDKey          = "Parent"
	traceIDVersion       = "1"
	traceIDDelimiter     = "-"
	isSampled            = "1"
	notSampled           = "0"

	traceFlagNone           = 0x0
	traceFlagSampled        = 0x1 << 0
	traceIDLength           = 35
	traceIDDelimitterIndex1 = 1
	traceIDDelimitterIndex2 = 10
	traceIDFirstPartLength  = 8
	sampledFlagLength       = 1
)

var (
	empty                    = trace.SpanContext{}
	errInvalidTraceHeader    = errors.New("invalid X-Amzn-Trace-Id header value, should contain 3 different part separated by ;")
	errMalformedTraceID      = errors.New("cannot decode trace ID from header")
	errLengthTraceIDHeader   = errors.New("incorrect length of X-Ray trace ID found, 35 character length expected")
	errInvalidTraceIDVersion = errors.New("invalid X-Ray trace ID header found, does not have valid trace ID version")
	errInvalidSpanIDLength   = errors.New("invalid span ID length, must be 16")
)

// Propagator serializes Span Context to/from AWS X-Ray headers.
//
// Example AWS X-Ray format:
//
// X-Amzn-Trace-Id: Root={traceId};Parent={parentId};Sampled={samplingFlag}.
type Propagator struct{}

// Asserts that the propagator implements the otel.TextMapPropagator interface at compile time.
var _ propagation.TextMapPropagator = &Propagator{}

// Inject injects a context to the carrier following AWS X-Ray format.
func (Propagator) Inject(ctx context.Context, carrier propagation.TextMapCarrier) {
	_ = "STUB: not implemented"
	return
}

// Extract gets a context from the carrier if it contains AWS X-Ray headers.
func (Propagator) Extract(ctx context.Context, carrier propagation.TextMapCarrier) context.Context {
	_ = "STUB: not implemented"
	// extract tracing information
	return *new(context.Context)
}

// extract extracts Span Context from context.
func extract(headerVal string) (trace.SpanContext, error) {
	_ = "STUB: not implemented"
	return *new(trace.SpanContext), nil
}

// last part

// extract parentId

// extract traceflag

// indexOf returns position of the first occurrence of a substr in str starting at pos index.
func indexOf(str, substr string, pos int) int { _ = "STUB: not implemented"; return 0 }

// parseTraceID returns trace ID if  valid else return invalid trace ID.
func parseTraceID(xrayTraceID string) (trace.TraceID, error) {
	_ = "STUB: not implemented"
	return *new(trace.TraceID), nil
}

// parseTraceFlag returns a parsed trace flag.
func parseTraceFlag(xraySampledFlag string) trace.TraceFlags {
	_ = "STUB: not implemented"
	// Use a direct comparison here (#7262).
	return *new(trace.TraceFlags)
}

// Fields returns list of fields used by HTTPTextFormat.
func (Propagator) Fields() []string { _ = "STUB: not implemented"; return nil }
