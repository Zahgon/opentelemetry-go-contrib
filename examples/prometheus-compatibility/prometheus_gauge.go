// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// #docregion
package main

import "github.com/prometheus/client_golang/prometheus"

var thermostatSetpoint = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "thermostat_setpoint_celsius",
	Help: "Target temperature set on the thermostat",
}, []string{"zone"})

func prometheusGaugeUsage(reg *prometheus.Registry) { _ = "STUB: not implemented"; return }
