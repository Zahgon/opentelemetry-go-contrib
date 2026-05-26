// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelconf // import "go.opentelemetry.io/contrib/otelconf"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ConsoleExporter) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// If key is present (even if empty object), ensure non-nil value.

// UnmarshalJSON implements json.Unmarshaler.
func (j *B3Propagator) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// If key is present (even if empty object), ensure non-nil value.

// UnmarshalJSON implements json.Unmarshaler.
func (j *B3MultiPropagator) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// If key is present (even if empty object), ensure non-nil value.

// UnmarshalJSON implements json.Unmarshaler.
func (j *BaggagePropagator) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// If key is present (even if empty object), ensure non-nil value.

// UnmarshalJSON implements json.Unmarshaler.
func (j *JaegerPropagator) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// If key is present (even if empty object), ensure non-nil value.

// UnmarshalJSON implements json.Unmarshaler.
func (j *OpenTracingPropagator) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// If key is present (even if empty object), ensure non-nil value.

// UnmarshalJSON implements json.Unmarshaler.
func (j *TraceContextPropagator) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// If key is present (even if empty object), ensure non-nil value.

// UnmarshalJSON implements json.Unmarshaler.
func (j *PushMetricExporter) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	// Use a shadow struct with a RawMessage field to detect key presence.
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SpanExporter) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	// Use a shadow struct with a RawMessage field to detect key presence.
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *LogRecordExporter) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	// Use a shadow struct with a RawMessage field to detect key presence.
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TextMapPropagator) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalJSON implements json.Unmarshaler.
func (j *BatchLogRecordProcessor) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Hydrate the exporter into the underlying field.

// UnmarshalJSON implements json.Unmarshaler.
func (j *BatchSpanProcessor) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// Hydrate the exporter into the underlying field.

// UnmarshalJSON implements json.Unmarshaler.
func (j *OpenTelemetryConfiguration) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Configure if the SDK is disabled or not.
// If omitted or null, false is used.

// Configure the log level of the internal logger used by the SDK.
// If omitted, info is used.

// UnmarshalJSON implements json.Unmarshaler.
func (j *PeriodicMetricReader) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// Hydrate the exporter into the underlying field.

// UnmarshalJSON implements json.Unmarshaler.
func (j *CardinalityLimits) UnmarshalJSON(value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SpanLimits) UnmarshalJSON(value []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalJSON implements json.Unmarshaler.
func (j *OTLPHttpMetricExporter) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *OTLPGrpcMetricExporter) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *OTLPHttpExporter) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalJSON implements json.Unmarshaler.
func (j *OTLPGrpcExporter) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalJSON implements json.Unmarshaler.
func (j *AttributeType) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalJSON implements json.Unmarshaler.
func (j *AttributeNameValue) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// json unmarshaller defaults to unmarshalling to float for int values

// UnmarshalJSON implements json.Unmarshaler.
func (j *SimpleLogRecordProcessor) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Hydrate the exporter into the underlying field.

// UnmarshalJSON implements json.Unmarshaler.
func (j *SimpleSpanProcessor) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// Hydrate the exporter into the underlying field.

// UnmarshalJSON implements json.Unmarshaler.
func (j *NameStringValuePair) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalJSON implements json.Unmarshaler.
func (j *InstrumentType) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalJSON implements json.Unmarshaler.
func (j *ExporterDefaultHistogramAggregation) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PullMetricReader) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// Hydrate the exporter into the underlying field.

// UnmarshalJSON implements json.Unmarshaler.
func (j *Sampler) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalJSON implements json.Unmarshaler.
func (j *MetricProducer) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }
