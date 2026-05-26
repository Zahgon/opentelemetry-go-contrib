// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelrestful // import "go.opentelemetry.io/contrib/instrumentation/github.com/emicklei/go-restful/otelrestful"

import (
	"github.com/emicklei/go-restful/v3"
)

// ScopeName is the instrumentation scope name.
const ScopeName = "go.opentelemetry.io/contrib/instrumentation/github.com/emicklei/go-restful/otelrestful"

// OTelFilter returns a restful.FilterFunction which will trace an incoming request.
//
// The service parameter should describe the name of the (virtual) server handling
// the request.  Options can be applied to configure the tracer and propagators
// used for this filter.
func OTelFilter(service string, opts ...Option) restful.FilterFunction {
	_ = "STUB: not implemented"
	return *new(restful.FilterFunction)
}

// Linking incoming span context if any for public endpoint.

// pass the span through the request context
