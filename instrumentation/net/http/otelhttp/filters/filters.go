// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package filters provides a set of filters useful with the
// otelhttp.WithFilter() option to control which inbound requests are traced.
package filters // import "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp/filters"

import (
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// Any takes a list of Filters and returns a Filter that
// returns true if any Filter in the list returns true.
func Any(fs ...otelhttp.Filter) otelhttp.Filter {
	_ = "STUB: not implemented"
	return *new(otelhttp.Filter)
}

// All takes a list of Filters and returns a Filter that
// returns true only if all Filters in the list return true.
func All(fs ...otelhttp.Filter) otelhttp.Filter {
	_ = "STUB: not implemented"
	return *new(otelhttp.Filter)
}

// None takes a list of Filters and returns a Filter that returns
// true only if none of the Filters in the list return true.
func None(fs ...otelhttp.Filter) otelhttp.Filter {
	_ = "STUB: not implemented"
	return *new(otelhttp.Filter)
}

// Not provides a convenience mechanism for inverting a Filter.
func Not(f otelhttp.Filter) otelhttp.Filter {
	_ = "STUB: not implemented"
	return *new(otelhttp.Filter)
}

// Hostname returns a Filter that returns true if the request's
// hostname matches the provided string.
func Hostname(h string) otelhttp.Filter { _ = "STUB: not implemented"; return *new(otelhttp.Filter) }

// Path returns a Filter that returns true if the request's
// path matches the provided string.
func Path(p string) otelhttp.Filter { _ = "STUB: not implemented"; return *new(otelhttp.Filter) }

// PathPrefix returns a Filter that returns true if the request's
// path starts with the provided string.
func PathPrefix(p string) otelhttp.Filter { _ = "STUB: not implemented"; return *new(otelhttp.Filter) }

// Query returns a Filter that returns true if the request
// includes a query parameter k with a value equal to v.
func Query(k, v string) otelhttp.Filter { _ = "STUB: not implemented"; return *new(otelhttp.Filter) }

// QueryContains returns a Filter that returns true if the request
// includes a query parameter k with a value that contains v.
func QueryContains(k, v string) otelhttp.Filter {
	_ = "STUB: not implemented"
	return *new(otelhttp.Filter)
}

// Method returns a Filter that returns true if the request
// method is equal to the provided value.
func Method(m string) otelhttp.Filter { _ = "STUB: not implemented"; return *new(otelhttp.Filter) }
