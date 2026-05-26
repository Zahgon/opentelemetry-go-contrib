// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// #docregion
package main

import "github.com/prometheus/client_golang/prometheus"

type deviceCountCollector struct{ desc *prometheus.Desc }

func newDeviceCountCollector() *deviceCountCollector { _ = "STUB: not implemented"; return nil }

func (c *deviceCountCollector) Describe(ch chan<- *prometheus.Desc) {
	_ = "STUB: not implemented"
	return
}
func (c *deviceCountCollector) Collect(ch chan<- prometheus.Metric) {
	_ = "STUB: not implemented"
	return
}

func prometheusUpDownCounterCallbackUsage(reg *prometheus.Registry) {
	_ = "STUB: not implemented"
	// The device manager maintains the count of connected devices.
	// Implement prometheus.Collector to report those values at scrape time.
	return
}
