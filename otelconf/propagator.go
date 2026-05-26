// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelconf // import "go.opentelemetry.io/contrib/otelconf"

import (
	"go.opentelemetry.io/otel/propagation"
)

func newPropagator(p *Propagator) (propagation.TextMapPropagator, error) {
	_ = "STUB: not implemented"
	return *new(propagation.TextMapPropagator), nil
}
