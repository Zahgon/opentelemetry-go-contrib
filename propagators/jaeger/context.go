// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package jaeger // import "go.opentelemetry.io/contrib/propagators/jaeger"

import "context"

type jaegerKeyType int

const (
	debugKey jaegerKeyType = iota
)

// withDebug returns a copy of parent with debug set as the debug flag value .
func withDebug(parent context.Context, debug bool) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// debugFromContext returns the debug value stored in ctx.
//
// If no debug value is stored in ctx false is returned.
func debugFromContext(ctx context.Context) bool { _ = "STUB: not implemented"; return false }
