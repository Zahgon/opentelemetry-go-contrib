// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Copyright (c) 2021 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
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

package jaegerremote // import "go.opentelemetry.io/contrib/samplers/jaegerremote"

import (
	"math"
	"sync"

	jaeger_api_v2 "github.com/jaegertracing/jaeger-idl/proto-gen/api_v2"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/trace"

	"go.opentelemetry.io/contrib/samplers/jaegerremote/internal/ratelimiter"
)

const (
	defaultMaxOperations = 2000
)

const (
	samplerTypeKey                = "jaeger.sampler.type"
	samplerParamKey               = "jaeger.sampler.param"
	samplerTypeValueProbabilistic = "probabilistic"
	samplerTypeValueRateLimiting  = "ratelimiting"
)

// -----------------------

// probabilisticSampler is a sampler that randomly samples a certain percentage
// of traces.
type probabilisticSampler struct {
	samplingRate       float64
	sampler            trace.Sampler
	attributes         []attribute.KeyValue
	attributesDisabled bool
}

// newProbabilisticSampler creates a sampler that randomly samples a certain percentage of traces specified by the
// samplingRate, in the range between 0.0 and 1.0. it utilizes the SDK `trace.TraceIDRatioBased` sampler.
func newProbabilisticSampler(samplingRate float64, attributesDisabled bool) *probabilisticSampler {
	_ = "STUB: not implemented"
	return nil
}

func (s *probabilisticSampler) init(samplingRate float64) *probabilisticSampler {
	s.samplingRate = math.Max(0.0, math.Min(samplingRate, 1.0))
	s.sampler = trace.TraceIDRatioBased(s.samplingRate)
	if s.attributesDisabled {
		return s
	}
	s.attributes = []attribute.KeyValue{attribute.String(samplerTypeKey, samplerTypeValueProbabilistic), attribute.Float64(samplerParamKey, s.samplingRate)}
	return s
}

// SamplingRate returns the sampling probability this sampled was constructed with.
func (s *probabilisticSampler) SamplingRate() float64 { _ = "STUB: not implemented"; return 0 }

func (s *probabilisticSampler) ShouldSample(p trace.SamplingParameters) trace.SamplingResult {
	_ = "STUB: not implemented"
	return *new(trace.SamplingResult)
}

// Equal compares with another sampler.
func (s *probabilisticSampler) Equal(other trace.Sampler) bool {
	_ = "STUB: not implemented"
	return false
}

// consider equal if within 0.000001%

// Update modifies in-place the sampling rate. Locking must be done externally.
func (s *probabilisticSampler) Update(samplingRate float64) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *probabilisticSampler) Description() string { _ = "STUB: not implemented"; return "" }

// -----------------------

// rateLimitingSampler samples at most maxTracesPerSecond. The distribution of sampled traces follows
// burstiness of the service, i.e. a service with uniformly distributed requests will have those
// requests sampled uniformly as well, but if requests are bursty, especially sub-second, then a
// number of sequential requests can be sampled each second.
type rateLimitingSampler struct {
	maxTracesPerSecond float64
	rateLimiter        *ratelimiter.RateLimiter
	attributes         []attribute.KeyValue
	attributesDisabled bool
}

// newRateLimitingSampler creates new rateLimitingSampler.
func newRateLimitingSampler(maxTracesPerSecond float64, attributesDisabled bool) *rateLimitingSampler {
	_ = "STUB: not implemented"
	return nil
}

func (s *rateLimitingSampler) init(maxTracesPerSecond float64) *rateLimitingSampler {
	if s.rateLimiter == nil {
		s.rateLimiter = ratelimiter.NewRateLimiter(maxTracesPerSecond, math.Max(maxTracesPerSecond, 1.0))
	} else {
		s.rateLimiter.Update(maxTracesPerSecond, math.Max(maxTracesPerSecond, 1.0))
	}
	s.maxTracesPerSecond = maxTracesPerSecond
	if s.attributesDisabled {
		return s
	}
	s.attributes = []attribute.KeyValue{attribute.String(samplerTypeKey, samplerTypeValueRateLimiting), attribute.Float64(samplerParamKey, s.maxTracesPerSecond)}
	return s
}

func (s *rateLimitingSampler) ShouldSample(p trace.SamplingParameters) trace.SamplingResult {
	_ = "STUB: not implemented"
	return *new(trace.SamplingResult)
}

// Update reconfigures the rate limiter, while preserving its accumulated balance.
// Locking must be done externally.
func (s *rateLimitingSampler) Update(maxTracesPerSecond float64) { _ = "STUB: not implemented"; return }

// Equal compares with another sampler.
func (s *rateLimitingSampler) Equal(other trace.Sampler) bool {
	_ = "STUB: not implemented"
	return false
}

func (*rateLimitingSampler) Description() string { _ = "STUB: not implemented"; return "" }

// -----------------------

// guaranteedThroughputProbabilisticSampler is a sampler that leverages both probabilisticSampler and
// rateLimitingSampler. The rateLimitingSampler is used as a guaranteed lower bound sampler such that
// every operation is sampled at least once in a time interval defined by the lowerBound. ie a lowerBound
// of 1.0 / (60 * 10) will sample an operation at least once every 10 minutes.
//
// The probabilisticSampler is given higher priority when tags are emitted, ie. if IsSampled() for both
// samplers return true, the tags for probabilisticSampler will be used.
type guaranteedThroughputProbabilisticSampler struct {
	probabilisticSampler *probabilisticSampler
	lowerBoundSampler    *rateLimitingSampler
	samplingRate         float64
	lowerBound           float64
	attributesDisabled   bool
}

func newGuaranteedThroughputProbabilisticSampler(lowerBound, samplingRate float64, attributesDisabled bool) *guaranteedThroughputProbabilisticSampler {
	_ = "STUB: not implemented"
	return nil
}

func (s *guaranteedThroughputProbabilisticSampler) setProbabilisticSampler(samplingRate float64) {
	_ = "STUB: not implemented"
	return
}

// since we don't validate samplingRate, sampler may have clamped it to [0, 1] interval

func (s *guaranteedThroughputProbabilisticSampler) ShouldSample(p trace.SamplingParameters) trace.SamplingResult {
	_ = "STUB: not implemented"
	return *new(trace.SamplingResult)
}

// this function should only be called while holding a Write lock.
func (s *guaranteedThroughputProbabilisticSampler) update(lowerBound, samplingRate float64) {
	_ = "STUB: not implemented"
	return
}

func (*guaranteedThroughputProbabilisticSampler) Description() string {
	_ = "STUB: not implemented"
	return ""
}

// -----------------------

// perOperationSampler is a delegating sampler that applies guaranteedThroughputProbabilisticSampler
// on a per-operation basis.
type perOperationSampler struct {
	sync.RWMutex

	samplers       map[string]*guaranteedThroughputProbabilisticSampler
	defaultSampler *probabilisticSampler
	lowerBound     float64
	maxOperations  int

	// see description in perOperationSamplerParams
	operationNameLateBinding bool
	attributesDisabled       bool
}

// perOperationSamplerParams defines parameters when creating perOperationSampler.
type perOperationSamplerParams struct {
	// Max number of operations that will be tracked. Other operations will be given default strategy.
	MaxOperations int

	// Opt-in feature for applications that require late binding of span name via explicit call to SetOperationName.
	// When this feature is enabled, the sampler will return retryable=true from OnCreateSpan(), thus leaving
	// the sampling decision as non-final (and the span as writeable). This may lead to degraded performance
	// in applications that always provide the correct span name on oteltrace creation.
	//
	// For backwards compatibility this option is off by default.
	OperationNameLateBinding bool

	// Initial configuration of the sampling strategies (usually retrieved from the backend by Remote Sampler).
	Strategies *jaeger_api_v2.PerOperationSamplingStrategies
}

// newPerOperationSampler returns a new perOperationSampler.
func newPerOperationSampler(params perOperationSamplerParams, attributesDisabled bool) *perOperationSampler {
	_ = "STUB: not implemented"
	return nil
}

func (s *perOperationSampler) ShouldSample(p trace.SamplingParameters) trace.SamplingResult {
	_ = "STUB: not implemented"
	return *new(trace.SamplingResult)
}

func (s *perOperationSampler) getSamplerForOperation(operation string) trace.Sampler {
	_ = "STUB: not implemented"
	return *new(trace.Sampler)
}

// Check if sampler has already been created

// Store only up to maxOperations of unique ops.

func (*perOperationSampler) Description() string { _ = "STUB: not implemented"; return "" }

func (s *perOperationSampler) update(strategies *jaeger_api_v2.PerOperationSamplingStrategies) {
	_ = "STUB: not implemented"
	return
}
