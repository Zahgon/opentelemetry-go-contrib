// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package runtime // import "go.opentelemetry.io/contrib/instrumentation/runtime"

import (
	"time"

	"go.opentelemetry.io/otel/metric"
)

// config contains optional settings for reporting runtime metrics.
type config struct {
	// MinimumReadMemStatsInterval sets the minimum interval
	// between calls to runtime.ReadMemStats().  Negative values
	// are ignored.
	MinimumReadMemStatsInterval time.Duration

	// MeterProvider sets the metric.MeterProvider.  If nil, the global
	// Provider will be used.
	MeterProvider metric.MeterProvider
}

// Option supports configuring optional settings for runtime metrics.
type Option interface {
	apply(*config)
}

// ProducerOption supports configuring optional settings for runtime metrics using a
// metric producer in addition to standard instrumentation.
type ProducerOption interface {
	Option
	applyProducer(*config)
}

// DefaultMinimumReadMemStatsInterval is the default minimum interval
// between calls to runtime.ReadMemStats().  Use the
// WithMinimumReadMemStatsInterval() option to modify this setting in
// Start().
const DefaultMinimumReadMemStatsInterval time.Duration = 15 * time.Second

// WithMinimumReadMemStatsInterval sets a minimum interval between calls to
// runtime.ReadMemStats(), which is a relatively expensive call to make
// frequently.  This setting is ignored when `d` is negative.
func WithMinimumReadMemStatsInterval(d time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type minimumReadMemStatsIntervalOption time.Duration

func (o minimumReadMemStatsIntervalOption) apply(c *config) { _ = "STUB: not implemented"; return }

func (o minimumReadMemStatsIntervalOption) applyProducer(c *config) {
	_ = "STUB: not implemented"

	// WithMeterProvider sets the Metric implementation to use for
	// reporting.  If this option is not used, the global metric.MeterProvider
	// will be used.  `provider` must be non-nil.
	return
}

func WithMeterProvider(provider metric.MeterProvider) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type metricProviderOption struct{ metric.MeterProvider }

func (o metricProviderOption) apply(c *config) { _ = "STUB: not implemented"; return }

// newConfig computes a config from the supplied Options.
func newConfig(opts ...Option) config { _ = "STUB: not implemented"; return *new(config) }

// newConfig computes a config from the supplied ProducerOptions.
func newProducerConfig(opts ...ProducerOption) config {
	_ = "STUB: not implemented"
	return *new(config)
}
