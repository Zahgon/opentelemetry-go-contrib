// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package xray // import "go.opentelemetry.io/contrib/propagators/aws/xray"

import (
	"context"
	"math/rand"
	"sync"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

// IDGenerator is used for generating a new traceID and spanID.
type IDGenerator struct {
	sync.Mutex
	randSource *rand.Rand
}

var _ sdktrace.IDGenerator = &IDGenerator{}

// NewSpanID returns a non-zero span ID from a randomly-chosen sequence.
func (gen *IDGenerator) NewSpanID(context.Context, trace.TraceID) trace.SpanID {
	_ = "STUB: not implemented"
	return *new(trace.SpanID)
}

// NewIDs returns a non-zero trace ID and a non-zero span ID.
// trace ID returned is based on AWS X-Ray TraceID format.
//   - https://docs.aws.amazon.com/xray/latest/devguide/xray-api-sendingdata.html#xray-api-traceids
//
// span ID is from a randomly-chosen sequence.
func (gen *IDGenerator) NewIDs(context.Context) (trace.TraceID, trace.SpanID) {
	_ = "STUB: not implemented"
	return *new(trace.TraceID), *new(trace.SpanID)
}

// NewIDGenerator returns an IDGenerator reference used for sending traces to AWS X-Ray.
func NewIDGenerator() *IDGenerator { _ = "STUB: not implemented"; return nil }

//nolint:gosec // G404: Use of weak random number generator (math/rand instead of crypto/rand) is ignored as this is not security-sensitive.

func getCurrentTimeHex() []uint8 { _ = "STUB: not implemented"; return nil }

// Ignore error since no expected error should result from this operation
// Odd-length strings and non-hex digits are the only 2 error conditions for hex.DecodeString()
// strconv.FromatInt() do not produce odd-length strings or non-hex digits
