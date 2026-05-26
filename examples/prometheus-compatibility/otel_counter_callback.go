// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// #docregion
package main

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var (
	zoneUpstairs   = attribute.String("zone", "upstairs")
	zoneDownstairs = attribute.String("zone", "downstairs")
)

func otelCounterCallbackUsage(meter metric.Meter) {
	_ = "STUB: not implemented"
	// Each zone has its own smart energy meter tracking cumulative joule totals.
	// Use an observable counter to report those values when metrics are collected.
	return
}
