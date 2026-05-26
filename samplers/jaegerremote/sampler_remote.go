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
	"net/http"
	"sync"
	"time"

	"go.opentelemetry.io/otel/sdk/trace"
)

const (
	defaultRemoteSamplingTimeout            = 10 * time.Second
	defaultSamplingRefreshInterval          = time.Minute
	defaultSamplingMaxOperations            = 256
	defaultSamplingOperationNameLateBinding = true
)

// SamplingStrategyFetcher is used to fetch sampling strategy updates from remote server.
type SamplingStrategyFetcher interface {
	Fetch(service string) ([]byte, error)
}

// samplingStrategyParser is used to parse sampling strategy updates. The output object
// should be of the type that is recognized by the SamplerUpdaters.
type samplingStrategyParser interface {
	Parse(response []byte) (any, error)
}

// samplerUpdater is used by Sampler to apply sampling strategies,
// retrieved from remote config server, to the current sampler. The updater can modify
// the sampler in-place if sampler supports it, or create a new one.
//
// If the strategy does not contain configuration for the sampler in question,
// updater must return modifiedSampler=nil to give other updaters a chance to inspect
// the sampling strategy response.
//
// Sampler invokes the updaters while holding a lock on the main sampler.
type samplerUpdater interface {
	Update(sampler trace.Sampler, strategy any) (modified trace.Sampler, err error)
}

// Sampler is a delegating sampler that polls a remote server
// for the appropriate sampling strategy, constructs a corresponding sampler and
// delegates to it for sampling decisions.
type Sampler struct {
	// These fields must be first in the struct because `sync/atomic` expects 64-bit alignment.
	// Cf. https://github.com/jaegertracing/jaeger-client-go/issues/155, https://pkg.go.dev/sync/atomic#pkg-note-BUG
	closed int64 // 0 - not closed, 1 - closed

	sync.RWMutex // used to serialize access to samplerConfig.sampler
	config

	serviceName string
	doneChan    chan *sync.WaitGroup
}

// New creates a sampler that periodically pulls
// the sampling strategy from an HTTP sampling server (e.g. jaeger-agent).
func New(
	serviceName string,
	opts ...Option,
) *Sampler {
	_ = "STUB: not implemented"
	return nil
}

// ShouldSample returns a sampling choice based on the passed sampling
// parameters.
func (s *Sampler) ShouldSample(p trace.SamplingParameters) trace.SamplingResult {
	_ = "STUB: not implemented"
	return *new(trace.SamplingResult)
}

// Close does a clean shutdown of the sampler, stopping any background
// go-routines it may have started.
func (s *Sampler) Close() { _ = "STUB: not implemented"; return }

// Description returns a human-readable name for the Sampler.
func (*Sampler) Description() string { _ = "STUB: not implemented"; return "" }

func (s *Sampler) pollController() { _ = "STUB: not implemented"; return }

func (s *Sampler) pollControllerWithTicker(ticker *time.Ticker) { _ = "STUB: not implemented"; return }

func (s *Sampler) setSampler(sampler trace.Sampler) { _ = "STUB: not implemented"; return }

// UpdateSampler forces the sampler to fetch sampling strategy from backend server.
// This function is called automatically on a timer, but can also be safely called manually, e.g. from tests.
func (s *Sampler) UpdateSampler() { _ = "STUB: not implemented"; return }

// NB: this function should only be called while holding a Write lock.
func (s *Sampler) updateSamplerViaUpdaters(strategy any) error {
	_ = "STUB: not implemented"
	return nil
}

// -----------------------

// probabilisticSamplerUpdater is used by Sampler to parse sampling configuration.
type probabilisticSamplerUpdater struct {
	attributesDisabled bool
}

// Update implements Update of samplerUpdater.
func (u *probabilisticSamplerUpdater) Update(sampler trace.Sampler, strategy any) (trace.Sampler, error) {
	_ = "STUB: not implemented"
	return *new(trace.Sampler), nil
}

// sanity signature check

// -----------------------

// rateLimitingSamplerUpdater is used by Sampler to parse sampling configuration.
type rateLimitingSamplerUpdater struct {
	attributesDisabled bool
}

// Update implements Update of samplerUpdater.
func (u *rateLimitingSamplerUpdater) Update(sampler trace.Sampler, strategy any) (trace.Sampler, error) {
	_ = "STUB: not implemented"
	return *new(trace.Sampler), nil
}

// sanity signature check

// -----------------------

// perOperationSamplerUpdater is used by Sampler to parse sampling configuration.
// Fields have the same meaning as in perOperationSamplerParams.
type perOperationSamplerUpdater struct {
	MaxOperations            int
	OperationNameLateBinding bool
	attributesDisabled       bool
}

// Update implements Update of samplerUpdater.
func (u *perOperationSamplerUpdater) Update(sampler trace.Sampler, strategy any) (trace.Sampler, error) {
	_ = "STUB: not implemented"
	return *new(trace.Sampler), nil
}

// sanity signature check

// -----------------------

type httpSamplingStrategyFetcher struct {
	serverURL  string
	httpClient http.Client
}

func newHTTPSamplingStrategyFetcher(serverURL string) *httpSamplingStrategyFetcher {
	_ = "STUB: not implemented"
	return nil
}

func (f *httpSamplingStrategyFetcher) Fetch(serviceName string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// -----------------------

type samplingStrategyParserImpl struct{}

func (*samplingStrategyParserImpl) Parse(response []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Official Jaeger Remote Sampling protocol contains enums encoded as strings.
// Legacy protocol contains enums as numbers.
// Gogo's jsonpb module can parse either format.
// Cf. https://github.com/open-telemetry/opentelemetry-go-contrib/issues/3184
