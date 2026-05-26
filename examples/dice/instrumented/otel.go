// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"

	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
)

// setupOTelSDK bootstraps the OpenTelemetry pipeline.
// If it does not return an error, make sure to call shutdown for proper cleanup.
func setupOTelSDK(ctx context.Context) (func(context.Context) error, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// shutdown calls cleanup functions registered via shutdownFuncs.
// The errors from the calls are joined.
// Each registered cleanup will be invoked once.

// handleErr calls shutdown for cleanup and makes sure that all errors are returned.

// Set up propagator.

// Set up trace provider.

// Set up meter provider.

// Set up logger provider.

func newPropagator() propagation.TextMapPropagator {
	_ = "STUB: not implemented"
	return *new(propagation.TextMapPropagator)
}

func newTracerProvider() (*trace.TracerProvider, error) { _ = "STUB: not implemented"; return nil, nil }

// Default is 5s. Set to 1s for demonstrative purposes.

func newMeterProvider() (*metric.MeterProvider, error) { _ = "STUB: not implemented"; return nil, nil }

// Default is 1m. Set to 3s for demonstrative purposes.

func newLoggerProvider() (*log.LoggerProvider, error) { _ = "STUB: not implemented"; return nil, nil }
