// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Server exemplifies the otelhttp instrumentation for a server.
package main

import (
	"context"
	"io"
	"log"
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func initTracer() (*sdktrace.TracerProvider, error) {
	_ = "STUB: not implemented"
	// Create stdout exporter to be able to retrieve
	// the collected spans.
	return nil, nil
}

// For the demonstration, use sdktrace.AlwaysSample sampler to sample all traces.
// In a production application, use sdktrace.ProbabilitySampler with a desired probability.

func initMeter() (*sdkmetric.MeterProvider, error) { _ = "STUB: not implemented"; return nil, nil }

func main() {
	tp, err := initTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	mp, err := initMeter()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := mp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down meter provider: %v", err)
		}
	}()

	uk := attribute.Key("username")

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		span := trace.SpanFromContext(ctx)
		bag := baggage.FromContext(ctx)
		span.AddEvent("handling this...", trace.WithAttributes(uk.String(bag.Member("username").Value())))

		_, _ = io.WriteString(w, "Hello, world!\n")
	}

	otelHandler := otelhttp.NewHandler(http.HandlerFunc(helloHandler), "Hello")

	http.Handle("/hello", otelHandler)
	err = http.ListenAndServe(":7777", nil) //nolint:gosec // Ignoring G114: Use of net/http serve function that has no support for setting timeouts.
	if err != nil {
		log.Fatal(err)
	}
}
