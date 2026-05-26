// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

/*
 *
 * Copyright 2014 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package test contains functions used by interop client/server.
//
// Copied from https://github.com/grpc/grpc-go/tree/v1.61.0/interop
// That package was not intended to be used by external code.
// See https://github.com/open-telemetry/opentelemetry-go-contrib/issues/4896
package test // import "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc/internal/test"

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	testpb "google.golang.org/grpc/interop/grpc_testing"
)

var (
	reqSizes                  = []int32{27182, 8, 1828, 45904}
	respSizes                 = []int32{31415, 9, 2653, 58979}
	largeReqSize              = 271828
	largeRespSize       int32 = 314159
	initialMetadataKey        = "x-grpc-test-echo-initial"
	trailingMetadataKey       = "x-grpc-test-echo-trailing-bin"

	logger = grpclog.Component("interop")
)

// ClientNewPayload returns a payload of the given type and size.
func ClientNewPayload(t testpb.PayloadType, size int) *testpb.Payload {
	_ = "STUB: not implemented"
	return nil
}

// DoEmptyUnaryCall performs a unary RPC with empty request and response messages.
func DoEmptyUnaryCall(ctx context.Context, tc testpb.TestServiceClient, args ...grpc.CallOption) {
	_ = "STUB: not implemented"
	return
}

// DoLargeUnaryCall performs a unary RPC with large payload in the request and response.
func DoLargeUnaryCall(ctx context.Context, tc testpb.TestServiceClient, args ...grpc.CallOption) {
	_ = "STUB: not implemented"
	return
}

// DoClientStreaming performs a client streaming RPC.
func DoClientStreaming(ctx context.Context, tc testpb.TestServiceClient, args ...grpc.CallOption) {
	_ = "STUB: not implemented"
	return
}

// DoServerStreaming performs a server streaming RPC.
func DoServerStreaming(ctx context.Context, tc testpb.TestServiceClient, args ...grpc.CallOption) {
	_ = "STUB: not implemented"
	return
}

// DoPingPong performs ping-pong style bi-directional streaming RPC.
func DoPingPong(ctx context.Context, tc testpb.TestServiceClient, args ...grpc.CallOption) {
	_ = "STUB: not implemented"
	return
}

// DoEmptyStream sets up a bi-directional streaming with zero message.
func DoEmptyStream(ctx context.Context, tc testpb.TestServiceClient, args ...grpc.CallOption) {
	_ = "STUB: not implemented"
	return
}

type testServer struct {
	testpb.UnimplementedTestServiceServer
}

// NewTestServer creates a test server for test service.  opts carries optional
// settings and does not need to be provided.  If multiple opts are provided,
// only the first one is used.
func NewTestServer() testpb.TestServiceServer {
	_ = "STUB: not implemented"
	return *new(testpb.TestServiceServer)
}

func (*testServer) EmptyCall(context.Context, *testpb.Empty) (*testpb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func serverNewPayload(t testpb.PayloadType, size int32) (*testpb.Payload, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*testServer) UnaryCall(ctx context.Context, in *testpb.SimpleRequest) (*testpb.SimpleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // Overflow handled by status code bounds in grpc_testing package.

func (*testServer) StreamingOutputCall(args *testpb.StreamingOutputCallRequest, stream testpb.TestService_StreamingOutputCallServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (*testServer) StreamingInputCall(stream testpb.TestService_StreamingInputCallServer) error {
	_ = "STUB: not implemented"
	return nil
}

// This could overflow, but given this is a test and the negative value
// should be detectable this should be good enough.
//nolint:gosec // See comment above.

func (*testServer) FullDuplexCall(stream testpb.TestService_FullDuplexCallServer) error {
	_ = "STUB: not implemented"
	return nil
}

// read done.

//nolint:gosec // Overflow handled by status code bounds in grpc_testing package.

func (*testServer) HalfDuplexCall(stream testpb.TestService_HalfDuplexCallServer) error {
	_ = "STUB: not implemented"
	return nil
}

// read done.
