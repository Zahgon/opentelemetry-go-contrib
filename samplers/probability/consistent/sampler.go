// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package consistent provides a consistent probability based sampler.
package consistent // import "go.opentelemetry.io/contrib/samplers/probability/consistent"

import (
	"math/rand"
	"sync"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type (
	// ProbabilityBasedOption is an option to the
	// ConssitentProbabilityBased sampler.
	ProbabilityBasedOption interface {
		apply(*consistentProbabilityBasedConfig)
	}

	consistentProbabilityBasedConfig struct {
		source rand.Source
	}

	consistentProbabilityBasedRandomSource struct {
		rand.Source
	}

	consistentProbabilityBased struct {
		// "LAC" is an abbreviation for the logarithm of
		// adjusted count.  Greater values have greater
		// representivity, therefore lesser sampling
		// probability.

		// lowLAC is the lower-probability log-adjusted count
		lowLAC uint8
		// highLAC is the higher-probability log-adjusted
		// count.  except for the zero probability special
		// case, highLAC == lowLAC - 1.
		highLAC uint8
		// lowProb is the probability that lowLAC should be used,
		// in the interval (0, 1].  For exact powers of two and the
		// special case of 0 probability, lowProb == 1.
		lowProb float64

		// lock protects rnd
		lock sync.Mutex
		rnd  *rand.Rand
	}
)

// WithRandomSource sets the source of the randomness used by the Sampler.
func WithRandomSource(source rand.Source) ProbabilityBasedOption {
	_ = "STUB: not implemented"
	return *new(ProbabilityBasedOption)
}

func (s consistentProbabilityBasedRandomSource) apply(cfg *consistentProbabilityBasedConfig) {
	_ = "STUB: not implemented"
	return

	// ProbabilityBased samples a given fraction of traces.  Based on the
	// OpenTelemetry specification, this Sampler supports only power-of-two
	// fractions.  When the input fraction is not a power of two, it will
	// be rounded down.
	// - Fractions >= 1 will always sample.
	// - Fractions < 2^-62 are treated as zero.
	//
	// This Sampler sets the OpenTelemetry tracestate p-value and/or r-value.
	//
	// To respect the parent trace's `SampledFlag`, this sampler should be
	// used as the root delegate of a `Parent` sampler.
}

func ProbabilityBased(fraction float64, opts ...ProbabilityBasedOption) sdktrace.Sampler {
	_ = "STUB: not implemented"
	return *new(sdktrace.Sampler)
}

//nolint:gosec // G404: Use of weak random number generator (math/rand instead of crypto/rand) is ignored as this is not security-sensitive.

//nolint:gosec // G404: Use of weak random number generator (math/rand instead of crypto/rand) is ignored as this is not security-sensitive.

func (cs *consistentProbabilityBased) newR() uint8 { _ = "STUB: not implemented"; return 0 }

//nolint:gosec // Int63 returns an always positive number.

func (cs *consistentProbabilityBased) lowChoice() bool { _ = "STUB: not implemented"; return false }

// ShouldSample implements "go.opentelemetry.io/otel/sdk/trace".Sampler.
func (cs *consistentProbabilityBased) ShouldSample(p sdktrace.SamplingParameters) sdktrace.SamplingResult {
	_ = "STUB: not implemented"
	return *new(sdktrace.SamplingResult)
}

// Note: this ignores whether psc.IsValid() because this
// allows other otel trace state keys to pass through even
// for root decisions.

// Note: a state.Insert(traceStateKey)
// follows, nothing else needs to be done here.

// Note: see the note in
// "go.opentelemetry.io/otel/trace".TraceState.Insert(). The
// error below is not a condition we're supposed to handle.

// Description returns "ProbabilityBased{%g}" with the configured probability.
func (cs *consistentProbabilityBased) Description() string { _ = "STUB: not implemented"; return "" }
