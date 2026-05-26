// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package autoexport // import "go.opentelemetry.io/contrib/exporters/autoexport"

import (
	"context"

	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/trace"
)

// noopSpanExporter is an implementation of trace.SpanExporter that performs no operations.
type noopSpanExporter struct{}

var _ trace.SpanExporter = noopSpanExporter{}

// ExportSpans is part of trace.SpanExporter interface.
func (noopSpanExporter) ExportSpans(context.Context, []trace.ReadOnlySpan) error {
	_ = "STUB: not implemented"

	// Shutdown is part of trace.SpanExporter interface.
	return nil
}

func (noopSpanExporter) Shutdown(context.Context) error {
	_ = "STUB: not implemented"

	// IsNoneSpanExporter returns true for the exporter returned by [NewSpanExporter]
	// when OTEL_TRACES_EXPORTER environment variable is set to "none".
	return nil
}

func IsNoneSpanExporter(e trace.SpanExporter) bool { _ = "STUB: not implemented"; return false }

type noopMetricReader struct {
	*metric.ManualReader
}

func newNoopMetricReader() noopMetricReader {
	_ = "STUB: not implemented"
	return *new(noopMetricReader)
}

// IsNoneMetricReader returns true for the exporter returned by [NewMetricReader]
// when OTEL_METRICS_EXPORTER environment variable is set to "none".
func IsNoneMetricReader(e metric.Reader) bool { _ = "STUB: not implemented"; return false }

type noopMetricProducer struct{}

func (noopMetricProducer) Produce(context.Context) ([]metricdata.ScopeMetrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newNoopMetricProducer() noopMetricProducer {
	_ = "STUB: not implemented"
	return *new(noopMetricProducer)
}

// noopLogExporter is an implementation of log.SpanExporter that performs no operations.
type noopLogExporter struct{}

var _ log.Exporter = noopLogExporter{}

// ExportSpans is part of log.Exporter interface.
func (noopLogExporter) Export(context.Context, []log.Record) error {
	_ = "STUB: not implemented"

	// Shutdown is part of log.Exporter interface.
	return nil
}

func (noopLogExporter) Shutdown(context.Context) error {
	_ = "STUB: not implemented"

	// ForceFlush is part of log.Exporter interface.
	return nil
}

func (noopLogExporter) ForceFlush(context.Context) error {
	_ = "STUB: not implemented"

	// IsNoneLogExporter returns true for the exporter returned by [NewLogExporter]
	// when OTEL_LOGSS_EXPORTER environment variable is set to "none".
	return nil
}

func IsNoneLogExporter(e log.Exporter) bool { _ = "STUB: not implemented"; return false }
