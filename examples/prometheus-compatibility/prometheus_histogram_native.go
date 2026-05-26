// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// #docregion
package main

import "github.com/prometheus/client_golang/prometheus"

var nativeDeviceCommandDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:                        "device_command_duration_seconds",
	Help:                        "Time to receive acknowledgment from a smart home device",
	NativeHistogramBucketFactor: 1.1,
}, []string{"device_type"})

func nativeHistogramUsage(reg *prometheus.Registry) { _ = "STUB: not implemented"; return }
