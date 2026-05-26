// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// #docregion
package main

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var (
	deviceThermostat = attribute.String("device_type", "thermostat")
	deviceLock       = attribute.String("device_type", "lock")
)

func otelUpDownCounterCallbackUsage(meter metric.Meter) {
	_ = "STUB: not implemented"
	// The device manager maintains the count of connected devices.
	// Use an observable up-down counter to report that value when metrics are collected.
	return
}
