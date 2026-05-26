// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package baggagecopy // import "go.opentelemetry.io/contrib/processors/baggagecopy"

import (
	"context"

	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Filter returns true if the baggage member should be added to a span.
type Filter func(member baggage.Member) bool

// AllowAllMembers allows all baggage members to be added to a span.
var AllowAllMembers Filter = func(baggage.Member) bool { return true }

// SpanProcessor is a [trace.SpanProcessor] implementation that adds baggage
// members onto a span as attributes.
type SpanProcessor struct {
	filter Filter
}

var _ trace.SpanProcessor = (*SpanProcessor)(nil)

// NewSpanProcessor returns a new [SpanProcessor].
//
// The Baggage span processor duplicates onto a span the attributes found
// in Baggage in the parent context at the moment the span is started.
// The passed filter determines which baggage members are added to the span.
//
// If filter is nil, all baggage members will be added.
func NewSpanProcessor(filter Filter) *SpanProcessor { _ = "STUB: not implemented"; return nil }

// OnStart is called when a span is started and adds span attributes for baggage contents.
func (processor SpanProcessor) OnStart(ctx context.Context, span trace.ReadWriteSpan) {
	_ = "STUB: not implemented"
	return
}

// OnEnd is called when span is finished and is a no-op for this processor.
func (SpanProcessor) OnEnd(trace.ReadOnlySpan) {
	_ = "STUB: not implemented"

	// Shutdown is called when the SDK shuts down and is a no-op for this processor.
	return
}

func (SpanProcessor) Shutdown(context.Context) error {
	_ = "STUB: not implemented"

	// ForceFlush exports all ended spans to the configured Exporter that have not yet
	// been exported and is a no-op for this processor.
	return nil
}

func (SpanProcessor) ForceFlush(context.Context) error { _ = "STUB: not implemented"; return nil }
