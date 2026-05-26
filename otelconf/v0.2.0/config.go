// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package otelconf provides an OpenTelemetry declarative configuration SDK.
package otelconf // import "go.opentelemetry.io/contrib/otelconf/v0.2.0"

import (
	"context"

	"go.opentelemetry.io/otel/log"
	nooplog "go.opentelemetry.io/otel/log/noop"
	"go.opentelemetry.io/otel/metric"
	noopmetric "go.opentelemetry.io/otel/metric/noop"
	"go.opentelemetry.io/otel/trace"
	nooptrace "go.opentelemetry.io/otel/trace/noop"
)

const (
	protocolProtobufHTTP = "http/protobuf"
	protocolProtobufGRPC = "grpc/protobuf"

	compressionGzip = "gzip"
	compressionNone = "none"
)

type configOptions struct {
	ctx                 context.Context
	opentelemetryConfig OpenTelemetryConfiguration
}

type shutdownFunc func(context.Context) error

func noopShutdown(context.Context) error {
	_ = "STUB: not implemented"

	// SDK is a struct that contains all the providers
	// configured via the configuration model.
	return nil
}

type SDK struct {
	meterProvider  metric.MeterProvider
	tracerProvider trace.TracerProvider
	loggerProvider log.LoggerProvider
	shutdown       shutdownFunc
}

// TracerProvider returns a configured trace.TracerProvider.
func (s *SDK) TracerProvider() trace.TracerProvider {
	_ = "STUB: not implemented"
	return *

	// MeterProvider returns a configured metric.MeterProvider.
	new(trace.TracerProvider)
}

func (s *SDK) MeterProvider() metric.MeterProvider {
	_ = "STUB: not implemented"
	return *

	// LoggerProvider returns a configured log.LoggerProvider.
	new(metric.MeterProvider)
}

func (s *SDK) LoggerProvider() log.LoggerProvider {
	_ = "STUB: not implemented"
	return *

	// Shutdown calls shutdown on all configured providers.
	new(log.LoggerProvider)
}

func (s *SDK) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

var noopSDK = SDK{
	loggerProvider: nooplog.LoggerProvider{},
	meterProvider:  noopmetric.MeterProvider{},
	tracerProvider: nooptrace.TracerProvider{},
	shutdown:       func(context.Context) error { return nil },
}

// NewSDK creates SDK providers based on the configuration model.
func NewSDK(opts ...ConfigurationOption) (SDK, error) {
	_ = "STUB: not implemented"
	return *new(SDK), nil
}

// ConfigurationOption configures options for providers.
type ConfigurationOption interface {
	apply(configOptions) configOptions
}

type configurationOptionFunc func(configOptions) configOptions

func (fn configurationOptionFunc) apply(cfg configOptions) configOptions {
	_ = "STUB: not implemented"

	// WithContext sets the context.Context for the SDK.
	return *new(configOptions)
}

func WithContext(ctx context.Context) ConfigurationOption {
	_ = "STUB: not implemented"
	return *new(ConfigurationOption)
}

// WithOpenTelemetryConfiguration sets the OpenTelemetryConfiguration used
// to produce the SDK.
func WithOpenTelemetryConfiguration(cfg OpenTelemetryConfiguration) ConfigurationOption {
	_ = "STUB: not implemented"
	return *new(ConfigurationOption)
}

// ParseYAML parses a YAML configuration file into an OpenTelemetryConfiguration.
func ParseYAML(file []byte) (*OpenTelemetryConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
