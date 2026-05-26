// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package b3 // import "go.opentelemetry.io/contrib/propagators/b3"

import "context"

type b3KeyType int

const (
	debugKey b3KeyType = iota
	deferredKey
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

// withDeferred returns a copy of parent with deferred set as the deferred flag value .
func withDeferred(parent context.Context, deferred bool) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// deferredFromContext returns the deferred value stored in ctx.
//
// If no deferred value is stored in ctx false is returned.
func deferredFromContext(ctx context.Context) bool { _ = "STUB: not implemented"; return false }
