// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// #docregion
package main

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// Preallocate attribute options when values are static to avoid per-call allocation.
var (
	zoneUpstairsOpts   = []metric.AddOption{metric.WithAttributes(attribute.String("zone", "upstairs"))}
	zoneDownstairsOpts = []metric.AddOption{metric.WithAttributes(attribute.String("zone", "downstairs"))}
)

func otelCounterUsage(ctx context.Context, meter metric.Meter) {
	_ = "STUB: not implemented"
	// No upfront label declaration: attributes are provided at record time.
	return
}
