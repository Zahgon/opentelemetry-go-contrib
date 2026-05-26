// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelconf // import "go.opentelemetry.io/contrib/otelconf/v0.2.0"

import (
	"context"

	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

func tracerProvider(cfg configOptions, res *resource.Resource) (trace.TracerProvider, shutdownFunc, error) {
	_ = "STUB: not implemented"
	return *new(trace.TracerProvider), *new(shutdownFunc), nil
}

func spanExporter(ctx context.Context, exporter SpanExporter) (sdktrace.SpanExporter, error) {
	_ = "STUB: not implemented"
	return *new(sdktrace.SpanExporter), nil
}

func spanProcessor(ctx context.Context, processor SpanProcessor) (sdktrace.SpanProcessor, error) {
	_ = "STUB: not implemented"
	return *new(sdktrace.SpanProcessor), nil
}

func otlpGRPCSpanExporter(ctx context.Context, otlpConfig *OTLP) (sdktrace.SpanExporter, error) {
	_ = "STUB: not implemented"
	return *new(sdktrace.SpanExporter), nil
}

// ParseRequestURI leaves the Host field empty when no
// scheme is specified (i.e. localhost:4317). This check is
// here to support the case where a user may not specify a
// scheme. The code does its best effort here by using
// otlpConfig.Endpoint as-is in that case.

// none requires no options

func otlpHTTPSpanExporter(ctx context.Context, otlpConfig *OTLP) (sdktrace.SpanExporter, error) {
	_ = "STUB: not implemented"
	return *new(sdktrace.SpanExporter), nil
}

func batchSpanProcessor(bsp *BatchSpanProcessor, exp sdktrace.SpanExporter) (sdktrace.SpanProcessor, error) {
	_ = "STUB: not implemented"
	return *new(sdktrace.SpanProcessor), nil
}
