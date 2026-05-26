// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelecho // import "go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"

import (
	"github.com/labstack/echo/v4"
)

const (
	tracerKey = "otel-go-contrib-tracer-labstack-echo"
	// ScopeName is the instrumentation scope name.
	ScopeName = "go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

// Middleware returns echo middleware which will trace incoming requests.
func Middleware(serverName string, opts ...Option) echo.MiddlewareFunc {
	_ = "STUB: not implemented"
	return *new(echo.MiddlewareFunc)
}

// pass the span through the request context

// serve the request to the next middleware

// Record the server-side attributes.

func spanNameFormatter(c echo.Context) string { _ = "STUB: not implemented"; return "" }
