// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package otelaws provides instrumentation for the AWS SDK.
package otelaws // import "go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws"

import (
	"context"

	"github.com/aws/smithy-go/middleware"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

const (
	// ScopeName is the instrumentation scope name.
	ScopeName = "go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws"
)

type spanTimestampKey struct{}

// AttributeBuilder returns an array of KeyValue pairs, it can be used to set custom attributes.
type AttributeBuilder func(ctx context.Context, in middleware.InitializeInput, out middleware.InitializeOutput) []attribute.KeyValue

type otelMiddlewares struct {
	tracer            trace.Tracer
	propagator        propagation.TextMapPropagator
	attributeBuilders []AttributeBuilder
}

func (otelMiddlewares) initializeMiddlewareBefore(stack *middleware.Stack) error {
	_ = "STUB: not implemented"
	return nil
}

func (m otelMiddlewares) initializeMiddlewareAfter(stack *middleware.Stack) error {
	_ = "STUB: not implemented"
	return nil
}

func (m otelMiddlewares) finalizeMiddlewareAfter(stack *middleware.Stack) error {
	_ = "STUB: not implemented"
	return nil
}

// Propagate the Trace information by injecting it into the HTTP request.

func (otelMiddlewares) deserializeMiddleware(stack *middleware.Stack) error {
	_ = "STUB: not implemented"
	return nil
}

// No raw response to wrap with.

func (m otelMiddlewares) buildAttributes(ctx context.Context, in middleware.InitializeInput, out middleware.InitializeOutput) (attributes []attribute.KeyValue) {
	_ = "STUB: not implemented"
	return nil
}

func spanName(serviceID, operation string) string { _ = "STUB: not implemented"; return "" }

// AppendMiddlewares attaches OTel middlewares to the AWS Go SDK V2 for instrumentation.
// OTel middlewares can be appended to either all aws clients or a specific operation.
// Please see more details in https://aws.github.io/aws-sdk-go-v2/docs/middleware/
func AppendMiddlewares(apiOptions *[]func(*middleware.Stack) error, opts ...Option) {
	_ = "STUB: not implemented"
	return
}
