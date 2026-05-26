// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package x // import "go.opentelemetry.io/contrib/otelconf/x"

import (
	"go.opentelemetry.io/otel/propagation"
)

func newPropagator(p *Propagator) (propagation.TextMapPropagator, error) {
	_ = "STUB: not implemented"
	return *new(propagation.TextMapPropagator), nil
}
