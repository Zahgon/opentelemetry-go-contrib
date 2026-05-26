// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package consistent // import "go.opentelemetry.io/contrib/samplers/probability/consistent"

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	traceStateKey       = "ot"
	pValueSubkey        = "p"
	rValueSubkey        = "r"
	pZeroValue          = 63
	invalidValue        = pZeroValue + 1 // invalid for p or r
	traceStateSizeLimit = 256
)

var (
	errTraceStateSyntax       = fmt.Errorf("otel tracestate: %w", strconv.ErrSyntax)
	errTraceStateInconsistent = errors.New("r-value and p-value are inconsistent")
)

type otelTraceState struct {
	rvalue  uint8 // valid in the interval [0, 62]
	pvalue  uint8 // valid in the interval [0, 63]
	unknown []string
}

func newTraceState() otelTraceState { _ = "STUB: not implemented"; return *new(otelTraceState) }

// out-of-range => !hasRValue()
// out-of-range => !hasPValue()

func (otts otelTraceState) serialize() string { _ = "STUB: not implemented"; return "" }

// Note: should this generate an explicit error?

func isValueByte(r byte) bool { _ = "STUB: not implemented"; return false }

func isLCAlphaNum(r byte) bool { _ = "STUB: not implemented"; return false }

func isLCAlpha(r byte) bool { _ = "STUB: not implemented"; return false }

func isUCAlpha(r byte) bool { _ = "STUB: not implemented"; return false }

func parseOTelTraceState(ts string, isSampled bool) (otelTraceState, error) {
	_ = "STUB: not implemented" //nolint:revive // ignore linter
	return *new(otelTraceState), nil
}

// Note: does the spec say how to handle duplicates?

// test for a trailing ;

// Note: set R before P, so that P won't propagate if R has an error.

// Invariant checking: unset P when the values are inconsistent.

// Note: the error ensures the parent-based
// sampler repairs the broken tracestate entry.

func parseNumber(key, input string, maximum uint8) (uint8, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//nolint:gosec // `value` is strictly less then the uint8 maximum. This cast is safe.

func parseError(key string, err error) error { _ = "STUB: not implemented"; return nil }

func (otts otelTraceState) hasRValue() bool { _ = "STUB: not implemented"; return false }

func (otts otelTraceState) hasPValue() bool { _ = "STUB: not implemented"; return false }
