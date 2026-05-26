// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package handler provides an HTTP handler.
package handler // import "go.opentelemetry.io/contrib/examples/passthrough/handler"

import (
	"context"
	"net/http"

	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// Handler is a minimal implementation of the handler and client from
// go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp for demonstration purposes.
// It handles an incoming http request, and makes an outgoing http request.
type Handler struct {
	propagators propagation.TextMapPropagator
	tracer      trace.Tracer
	next        func(r *http.Request)
}

// New returns a new Handler that will trace requests before handing them off
// to next.
func New(next func(r *http.Request)) *Handler {
	_ = "STUB: not implemented"
	// Like most instrumentation packages, this handler defaults to using the
	// global propagators and tracer providers.
	return nil
}

// HandleHTTPReq mimics what an instrumented http server does.
func (h *Handler) HandleHTTPReq(r *http.Request) { _ = "STUB: not implemented"; return }

// Pretend to do work

// makeOutgoingRequest mimics what an instrumented http client does.
func (h *Handler) makeOutgoingRequest(ctx context.Context) {
	_ = "STUB: not implemented"
	// make a new http request
	return
}
