// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// #docregion
package main

import "github.com/prometheus/client_golang/prometheus"

var hvacOnTime = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "hvac_on_seconds_total",
	Help: "Total time the HVAC system has been running, in seconds",
}, []string{"zone"})

func prometheusCounterUsage(reg *prometheus.Registry) { _ = "STUB: not implemented"; return }

// Pre-bind to label value sets: subsequent calls avoid the series lookup.

// Pre-initialize a series so it appears in /metrics with value 0.
