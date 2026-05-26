// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelconf // import "go.opentelemetry.io/contrib/otelconf/v0.3.0"

import (
	"context"

	"go.opentelemetry.io/otel/log"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
)

func loggerProvider(cfg configOptions, res *resource.Resource) (log.LoggerProvider, shutdownFunc, error) {
	_ = "STUB: not implemented"
	return *new(log.LoggerProvider), *new(shutdownFunc), nil
}

func logProcessor(ctx context.Context, processor LogRecordProcessor) (sdklog.Processor, error) {
	_ = "STUB: not implemented"
	return *new(sdklog.Processor), nil
}

func logExporter(ctx context.Context, exporter LogRecordExporter) (sdklog.Exporter, error) {
	_ = "STUB: not implemented"
	return *new(sdklog.Exporter), nil
}

func batchLogProcessor(blp *BatchLogRecordProcessor, exp sdklog.Exporter) (*sdklog.BatchProcessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func otlpHTTPLogExporter(ctx context.Context, otlpConfig *OTLP) (sdklog.Exporter, error) {
	_ = "STUB: not implemented"
	return *new(sdklog.Exporter), nil
}

func otlpGRPCLogExporter(ctx context.Context, otlpConfig *OTLP) (sdklog.Exporter, error) {
	_ = "STUB: not implemented"
	return *new(sdklog.Exporter), nil
}

// ParseRequestURI leaves the Host field empty when no
// scheme is specified (i.e. localhost:4317). This check is
// here to support the case where a user may not specify a
// scheme. The code does its best effort here by using
// otlpConfig.Endpoint as-is in that case

// none requires no options
