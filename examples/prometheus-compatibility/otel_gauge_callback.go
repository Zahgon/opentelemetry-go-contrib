// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// #docregion
package main

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var (
	roomLivingRoom = attribute.String("room", "living_room")
	roomBedroom    = attribute.String("room", "bedroom")
)

func otelGaugeCallbackUsage(meter metric.Meter) {
	_ = "STUB: not implemented"
	// Temperature sensors maintain their own readings in firmware.
	// Use an observable gauge to report those values when metrics are collected.
	return
}
