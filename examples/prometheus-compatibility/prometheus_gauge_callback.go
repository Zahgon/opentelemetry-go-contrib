// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// #docregion
package main

import "github.com/prometheus/client_golang/prometheus"

type temperatureCollector struct{ desc *prometheus.Desc }

func newTemperatureCollector() *temperatureCollector { _ = "STUB: not implemented"; return nil }

func (c *temperatureCollector) Describe(ch chan<- *prometheus.Desc) {
	_ = "STUB: not implemented"
	return
}
func (c *temperatureCollector) Collect(ch chan<- prometheus.Metric) {
	_ = "STUB: not implemented"
	return
}

func prometheusGaugeCallbackUsage(reg *prometheus.Registry) {
	_ = "STUB: not implemented"
	// Temperature sensors maintain their own readings in firmware.
	// Implement prometheus.Collector to report those values at scrape time.
	return
}
