// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Server exemplifies use of the otelgrpc instrumentation for a gRPC server.
package main

import (
	"context"
	"log"
	"net"

	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc/example/api"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc/example/config"
)

var tracer = otel.Tracer("grpc-example")

// server is used to implement api.HelloServiceServer.
type server struct {
	api.HelloServiceServer
}

// SayHello implements api.HelloServiceServer.
func (s *server) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*server) workHard(ctx context.Context) { _ = "STUB: not implemented"; return }

func (*server) SayHelloServerStream(in *api.HelloRequest, out api.HelloService_SayHelloServerStreamServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (*server) SayHelloClientStream(stream api.HelloService_SayHelloClientStreamServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (*server) SayHelloBidiStream(stream api.HelloService_SayHelloBidiStreamServer) error {
	_ = "STUB: not implemented"
	return nil
}

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

	lis, err := (&net.ListenConfig{}).Listen(context.Background(), "tcp", "127.0.0.1:7777")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	api.RegisterHelloServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
