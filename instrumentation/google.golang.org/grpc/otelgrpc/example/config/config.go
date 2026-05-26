// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package config provides initialization of OpenTelemetry setup.
package config // import "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc/example/config"

import (
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// Init configures an OpenTelemetry exporter and trace provider.
func Init() (*sdktrace.TracerProvider, error) { _ = "STUB: not implemented"; return nil, nil }
