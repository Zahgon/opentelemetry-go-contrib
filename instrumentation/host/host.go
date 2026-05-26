// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package host // import "go.opentelemetry.io/contrib/instrumentation/host"

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// ScopeName is the instrumentation scope name.
const ScopeName = "go.opentelemetry.io/contrib/instrumentation/host"

// Host reports the work-in-progress conventional host metrics specified by OpenTelemetry.
type host struct {
	config config
	meter  metric.Meter
}

// config contains optional settings for reporting host metrics.
type config struct {
	// MeterProvider sets the metric.MeterProvider.  If nil, the global
	// Provider will be used.
	MeterProvider metric.MeterProvider
}

// Option supports configuring optional settings for host metrics.
type Option interface {
	apply(*config)
}

// WithMeterProvider sets the Metric implementation to use for
// reporting.  If this option is not used, the global metric.MeterProvider
// will be used.  `provider` must be non-nil.
func WithMeterProvider(provider metric.MeterProvider) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type metricProviderOption struct{ metric.MeterProvider }

func (o metricProviderOption) apply(c *config) { _ = "STUB: not implemented"; return }

// Attribute sets.
var (
	// Attribute sets for CPU time measurements.

	// Deprecated: Use go.opentelemetry.io/otel/semconv instead.
	AttributeCPUTimeUser = attribute.NewSet(attribute.String("state", "user"))
	// Deprecated: Use go.opentelemetry.io/otel/semconv instead.
	AttributeCPUTimeSystem = attribute.NewSet(attribute.String("state", "system"))
	// Deprecated: Use go.opentelemetry.io/otel/semconv instead.
	AttributeCPUTimeOther = attribute.NewSet(attribute.String("state", "other"))
	// Deprecated: Use go.opentelemetry.io/otel/semconv instead.
	AttributeCPUTimeIdle = attribute.NewSet(attribute.String("state", "idle"))

	// Attribute sets used for Memory measurements.

	// Deprecated: Use go.opentelemetry.io/otel/semconv instead.
	AttributeMemoryAvailable = attribute.NewSet(attribute.String("state", "available"))
	// Deprecated: Use go.opentelemetry.io/otel/semconv instead.
	AttributeMemoryUsed = attribute.NewSet(attribute.String("state", "used"))

	// Attribute sets used for Network measurements.

	// Deprecated: Use go.opentelemetry.io/otel/semconv instead.
	AttributeNetworkTransmit = attribute.NewSet(attribute.String("direction", "transmit"))
	// Deprecated: Use go.opentelemetry.io/otel/semconv instead.
	AttributeNetworkReceive = attribute.NewSet(attribute.String("direction", "receive"))
)

// newConfig computes a config from a list of Options.
func newConfig(opts ...Option) config { _ = "STUB: not implemented"; return *new(config) }

// Start initializes reporting of host metrics using the supplied config.
func Start(opts ...Option) error { _ = "STUB: not implemented"; return nil }

func (h *host) register() error { _ = "STUB: not implemented"; return nil }

// lock prevents a race between batch observer and instrument registration.

// This follows the OpenTelemetry Collector's "hostmetrics"
// receiver/hostmetricsreceiver/internal/scraper/processscraper
// measures User and System IOwait time.
// TODO: the Collector has per-OS compilation modules to support
// specific metrics that are not universal.

// TODO(#244): "other" is a placeholder for actually dealing
// with these states.  Do users actually want this
// (unconditionally)?  How should we handle "iowait"
// if not all systems expose it?  Should we break
// these down by CPU?  If so, are users going to want
// to aggregate in-process?  See:
// https://github.com/open-telemetry/opentelemetry-go-contrib/issues/244

// Host memory usage

// Host memory utilization

// Host network usage
//
// TODO: These can be broken down by network
// interface, with similar questions to those posed
// about per-CPU measurements above.

func clampInt64(v uint64) int64 { _ = "STUB: not implemented"; return 0 }
