// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelgrpc // import "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

import (
	"context"

	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc/metadata"
)

type metadataSupplier struct {
	metadata metadata.MD
}

// assert that metadataSupplier implements the TextMapCarrier interface.
var _ propagation.TextMapCarrier = metadataSupplier{}

func (s metadataSupplier) Get(key string) string { _ = "STUB: not implemented"; return "" }

func (s metadataSupplier) Set(key, value string) { _ = "STUB: not implemented"; return }

func (s metadataSupplier) Keys() []string { _ = "STUB: not implemented"; return nil }

func inject(ctx context.Context, propagators propagation.TextMapPropagator) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func extract(ctx context.Context, propagators propagation.TextMapPropagator) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
