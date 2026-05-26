// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package deprecatedruntime // import "go.opentelemetry.io/contrib/instrumentation/runtime/internal/deprecatedruntime"

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/metric"
)

// Runtime reports the work-in-progress conventional runtime metrics specified by OpenTelemetry.
type runtime struct {
	minimumReadMemStatsInterval time.Duration
	meter                       metric.Meter
}

// Start initializes reporting of runtime metrics using the supplied config.
func Start(meter metric.Meter, minimumReadMemStatsInterval time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *runtime) register() error { _ = "STUB: not implemented"; return nil }

func (r *runtime) registerMemStats() error { _ = "STUB: not implemented"; return nil }

// TODO: is ptrLookups useful? I've not seen a value
// other than zero.

// lock prevents a race between batch observer and instrument registration.

// FYI see https://github.com/golang/go/issues/32284 to help
// understand the meaning of this value.

// Note that the following could be derived as a sum of
// individual pauses, but we may lose individual pauses if the
// observation interval is too slow.

func clampUint64(v uint64) int64 { _ = "STUB: not implemented"; return 0 }

func computeGCPauses(
	ctx context.Context,
	recorder metric.Int64Histogram,
	circular []uint64,
	lastNumGC, currentNumGC uint32,
) {
	_ = "STUB: not implemented"
	return
}

// There were > 256 collections, some may have been lost.

// Only the case in error situations.

// wrap around the circular buffer

func recordGCPauses(
	ctx context.Context,
	recorder metric.Int64Histogram,
	pauses []uint64,
) {
	_ = "STUB: not implemented"
	return
}
