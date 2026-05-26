// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheus // import "go.opentelemetry.io/contrib/bridges/prometheus"

import (
	"github.com/prometheus/client_golang/prometheus"
)

// config contains options for the producer.
type config struct {
	gatherers []prometheus.Gatherer
}

// newConfig creates a validated config configured with options.
func newConfig(opts ...Option) config { _ = "STUB: not implemented"; return *new(config) }

// Option sets producer option values.
type Option interface {
	apply(config) config
}

type optionFunc func(config) config

func (fn optionFunc) apply(cfg config) config {
	_ = "STUB: not implemented"

	// WithGatherer configures which prometheus Gatherer the Bridge will gather
	// from. If no registerer is used the prometheus DefaultGatherer is used.
	return *new(config)
}

func WithGatherer(gatherer prometheus.Gatherer) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
