// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Based on https://github.com/DataDog/dd-trace-go/blob/8fb554ff7cf694267f9077ae35e27ce4689ed8b6/contrib/gin-gonic/gin/gintrace.go

package otelgin // import "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

import (
	"github.com/gin-gonic/gin"
)

const (
	tracerKey = "otel-go-contrib-tracer"
	// ScopeName is the instrumentation scope name.
	ScopeName = "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

// Middleware returns middleware that will trace incoming requests.
// The service parameter should describe the name of the (virtual)
// server handling the request.
func Middleware(service string, opts ...Option) gin.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(gin.HandlerFunc)
}

// Serve the request to the next middleware
// if a filter rejects the request.

// Serve the request to the next middleware
// if a filter rejects the request.

// Gin's ClientIP method can detect the client's IP from various headers set by proxies, and it's configurable

// pass the span through the request context

// serve the request to the next middleware

//nolint:forbidigo // TODO: https://github.com/open-telemetry/opentelemetry-go-contrib/issues/8441

// Record the server-side attributes.

// HTML will trace the rendering of the template as a child of the
// span in the given context. This is a replacement for
// gin.Context.HTML function - it invokes the original function after
// setting up the span.
func HTML(c *gin.Context, code int, name string, obj any) { _ = "STUB: not implemented"; return }
