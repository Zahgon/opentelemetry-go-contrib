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
	deviceThermostatOpts = []metric.RecordOption{metric.WithAttributes(attribute.String("device_type", "thermostat"))}
	deviceLockOpts       = []metric.RecordOption{metric.WithAttributes(attribute.String("device_type", "lock"))}
)

func otelHistogramUsage(ctx context.Context, meter metric.Meter) {
	_ = "STUB: not implemented"
	// WithExplicitBucketBoundaries sets default boundaries as a hint to the SDK.
	// Views configured at the SDK level take precedence over this hint.
	return
}
