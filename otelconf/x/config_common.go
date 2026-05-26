// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package x // import "go.opentelemetry.io/contrib/otelconf/x"

import (
	"context"

	sdklog "go.opentelemetry.io/otel/sdk/log"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

const (
	compressionGzip = "gzip"
	compressionNone = "none"
)

var enumValuesAttributeType = []any{
	nil,
	"string",
	"bool",
	"int",
	"double",
	"string_array",
	"bool_array",
	"int_array",
	"double_array",
}

var enumValuesViewSelectorInstrumentType = []any{
	"counter",
	"gauge",
	"histogram",
	"observable_counter",
	"observable_gauge",
	"observable_up_down_counter",
	"up_down_counter",
}

var enumValuesOTLPMetricDefaultHistogramAggregation = []any{
	"explicit_bucket_histogram",
	"base2_exponential_bucket_histogram",
}

type configOptions struct {
	ctx                   context.Context
	opentelemetryConfig   OpenTelemetryConfiguration
	loggerProviderOptions []sdklog.LoggerProviderOption
	meterProviderOptions  []sdkmetric.Option
	tracerProviderOptions []sdktrace.TracerProviderOption
}

type shutdownFunc func(context.Context) error

func noopShutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

type errBound struct {
	Field string
	Bound int
	Op    string
}

func (e *errBound) Error() string { _ = "STUB: not implemented"; return "" }

func (e *errBound) Is(target error) bool { _ = "STUB: not implemented"; return false }

type errRequired struct {
	Object any
	Field  string
}

func (e *errRequired) Error() string { _ = "STUB: not implemented"; return "" }

func (e *errRequired) Is(target error) bool { _ = "STUB: not implemented"; return false }

type errUnmarshal struct {
	Object any
}

func (e *errUnmarshal) Error() string { _ = "STUB: not implemented"; return "" }

func (e *errUnmarshal) Is(target error) bool { _ = "STUB: not implemented"; return false }

// newErrGreaterOrEqualZero creates a new error indicating that the field must be greater than
// or equal to zero.
func newErrGreaterOrEqualZero(field string) error { _ = "STUB: not implemented"; return nil }

// newErrGreaterThanZero creates a new error indicating that the field must be greater
// than zero.
func newErrGreaterThanZero(field string) error { _ = "STUB: not implemented"; return nil }

// newErrRequired creates a new error indicating that the exporter field is required.
func newErrRequired(object any, field string) error { _ = "STUB: not implemented"; return nil }

// newErrUnmarshal creates a new error indicating that an error occurred during unmarshaling.
func newErrUnmarshal(object any) error { _ = "STUB: not implemented"; return nil }

type errInvalid struct {
	Identifier string
}

func (e *errInvalid) Error() string { _ = "STUB: not implemented"; return "" }

func (*errInvalid) Is(target error) bool { _ = "STUB: not implemented"; return false }

// newErrInvalid creates a new error indicating that an error occurred due to misconfiguration.
func newErrInvalid(id string) error { _ = "STUB: not implemented"; return nil }

// unmarshalSamplerTypes handles always_on and always_off sampler unmarshaling.
func unmarshalSamplerTypes(raw map[string]any, plain *Sampler) {
	_ = "STUB: not implemented"
	// always_on can be nil, must check and set here
	return
}

// always_off can be nil, must check and set here

// unmarshalMetricProducer handles opencensus metric producer unmarshaling.
func unmarshalMetricProducer(raw map[string]any, plain *MetricProducer) {
	_ = "STUB: not implemented"
	// opencensus can be nil, must check and set here
	return
}

// validatePeriodicMetricReader handles validation for PeriodicMetricReader.
func validatePeriodicMetricReader(plain *PeriodicMetricReader) error {
	_ = "STUB: not implemented"
	return nil
}

// validateBatchLogRecordProcessor handles validation for BatchLogRecordProcessor.
func validateBatchLogRecordProcessor(plain *BatchLogRecordProcessor) error {
	_ = "STUB: not implemented"
	return nil
}

// validateBatchSpanProcessor handles validation for BatchSpanProcessor.
func validateBatchSpanProcessor(plain *BatchSpanProcessor) error {
	_ = "STUB: not implemented"
	return nil
}

// validateCardinalityLimits handles validation for CardinalityLimits.
func validateCardinalityLimits(plain *CardinalityLimits) error {
	_ = "STUB: not implemented"
	return nil
}

// validateSpanLimits handles validation for SpanLimits.
func validateSpanLimits(plain *SpanLimits) error { _ = "STUB: not implemented"; return nil }

func ptr[T any](v T) *T {
	_ = "STUB: not implemented"

	// validateOTLPHTTPEncoding validates the encoding configuration.
	// The Go SDK only supports protobuf encoding for OTLP HTTP exporters.
	return nil
}

func validateOTLPHTTPEncoding(encoding *OTLPHttpEncoding) error {
	_ = "STUB: not implemented"
	return nil
}

// createHeadersConfig combines the two header config fields. Headers take precedence over headersList.
func createHeadersConfig(headers []NameStringValuePair, headersList *string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parsing follows https://github.com/open-telemetry/opentelemetry-configuration/blob/568e5080816d40d75792eb754fc96bde09654159/schema/type_descriptions.yaml#L584.

// Headers take precedence over HeadersList, so this has to be after HeadersList is processed.

// supportedInstrumentType return an error if the instrument type is not supported.
func supportedInstrumentType(in InstrumentType) error { _ = "STUB: not implemented"; return nil }

// supportedHistogramAggregation return an error if the histogram aggregation is not supported.
func supportedHistogramAggregation(in ExporterDefaultHistogramAggregation) error {
	_ = "STUB: not implemented"
	return nil
}
