// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package consistent // import "go.opentelemetry.io/contrib/samplers/probability/consistent"

import (
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type (
	parentProbabilitySampler struct {
		delegate sdktrace.Sampler
	}
)

// ParentProbabilityBased is an implementation of the OpenTelemetry
// Trace Sampler interface that provides additional checks for tracestate
// Probability Sampling fields.
func ParentProbabilityBased(root sdktrace.Sampler, samplers ...sdktrace.ParentBasedSamplerOption) sdktrace.Sampler {
	_ = "STUB: not implemented"
	return *new(sdktrace.Sampler)
}

// ShouldSample implements "go.opentelemetry.io/otel/sdk/trace".Sampler.
func (p *parentProbabilitySampler) ShouldSample(params sdktrace.SamplingParameters) sdktrace.SamplingResult {
	_ = "STUB: not implemented"
	return *new(sdktrace.SamplingResult)
}

// Note: We do not check psc.IsValid(), i.e., we repair the tracestate
// with or without a parent TraceId and SpanId.

// Note: see the note in
// "go.opentelemetry.io/otel/trace".TraceState.Insert(). The
// error below is not a condition we're supposed to handle.

// Fix the broken tracestate before calling the delegate.

// Description returns the same description as the built-in
// ParentBased sampler, with "ParentBased" replaced by
// "ParentProbabilityBased".
func (p *parentProbabilitySampler) Description() string { _ = "STUB: not implemented"; return "" }
