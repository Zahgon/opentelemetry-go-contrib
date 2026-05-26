// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelgrpc // import "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	oldrpcconv "go.opentelemetry.io/otel/semconv/v1.37.0/rpcconv" //nolint:depguard // Use of v1.37.0 is required for backward compatibility stability opt-in.
	"go.opentelemetry.io/otel/semconv/v1.41.0/rpcconv"
	"go.opentelemetry.io/otel/trace"

	grpc_codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/status"
)

type gRPCContextKey struct{}

type gRPCContext struct {
	metricAttrs []attribute.KeyValue
	record      bool
}

type serverHandler struct {
	*config

	tracer trace.Tracer

	duration    rpcconv.ServerCallDuration
	oldDuration oldrpcconv.ServerDuration
}

// NewServerHandler creates a stats.Handler for a gRPC server.
func NewServerHandler(opts ...Option) stats.Handler {
	_ = "STUB: not implemented"
	return *new(stats.Handler)
}

// TagConn can attach some information to the given context.
func (*serverHandler) TagConn(ctx context.Context, _ *stats.ConnTagInfo) context.Context {
	_ = "STUB: not implemented"

	// HandleConn processes the Conn stats.
	return *new(context.Context)
}

func (*serverHandler) HandleConn(context.Context, stats.ConnStats) {
	_ = "STUB: not implemented"

	// TagRPC can attach some information to the given context.
	return
}

func (h *serverHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Combine both. We append New last so its rpc.method (fully qualified) wins when deduplicated.

// New convention
// semconvModeNew

// Make a new slice to avoid aliasing into the same attrs slice used by metrics.

// Linking incoming span context if any for public endpoint.

// HandleRPC processes the RPC stats.
func (h *serverHandler) HandleRPC(ctx context.Context, rs stats.RPCStats) {
	_ = "STUB: not implemented"
	return
}

type clientHandler struct {
	*config

	tracer trace.Tracer

	duration    rpcconv.ClientCallDuration
	oldDuration oldrpcconv.ClientDuration
}

// NewClientHandler creates a stats.Handler for a gRPC client.
func NewClientHandler(opts ...Option) stats.Handler {
	_ = "STUB: not implemented"
	return *new(stats.Handler)
}

// TagRPC can attach some information to the given context.
func (h *clientHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Combine both. We append New last so its rpc.method (fully qualified) wins when deduplicated.

// New convention
// semconvModeNew

// Make a new slice to avoid aliasing into the same attrs slice used by metrics.

// HandleRPC processes the RPC stats.
func (h *clientHandler) HandleRPC(ctx context.Context, rs stats.RPCStats) {
	_ = "STUB: not implemented"
	return
}

// TagConn can attach some information to the given context.
func (*clientHandler) TagConn(ctx context.Context, _ *stats.ConnTagInfo) context.Context {
	_ = "STUB: not implemented"

	// HandleConn processes the Conn stats.
	return *new(context.Context)
}

func (*clientHandler) HandleConn(context.Context, stats.ConnStats) {
	_ = "STUB: not implemented"
	// no-op
	return
}

func (*config) handleRPC(
	ctx context.Context,
	rs stats.RPCStats,
	duration metric.Float64Histogram,
	oldDuration metric.Float64Histogram,
	recordStatus func(*status.Status) (codes.Code, string),
) {
	_ = "STUB: not implemented"
	return
}

// TODO: add server.address and server.port to metrics once the API supports opt-in attributes.

// Don't use gctx.metricAttrSet here, because it requires passing
// multiple RecordOptions, which would call metric.mergeSets and
// allocate a new set for each Record call.

// Allocate vararg slice once.

// Use floating point division here for higher precision (instead of Millisecond method).
// Measure right before calling Record() to capture as much elapsed time as possible.

func canonicalString(code grpc_codes.Code) string { _ = "STUB: not implemented"; return "" }
