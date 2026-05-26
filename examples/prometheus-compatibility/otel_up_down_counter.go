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
	deviceThermostatAddOpts = []metric.AddOption{metric.WithAttributes(attribute.String("device_type", "thermostat"))}
	deviceLockAddOpts       = []metric.AddOption{metric.WithAttributes(attribute.String("device_type", "lock"))}
)

func otelUpDownCounterUsage(ctx context.Context, meter metric.Meter) {
	_ = "STUB: not implemented"
	return
}

// Add() accepts positive and negative values.
