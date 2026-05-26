// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Copyright 2017, OpenCensus Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package zpages // import "go.opentelemetry.io/contrib/zpages"

import (
	"context"
	"sync"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

var _ sdktrace.SpanProcessor = (*SpanProcessor)(nil)

// perMethodSummary is a summary of the spans stored for a single span name.
type perMethodSummary struct {
	activeSpans  int
	latencySpans []int
	errorSpans   int
}

// SpanProcessor is an sdktrace.SpanProcessor implementation that exposes zpages functionality for opentelemetry-go.
//
// It tracks all active spans, and stores samples of spans based on latency for non errored spans,
// and samples for errored spans.
type SpanProcessor struct {
	// Cannot keep track of the active Spans per name because the Span interface,
	// allows the name to be changed, and that will leak memory.
	activeSpansStore sync.Map
	spanSampleStores sync.Map
}

// NewSpanProcessor returns a new SpanProcessor.
func NewSpanProcessor() *SpanProcessor { _ = "STUB: not implemented"; return nil }

// OnStart adds span as active and reports it with zpages.
func (ssm *SpanProcessor) OnStart(_ context.Context, span sdktrace.ReadWriteSpan) {
	_ = "STUB: not implemented"
	return
}

// OnEnd processes all spans and reports them with zpages.
func (ssm *SpanProcessor) OnEnd(span sdktrace.ReadOnlySpan) { _ = "STUB: not implemented"; return }

// Shutdown does nothing.
func (*SpanProcessor) Shutdown(context.Context) error {
	_ = "STUB: not implemented"
	// Do nothing
	return nil
}

// ForceFlush does nothing.
func (*SpanProcessor) ForceFlush(context.Context) error {
	_ = "STUB: not implemented"
	// Do nothing
	return nil
}

// spanStoreForName returns the sampleStore for the given name.
//
// It returns nil if it doesn't exist.
func (ssm *SpanProcessor) spanStoreForName(name string) *sampleStore {
	_ = "STUB: not implemented"
	return nil
}

// spansPerMethod returns a summary of what spans are being stored for each span name.
func (ssm *SpanProcessor) spansPerMethod() map[string]*perMethodSummary {
	_ = "STUB: not implemented"
	return nil
}

// activeSpans returns the active spans for the given name.
func (ssm *SpanProcessor) activeSpans(name string) []sdktrace.ReadOnlySpan {
	_ = "STUB: not implemented"
	return nil
}

// errorSpans returns a sample of error spans.
func (ssm *SpanProcessor) errorSpans(name string) []sdktrace.ReadOnlySpan {
	_ = "STUB: not implemented"
	return nil
}

// spansByLatency returns a sample of successful spans.
//
// minLatency is the minimum latency of spans to be returned.
// maxDuration, if nonzero, is the maximum latency of spans to be returned.
func (ssm *SpanProcessor) spansByLatency(name string, latencyBucketIndex int) []sdktrace.ReadOnlySpan {
	_ = "STUB: not implemented"
	return nil
}

// sampleStore stores a sampled of spans for a particular span name.
//
// It contains sample of spans for error requests (status code is codes.Error);
// and a sample of spans for successful requests, bucketed by latency.
type sampleStore struct {
	sync.Mutex // protects everything below.
	latency    []*bucket
	errors     *bucket
}

// newSampleStore creates a sampleStore.
func newSampleStore(latencyBucketSize, errorBucketSize uint) *sampleStore {
	_ = "STUB: not implemented"
	return nil
}

func (ss *sampleStore) perMethodSummary() *perMethodSummary { _ = "STUB: not implemented"; return nil }

func (ss *sampleStore) spansByLatency(latencyBucketIndex int) []sdktrace.ReadOnlySpan {
	_ = "STUB: not implemented"
	return nil
}

func (ss *sampleStore) errorSpans() []sdktrace.ReadOnlySpan { _ = "STUB: not implemented"; return nil }

// sampleSpan removes adds to the corresponding latency or error bucket.
func (ss *sampleStore) sampleSpan(span sdktrace.ReadOnlySpan) { _ = "STUB: not implemented"; return }

// In case of time skew or wrong time, sample as 0 latency.

func spanKey(sc trace.SpanContext) [24]byte { _ = "STUB: not implemented"; return nil }
