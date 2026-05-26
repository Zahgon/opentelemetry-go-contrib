// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// #docregion createExponentialExporter
package main

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
)

func createExponentialExporter(ctx context.Context) (*otlpmetrichttp.Exporter, error) {
	_ = "STUB: not implemented"
	// Configure the exporter to use exponential histograms for all histogram instruments.
	// This is the preferred approach — it applies globally without modifying instrumentation code.
	return nil, nil
}

// #enddocregion createExponentialExporter
