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
	summaryThermostatOpts = []metric.RecordOption{metric.WithAttributes(attribute.String("device_type", "thermostat"))}
	summaryLockOpts       = []metric.RecordOption{metric.WithAttributes(attribute.String("device_type", "lock"))}
)

func summaryReplacement(ctx context.Context, meter metric.Meter) {
	_ = "STUB: not implemented"
	// No explicit bucket boundaries: captures count and sum only.
	// For quantile estimation, prefer a base2 exponential histogram instead.
	return
}

// no boundaries
