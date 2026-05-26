// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Client exemplifies use of the otelgrpc instrumentation for a gRPC client.
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc/example/api"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc/example/config"
)

func main() {
	tp, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	var conn *grpc.ClientConn
	conn, err = grpc.NewClient(
		"127.0.0.1:7777", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer func() { _ = conn.Close() }()

	c := api.NewHelloServiceClient(conn)

	if err := callSayHello(c); err != nil {
		log.Fatal(err)
	}
	if err := callSayHelloClientStream(c); err != nil {
		log.Fatal(err)
	}
	if err := callSayHelloServerStream(c); err != nil {
		log.Fatal(err)
	}
	if err := callSayHelloBidiStream(c); err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Millisecond)
}

func callSayHello(c api.HelloServiceClient) error { _ = "STUB: not implemented"; return nil }

func callSayHelloClientStream(c api.HelloServiceClient) error {
	_ = "STUB: not implemented"
	return nil
}

func callSayHelloServerStream(c api.HelloServiceClient) error {
	_ = "STUB: not implemented"
	return nil
}

func callSayHelloBidiStream(c api.HelloServiceClient) error { _ = "STUB: not implemented"; return nil }

//nolint:revive  // This acts as its own main func.

//nolint:revive  // This acts as its own main func.

//nolint:revive  // This acts as its own main func.

// Wait until client and server both closed the connection.
