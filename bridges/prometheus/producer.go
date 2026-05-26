// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheus // import "go.opentelemetry.io/contrib/bridges/prometheus"

import (
	"context"
	"errors"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

const (
	scopeName    = "go.opentelemetry.io/contrib/bridges/prometheus"
	traceIDLabel = "trace_id"
	spanIDLabel  = "span_id"
)

var (
	errUnsupportedType = errors.New("unsupported metric type")
	processStartTime   = time.Now()
)

type producer struct {
	gatherers prometheus.Gatherers
}

// NewMetricProducer returns a metric.Producer that fetches metrics from
// Prometheus. This can be used to allow Prometheus instrumentation to be
// added to an OpenTelemetry export pipeline.
func NewMetricProducer(opts ...Option) metric.Producer {
	_ = "STUB: not implemented"
	return *new(metric.Producer)
}

func (p *producer) Produce(context.Context) ([]metricdata.ScopeMetrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertPrometheusMetricsInto(promMetrics []*dto.MetricFamily, now time.Time) ([]metricdata.Metrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This shouldn't ever happen

// MetricType_GAUGE_HISTOGRAM, MetricType_UNTYPED

func isExponentialHistogram(hist *dto.Histogram) bool {
	_ = "STUB: not implemented"
	// The prometheus go client ensures at least one of these is non-zero
	// so it can be distinguished from a fixed-bucket histogram.
	// https://github.com/prometheus/client_golang/blob/7ac90362b02729a65109b33d172bafb65d7dab50/prometheus/histogram.go#L818
	return false
}

func convertGauge(metrics []*dto.Metric, now time.Time) metricdata.Gauge[float64] {
	_ = "STUB: not implemented"
	return nil
}

func convertCounter(metrics []*dto.Metric, now time.Time) metricdata.Sum[float64] {
	_ = "STUB: not implemented"
	return nil
}

func convertExponentialHistogram(metrics []*dto.Metric, now time.Time) metricdata.ExponentialHistogram[float64] {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Support exemplars

func convertExponentialBuckets(bucketSpans []*dto.BucketSpan, deltas []int64) metricdata.ExponentialBucket {
	_ = "STUB: not implemented"
	return *new(metricdata.ExponentialBucket)
}

// Prometheus Native Histograms buckets are indexed by upper boundary
// while Exponential Histograms are indexed by lower boundary, the result
// being that the Offset fields are different-by-one.

// We will have one bucket count for each delta, and zeros for the offsets
// after the initial offset.

// Do not insert zeroes if this is the first bucketSpan, since those
// zeroes are accounted for in the Offset field.

// Increase the count index by the Offset to insert Offset zeroes

// Convert deltas to the cumulative number of observations

// count should always be positive after accounting for deltas

func convertHistogram(metrics []*dto.Metric, now time.Time) metricdata.Histogram[float64] {
	_ = "STUB: not implemented"
	return nil
}

func convertBuckets(buckets []*dto.Bucket, sampleCount uint64) ([]float64, []uint64, []metricdata.Exemplar[float64]) {
	_ = "STUB: not implemented"
	return nil,

		// This should never happen
		nil, nil
}

// buckets will only include the +Inf bucket if there is an exemplar for it
// https://github.com/prometheus/client_golang/blob/d038ab96c0c7b9cd217a39072febd610bcdf1fd8/prometheus/metric.go#L189
// we need to handle the case where it is present, or where it is missing.

// The last bound may be the +Inf bucket, which is implied in OTel, but
// is explicit in Prometheus. Skip the last boundary if it is the +Inf
// bound.

// The Inf bucket was missing, so set the last bucket counts to the
// overall count

func convertSummary(metrics []*dto.Metric, now time.Time) metricdata.Summary {
	_ = "STUB: not implemented"
	return *new(metricdata.Summary)
}

func convertQuantiles(quantiles []*dto.Quantile) []metricdata.QuantileValue {
	_ = "STUB: not implemented"
	return nil
}

func convertLabels(labels []*dto.LabelPair) attribute.Set {
	_ = "STUB: not implemented"
	return *new(attribute.Set)
}

func convertExemplar(exemplar *dto.Exemplar) metricdata.Exemplar[float64] {
	_ = "STUB: not implemented"
	return nil
}

// find the trace ID and span ID in attributes, if it exists

type multierr []error

func (e multierr) errOrNil() error { _ = "STUB: not implemented"; return nil }

func (e multierr) Error() string { _ = "STUB: not implemented"; return "" }
