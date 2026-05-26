// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelconf // import "go.opentelemetry.io/contrib/otelconf"

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/resource"
)

var zeroScope instrumentation.Scope

const instrumentKindUndefined = sdkmetric.InstrumentKind(0)

func meterProvider(cfg configOptions, res *resource.Resource) (metric.MeterProvider, shutdownFunc, error) {
	_ = "STUB: not implemented"
	return *new(metric.MeterProvider), *new(shutdownFunc), nil
}

func metricReader(ctx context.Context, r MetricReader) (sdkmetric.Reader, error) {
	_ = "STUB: not implemented"
	return *new(sdkmetric.Reader), nil
}

func pullReader(_ context.Context, _ PullMetricExporter) (sdkmetric.Reader, error) {
	_ = "STUB: not implemented"
	return *new(sdkmetric.Reader), nil
}

func periodicExporter(ctx context.Context, exporter PushMetricExporter, opts ...sdkmetric.PeriodicReaderOption) (sdkmetric.Reader, error) {
	_ = "STUB: not implemented"
	return *new(sdkmetric.Reader), nil
}

func otlpHTTPMetricExporter(ctx context.Context, otlpConfig *OTLPHttpMetricExporter) (sdkmetric.Exporter, error) {
	_ = "STUB: not implemented"
	return *new(sdkmetric.Exporter), nil
}

func otlpGRPCMetricExporter(ctx context.Context, otlpConfig *OTLPGrpcMetricExporter) (sdkmetric.Exporter, error) {
	_ = "STUB: not implemented"
	return *new(sdkmetric.Exporter), nil
}

// ParseRequestURI leaves the Host field empty when no
// scheme is specified (i.e. localhost:4317). This check is
// here to support the case where a user may not specify a
// scheme. The code does its best effort here by using
// otlpConfig.Endpoint as-is in that case

// none requires no options

func cumulativeTemporality(sdkmetric.InstrumentKind) metricdata.Temporality {
	_ = "STUB: not implemented"
	return *new(metricdata.Temporality)
}

func deltaTemporality(ik sdkmetric.InstrumentKind) metricdata.Temporality {
	_ = "STUB: not implemented"
	return *new(metricdata.Temporality)
}

func lowMemory(ik sdkmetric.InstrumentKind) metricdata.Temporality {
	_ = "STUB: not implemented"
	return *new(metricdata.Temporality)
}

// newIncludeExcludeFilter returns a Filter that includes attributes
// in the include list and excludes attributes in the excludes list.
// It returns an error if an attribute is in both lists
//
// If IncludeExclude is empty an include-all filter is returned.
func newIncludeExcludeFilter(lists *IncludeExclude) (attribute.Filter, error) {
	_ = "STUB: not implemented"
	return *new(attribute.Filter), nil
}

// check if a value is excluded first

func view(v View) (sdkmetric.View, error) {
	_ = "STUB: not implemented"
	return *new(sdkmetric.View), nil
}

func instrument(vs ViewSelector) (sdkmetric.Instrument, error) {
	_ = "STUB: not implemented"
	return *new(sdkmetric.Instrument), nil
}

func stream(vs ViewStream) (sdkmetric.Stream, error) {
	_ = "STUB: not implemented"
	return *new(sdkmetric.Stream), nil
}

func aggregation(aggr *Aggregation) sdkmetric.Aggregation {
	_ = "STUB: not implemented"
	return *new(sdkmetric.Aggregation)
}

// Need to negate because config has the positive action RecordMinMax.

// TODO: Understand what to set here.

// Need to negate because config has the positive action RecordMinMax.

func instrumentKind(vsit *InstrumentType) (sdkmetric.InstrumentKind, error) {
	_ = "STUB: not implemented"

	// Equivalent to instrumentKindUndefined.
	return *new(sdkmetric.InstrumentKind), nil
}

func instrumentIsEmpty(i sdkmetric.Instrument) bool { _ = "STUB: not implemented"; return false }

func boolOrFalse(pBool *bool) bool { _ = "STUB: not implemented"; return false }

func int32OrZero(pInt *int) int32 { _ = "STUB: not implemented"; return 0 }

func strOrEmpty(pStr *string) string { _ = "STUB: not implemented"; return "" }
