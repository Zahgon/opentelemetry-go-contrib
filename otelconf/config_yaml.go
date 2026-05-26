// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelconf // import "go.opentelemetry.io/contrib/otelconf"

import (
	"go.yaml.in/yaml/v3"
)

// hasYAMLMapKey reports whether the provided mapping node contains the given
// key. It assumes the node is a mapping node and performs a linear scan of its
// key nodes.
func hasYAMLMapKey(node *yaml.Node, key string) bool { _ = "STUB: not implemented"; return false }

// UnmarshalYAML implements yaml.Unmarshaler.

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *PushMetricExporter) UnmarshalYAML(node *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// console can be nil, must check and set here

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *OpenTelemetryConfiguration) UnmarshalYAML(node *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// Configure the log level of the internal logger used by the SDK.
// If omitted, info is used.

// Configure the log level of the internal logger used by the SDK.
// If omitted, info is used.

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *SpanExporter) UnmarshalYAML(node *yaml.Node) error { _ = "STUB: not implemented"; return nil }

// console can be nil, must check and set here

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *LogRecordExporter) UnmarshalYAML(node *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// console can be nil, must check and set here

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *TextMapPropagator) UnmarshalYAML(node *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// b3 can be nil, must check and set here

// b3multi can be nil, must check and set here

// baggage can be nil, must check and set here

// jaeger can be nil, must check and set here

// ottrace can be nil, must check and set here

// tracecontext can be nil, must check and set here

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *BatchLogRecordProcessor) UnmarshalYAML(node *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *Sampler) UnmarshalYAML(node *yaml.Node) error { _ = "STUB: not implemented"; return nil }

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *MetricProducer) UnmarshalYAML(node *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *BatchSpanProcessor) UnmarshalYAML(node *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *PeriodicMetricReader) UnmarshalYAML(node *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *CardinalityLimits) UnmarshalYAML(node *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *SpanLimits) UnmarshalYAML(node *yaml.Node) error { _ = "STUB: not implemented"; return nil }

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *OTLPHttpMetricExporter) UnmarshalYAML(node *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *OTLPGrpcMetricExporter) UnmarshalYAML(node *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *OTLPHttpExporter) UnmarshalYAML(node *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *OTLPGrpcExporter) UnmarshalYAML(node *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *AttributeType) UnmarshalYAML(node *yaml.Node) error { _ = "STUB: not implemented"; return nil }

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *AttributeNameValue) UnmarshalYAML(node *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// yaml unmarshaller defaults to unmarshalling to int

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *SimpleLogRecordProcessor) UnmarshalYAML(node *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *SimpleSpanProcessor) UnmarshalYAML(node *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *NameStringValuePair) UnmarshalYAML(node *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *InstrumentType) UnmarshalYAML(node *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *ExporterDefaultHistogramAggregation) UnmarshalYAML(node *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *PullMetricReader) UnmarshalYAML(node *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}
