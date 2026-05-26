// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelconf // import "go.opentelemetry.io/contrib/otelconf/v0.2.0"

import (
	"context"
	"net/http"

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

func pullReader(ctx context.Context, exporter MetricExporter) (sdkmetric.Reader, error) {
	_ = "STUB: not implemented"
	return *new(sdkmetric.Reader), nil
}

func periodicExporter(ctx context.Context, exporter MetricExporter, opts ...sdkmetric.PeriodicReaderOption) (sdkmetric.Reader, error) {
	_ = "STUB: not implemented"
	return *new(sdkmetric.Reader), nil
}

func otlpHTTPMetricExporter(ctx context.Context, otlpConfig *OTLPMetric) (sdkmetric.Exporter, error) {
	_ = "STUB: not implemented"
	return *new(sdkmetric.Exporter), nil
}

func otlpGRPCMetricExporter(ctx context.Context, otlpConfig *OTLPMetric) (sdkmetric.Exporter, error) {
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

func prometheusReader(ctx context.Context, prometheusConfig *Prometheus) (sdkmetric.Reader, error) {
	_ = "STUB: not implemented"
	return *new(sdkmetric.Reader), nil
}

// NoTranslation preserves OTel dot-style label names (e.g. service.name)
// and suppresses metric name suffixes. The exporter's newConfig will
// automatically set withoutCounterSuffixes and withoutUnits when
// ShouldAddSuffixes() is false.

// Timeouts are necessary to make a server resilient to attacks, but ListenAndServe doesn't set any.
// We use values from this example: https://blog.cloudflare.com/exposing-go-on-the-internet/#:~:text=There%20are%20three%20main%20timeouts

// Remove surrounding "[]" from the host definition to allow users to define the host as "[::1]" or "::1".

// Only for testing reasons, add the address to the http Server, will not be used.

type readerWithServer struct {
	sdkmetric.Reader
	server *http.Server
}

func (rws readerWithServer) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func view(v View) (sdkmetric.View, error) {
	_ = "STUB: not implemented"
	return *new(sdkmetric.View), nil
}

func instrument(vs ViewSelector) (sdkmetric.Instrument, error) {
	_ = "STUB: not implemented"
	return *new(sdkmetric.Instrument), nil
}

func stream(vs *ViewStream) sdkmetric.Stream {
	_ = "STUB: not implemented"
	return *new(sdkmetric.Stream)
}

func attributeFilter(attributeKeys []string) attribute.Filter {
	_ = "STUB: not implemented"
	return *new(attribute.Filter)
}

func aggregation(aggr *ViewStreamAggregation) sdkmetric.Aggregation {
	_ = "STUB: not implemented"
	return *new(sdkmetric.Aggregation)
}

// Need to negate because config has the positive action RecordMinMax.

// TODO: Understand what to set here.

// Need to negate because config has the positive action RecordMinMax.

func instrumentKind(vsit *ViewSelectorInstrumentType) (sdkmetric.InstrumentKind, error) {
	_ = "STUB: not implemented"

	// Equivalent to instrumentKindUndefined.
	return *new(sdkmetric.InstrumentKind), nil
}

func instrumentIsEmpty(i sdkmetric.Instrument) bool { _ = "STUB: not implemented"; return false }

func boolOrFalse(pBool *bool) bool { _ = "STUB: not implemented"; return false }

func int32OrZero(pInt *int) int32 { _ = "STUB: not implemented"; return 0 }

func strOrEmpty(pStr *string) string { _ = "STUB: not implemented"; return "" }
