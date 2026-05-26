// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// #docregion
package main

import "github.com/prometheus/client_golang/prometheus"

// Prometheus uses Gauge for values that can increase or decrease.
var devicesConnected = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "devices_connected",
	Help: "Number of smart home devices currently connected",
}, []string{"device_type"})

func prometheusUpDownCounterUsage(reg *prometheus.Registry) { _ = "STUB: not implemented"; return }

// Increment when a device connects, decrement when it disconnects.
