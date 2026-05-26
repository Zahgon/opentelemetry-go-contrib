// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package filters provides a set of filters useful with the
// [otelgrpc.WithFilter] option to control which inbound requests are instrumented.
package filters // import "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc/filters"

import (
	"google.golang.org/grpc/stats"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
)

type gRPCPath struct {
	service string
	method  string
}

// splitFullMethod splits path defined in gRPC protocol
// and returns as gRPCPath object that has divided service and method names
// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md
// If name is not FullMethod, returned gRPCPath has empty service field.
func splitFullMethod(i *stats.RPCTagInfo) gRPCPath {
	_ = "STUB: not implemented"
	return *new(gRPCPath)
}

// Any takes a list of Filters and returns a Filter that
// returns true if any Filter in the list returns true.
func Any(fs ...otelgrpc.Filter) otelgrpc.Filter {
	_ = "STUB: not implemented"
	return *new(otelgrpc.Filter)
}

// All takes a list of Filters and returns a Filter that
// returns true only if all Filters in the list return true.
func All(fs ...otelgrpc.Filter) otelgrpc.Filter {
	_ = "STUB: not implemented"
	return *new(otelgrpc.Filter)
}

// None takes a list of Filters and returns a Filter that returns
// true only if none of the Filters in the list return true.
func None(fs ...otelgrpc.Filter) otelgrpc.Filter {
	_ = "STUB: not implemented"
	return *

	// Not provides a convenience mechanism for inverting a Filter.
	new(otelgrpc.Filter)
}

func Not(f otelgrpc.Filter) otelgrpc.Filter {
	_ = "STUB: not implemented"
	return *new(otelgrpc.Filter)
}

// MethodName returns a Filter that returns true if the request's
// method name matches the provided string n.
func MethodName(n string) otelgrpc.Filter { _ = "STUB: not implemented"; return *new(otelgrpc.Filter) }

// MethodPrefix returns a Filter that returns true if the request's
// method starts with the provided string pre.
func MethodPrefix(pre string) otelgrpc.Filter {
	_ = "STUB: not implemented"
	return *new(otelgrpc.Filter)
}

// FullMethodName returns a Filter that returns true if the request's
// full RPC method string, i.e. /package.service/method, starts with
// the provided string n.
func FullMethodName(n string) otelgrpc.Filter {
	_ = "STUB: not implemented"
	return *new(otelgrpc.Filter)
}

// ServiceName returns a Filter that returns true if the request's
// service name, i.e. package.service, matches s.
func ServiceName(s string) otelgrpc.Filter { _ = "STUB: not implemented"; return *new(otelgrpc.Filter) }

// ServicePrefix returns a Filter that returns true if the request's
// service name, i.e. package.service, starts with the provided string pre.
func ServicePrefix(pre string) otelgrpc.Filter {
	_ = "STUB: not implemented"
	return *new(otelgrpc.Filter)
}

// HealthCheck returns a Filter that returns true if the request's
// service name is health check defined by gRPC Health Checking Protocol.
// https://github.com/grpc/grpc/blob/master/doc/health-checking.md
func HealthCheck() otelgrpc.Filter { _ = "STUB: not implemented"; return *new(otelgrpc.Filter) }
