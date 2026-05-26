// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// #docregion
package main

import "github.com/prometheus/client_golang/prometheus"

type energyCollector struct{ desc *prometheus.Desc }

func newEnergyCollector() *energyCollector { _ = "STUB: not implemented"; return nil }

func (c *energyCollector) Describe(ch chan<- *prometheus.Desc) { _ = "STUB: not implemented"; return }
func (c *energyCollector) Collect(ch chan<- prometheus.Metric) { _ = "STUB: not implemented"; return }

func prometheusCounterCallbackUsage(reg *prometheus.Registry) {
	_ = "STUB: not implemented"
	// Each zone has its own smart energy meter tracking cumulative joule totals.
	// Implement prometheus.Collector to report those values at scrape time.
	return
}
