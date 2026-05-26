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
	"time"

	"github.com/go-logr/logr"
	"go.opentelemetry.io/otel/sdk/trace"
)

type config struct {
	sampler                 trace.Sampler
	samplingServerURL       string
	samplingRefreshInterval time.Duration
	samplingFetcher         SamplingStrategyFetcher
	samplingParser          samplingStrategyParser
	updaters                []samplerUpdater
	posParams               perOperationSamplerParams
	logger                  logr.Logger
	attributesDisabled      bool
}

func getEnvOptions() ([]Option, []error) {
	_ = "STUB: not implemented"

	// list of errors which will be logged once logger is set by the user
	return nil, nil
}

// newConfig returns an appropriately configured config.
func newConfig(options ...Option) config { _ = "STUB: not implemented"; return *new(config) }

// Option applies configuration settings to a Sampler.
type Option interface {
	apply(*config)
}

type optionFunc func(*config)

func (fn optionFunc) apply(c *config) {
	_ = "STUB: not implemented"

	// WithInitialSampler creates a Option that sets the initial sampler
	// to use before a remote sampler is created and used.
	return
}

func WithInitialSampler(sampler trace.Sampler) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSamplingServerURL creates a Option that sets the sampling server url
// of the local agent that contains the sampling strategies.
func WithSamplingServerURL(samplingServerURL string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// The default port of jaeger agent is 5778, but there are other ports specified by the user, so the sampling address and fetch address are strongly bound

// WithMaxOperations creates a Option that sets the maximum number of
// operations the sampler will keep track of.
func WithMaxOperations(maxOperations int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithOperationNameLateBinding creates a Option that sets the respective
// field in the perOperationSamplerParams.
func WithOperationNameLateBinding(enable bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSamplingRefreshInterval creates a Option that sets how often the
// sampler will poll local agent for the appropriate sampling strategy.
func WithSamplingRefreshInterval(samplingRefreshInterval time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithLogger configures the sampler to log operation and debug information with logger.
func WithLogger(logger logr.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSamplingStrategyFetcher creates an Option that initializes the sampling strategy fetcher.
// Custom fetcher can be used for setting custom headers, timeouts, etc., or getting
// sampling strategies from a different source, like files.
func WithSamplingStrategyFetcher(fetcher SamplingStrategyFetcher) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithAttributesDisabled configures the sampler to disable setting attributes jaeger.sampler.type and jaeger.sampler.param.
func WithAttributesDisabled() Option { _ = "STUB: not implemented"; return *new(Option) }

// samplingStrategyParser creates a Option that initializes sampling strategy parser.
func withSamplingStrategyParser(parser samplingStrategyParser) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
