// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Passthrough exemplifies distributed context propagation.
package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/contrib/examples/passthrough/handler"
)

const name = "go.opentelemetry.io/contrib/examples/passthrough"

func main() {
	ctx := context.Background()

	initPassthroughGlobals()
	tp, err := nonGlobalTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = tp.Shutdown(ctx) }()

	// make an initial http request
	r, err := http.NewRequestWithContext(ctx, "", "", http.NoBody)
	if err != nil {
		panic(err)
	}

	// This is roughly what an instrumented http client does.
	log.Println("The \"make outer request\" span should be recorded, because it is recorded with a Tracer from the SDK TracerProvider")
	var span trace.Span
	tracer := tp.Tracer(name)
	ctx, span = tracer.Start(ctx, "make outer request")
	defer span.End()
	r = r.WithContext(ctx)
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(r.Header))

	backendFunc := func(r *http.Request) {
		// This is roughly what an instrumented http server does.
		ctx := r.Context()

		tp := trace.SpanFromContext(ctx).TracerProvider()
		tracer := tp.Tracer(name)

		ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(r.Header))
		log.Println("The \"handle inner request\" span should be recorded, because it is recorded with a Tracer from the SDK TracerProvider")
		_, span := tracer.Start(ctx, "handle inner request")
		defer span.End()

		// Do "backend work"
		time.Sleep(time.Second)
	}
	// This handler will be a passthrough, since we didn't set a global TracerProvider
	passthroughHandler := handler.New(backendFunc)
	passthroughHandler.HandleHTTPReq(r)
}

func initPassthroughGlobals() {
	_ = "STUB: not implemented"
	// We explicitly DO NOT set the global TracerProvider using otel.SetTracerProvider().
	// The unset TracerProvider returns a "non-recording" span, but still passes through context.
	return
}

// nonGlobalTracer creates a trace provider instance for testing, but doesn't
// set it as the global tracer provider.
func nonGlobalTracer() (*sdktrace.TracerProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
