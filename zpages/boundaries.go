// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zpages // import "go.opentelemetry.io/contrib/zpages"

import (
	"time"
)

const (
	zeroDuration = time.Duration(0)
	maxDuration  = time.Duration(1<<63 - 1)
)

var defaultBoundaries = newBoundaries([]time.Duration{
	10 * time.Microsecond,
	100 * time.Microsecond,
	time.Millisecond,
	10 * time.Millisecond,
	100 * time.Millisecond,
	time.Second,
	10 * time.Second,
	100 * time.Second,
})

// boundaries represents the interval bounds for the latency based samples.
type boundaries struct {
	durations []time.Duration
}

// newBoundaries returns a new boundaries.
func newBoundaries(durations []time.Duration) *boundaries { _ = "STUB: not implemented"; return nil }

// numBuckets returns the number of buckets needed for these boundaries.
func (lb boundaries) numBuckets() int { _ = "STUB: not implemented"; return 0 }

// getBucketIndex returns the appropriate bucket index for a given latency.
func (lb boundaries) getBucketIndex(latency time.Duration) int { _ = "STUB: not implemented"; return 0 }
