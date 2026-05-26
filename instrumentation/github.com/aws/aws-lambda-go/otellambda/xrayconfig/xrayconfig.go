// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package xrayconfig provides AWS XRAY configuration for otellambda.
package xrayconfig // import "go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda/xrayconfig"

import (
	"context"

	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace" //nolint:depguard // NewTracerProvider requires the SDK

	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"
)

func xrayEventToCarrier([]byte) propagation.TextMapCarrier {
	_ = "STUB: not implemented"
	return *new(propagation.TextMapCarrier)
}

// NewTracerProvider returns a TracerProvider configured with an exporter,
// ID generator, and lambda resource detector to send trace data to AWS X-Ray
// via a Collector instance listening on localhost.
func NewTracerProvider(ctx context.Context) (*sdktrace.TracerProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WithEventToCarrier returns an otellambda.Option to enable
// an otellambda.EventToCarrier function which reads the XRay trace
// information from the environment and returns this information in
// a propagation.HeaderCarrier.
func WithEventToCarrier() otellambda.Option {
	_ = "STUB: not implemented"
	return *new(otellambda.Option)
}

// WithPropagator returns an otellambda.Option to enable the xray.Propagator.
func WithPropagator() otellambda.Option { _ = "STUB: not implemented"; return *new(otellambda.Option) }

// WithRecommendedOptions returns a list of all otellambda.Option(s)
// recommended for the otellambda package when using AWS XRay.
func WithRecommendedOptions(tp *sdktrace.TracerProvider) []otellambda.Option {
	_ = "STUB: not implemented"
	return nil
}
