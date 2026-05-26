// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otellambda // import "go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"

import (
	"context"
	"log"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const (
	// ScopeName is the instrumentation scope name.
	ScopeName = "go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"
)

var errorLogger = log.New(log.Writer(), "OTel Lambda Error: ", 0)

type instrumentor struct {
	configuration config
	resAttrs      []attribute.KeyValue
	tracer        trace.Tracer
}

func newInstrumentor(opts ...Option) instrumentor {
	_ = "STUB: not implemented"
	return *new(instrumentor)
}

// Logic to start OTel Tracing.
func (i *instrumentor) tracingBegin(ctx context.Context, eventJSON []byte) (context.Context, trace.Span) {
	_ = "STUB: not implemented"
	// Add trace id to context
	return *new(context.Context), *new(trace.Span)
}

// Some resource attrs added as span attrs because lambda
// resource detectors are created before a lambda
// invocation and therefore lack lambdacontext.
// Create these attrs upon first invocation

// Logic to wrap up OTel Tracing.
func (i *instrumentor) tracingEnd(ctx context.Context, span trace.Span) {
	_ = "STUB: not implemented"

	// force flush any tracing data since lambda may freeze
	return
}
