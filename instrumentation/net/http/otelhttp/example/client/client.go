// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Client exemplifies the otelhttp instrumentation for a client.
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/baggage"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.41.0"
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
	url := flag.String("server", "http://localhost:7777/hello", "server url")
	flag.Parse()

	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}

	bag, _ := baggage.Parse("username=donuts")
	ctx := baggage.ContextWithBaggage(context.Background(), bag)

	var body []byte

	tr := otel.Tracer("example/client")
	err = func(ctx context.Context) error {
		ctx, span := tr.Start(ctx, "say hello", trace.WithAttributes(semconv.ServicePeerName("ExampleService")))
		defer span.End()
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, *url, http.NoBody)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Sending request...\n")
		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		body, err = io.ReadAll(res.Body)
		_ = res.Body.Close()

		return err
	}(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Response Received: %s\n\n\n", body)
	fmt.Printf("Waiting for few seconds to export spans ...\n\n")
	time.Sleep(10 * time.Second)
	fmt.Printf("Inspect traces on stdout\n")
}
