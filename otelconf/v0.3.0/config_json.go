// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelconf // import "go.opentelemetry.io/contrib/otelconf/v0.3.0"

// MarshalJSON implements json.Marshaler.
func (j *AttributeNameValueType) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var enumValuesAttributeNameValueType = []any{
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

// UnmarshalJSON implements json.Unmarshaler.
func (j *AttributeNameValueType) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *BatchLogRecordProcessor) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *BatchSpanProcessor) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalJSON implements json.Unmarshaler.
func (j *GeneralInstrumentationPeerServiceMappingElem) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NameStringValuePair) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

var enumValuesOTLPMetricDefaultHistogramAggregation = []any{
	"explicit_bucket_histogram",
	"base2_exponential_bucket_histogram",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *OTLPMetricDefaultHistogramAggregation) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *OTLPMetric) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalJSON implements json.Unmarshaler.
func (j *OTLP) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalJSON implements json.Unmarshaler.
func (j *OpenTelemetryConfiguration) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PeriodicMetricReader) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalJSON implements json.Unmarshaler.
func (j *PullMetricReader) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalJSON implements json.Unmarshaler.
func (j *SimpleLogRecordProcessor) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SimpleSpanProcessor) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

var enumValuesViewSelectorInstrumentType = []any{
	"counter",
	"histogram",
	"observable_counter",
	"observable_gauge",
	"observable_up_down_counter",
	"up_down_counter",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ViewSelectorInstrumentType) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Zipkin) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalJSON implements json.Unmarshaler.
func (j *AttributeNameValue) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }
