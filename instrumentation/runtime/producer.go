// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package runtime // import "go.opentelemetry.io/contrib/instrumentation/runtime"

import (
	"context"
	"runtime/metrics"
	"sync"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

var histogramMetrics = []string{goSchedLatencies}

// Producer is a metric.Producer, which provides precomputed histogram metrics from the go runtime.
type Producer struct {
	lock      sync.Mutex
	collector *goCollector
}

var _ metric.Producer = (*Producer)(nil)

// NewProducer creates a Producer which provides precomputed histogram metrics from the go runtime.
//
// Metrics emitted by NewProducer include:
//
//	go.schedule.duration    s             The time goroutines have spent in the scheduler in a runnable state before actually running.
func NewProducer(opts ...ProducerOption) *Producer { _ = "STUB: not implemented"; return nil }

// Produce returns precomputed histogram metrics from the go runtime, or an error if unsuccessful.
func (p *Producer) Produce(context.Context) ([]metricdata.ScopeMetrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use the last collection time (which may or may not be now) for the timestamp.

var emptySet = attribute.EmptySet()

func convertRuntimeHistogram(runtimeHist *metrics.Float64Histogram, ts time.Time) []metricdata.HistogramDataPoint[float64] {
	_ = "STUB: not implemented"
	return nil
}

// runtime histograms are guaranteed to have at least two bucket boundaries.

// trim the first bucket since it is a lower bound. OTel histogram boundaries only have an upper bound.

// trim the last bucket if it is +Inf, since the +Inf boundary is implicit in OTel.

// if the last bucket is not +Inf, append an extra zero count since
// the implicit +Inf bucket won't have any observations.

// This computed sum is an underestimate, since it assumes each
// observation happens at the bucket's lower bound.
