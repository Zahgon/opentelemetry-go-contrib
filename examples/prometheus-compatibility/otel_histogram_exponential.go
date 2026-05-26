// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// #docregion createExponentialProvider
package main

import sdkmetric "go.opentelemetry.io/otel/sdk/metric"

func createExponentialProvider(reader sdkmetric.Reader) *sdkmetric.MeterProvider {
	_ = "STUB: not implemented"
	// Configure base2 exponential histograms for all histogram instruments via a view.
	return nil
}

// #enddocregion createExponentialProvider

// #docregion createExponentialView
func createExponentialView() sdkmetric.View {
	_ = "STUB: not implemented"
	// Use a view for per-instrument control — select a specific instrument by name
	// to use exponential histograms while keeping explicit buckets for others.
	return *new(sdkmetric.View)
}

// #enddocregion createExponentialView
