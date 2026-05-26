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
	zoneUpstairsGaugeOpts   = []metric.RecordOption{metric.WithAttributes(attribute.String("zone", "upstairs"))}
	zoneDownstairsGaugeOpts = []metric.RecordOption{metric.WithAttributes(attribute.String("zone", "downstairs"))}
)

func otelGaugeUsage(ctx context.Context, meter metric.Meter) { _ = "STUB: not implemented"; return }
