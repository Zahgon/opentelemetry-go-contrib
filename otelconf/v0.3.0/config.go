// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package otelconf provides an OpenTelemetry declarative configuration SDK.
package otelconf // import "go.opentelemetry.io/contrib/otelconf/v0.3.0"

import (
	"context"

	"go.opentelemetry.io/otel/log"
	nooplog "go.opentelemetry.io/otel/log/noop"
	"go.opentelemetry.io/otel/metric"
	noopmetric "go.opentelemetry.io/otel/metric/noop"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdkresource "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	nooptrace "go.opentelemetry.io/otel/trace/noop"
)

const (
	protocolProtobufHTTP = "http/protobuf"
	protocolProtobufGRPC = "grpc"

	compressionGzip = "gzip"
	compressionNone = "none"
)

type configOptions struct {
	ctx                   context.Context
	opentelemetryConfig   OpenTelemetryConfiguration
	loggerProviderOptions []sdklog.LoggerProviderOption
	meterProviderOptions  []sdkmetric.Option
	tracerProviderOptions []sdktrace.TracerProviderOption
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
	resource       *sdkresource.Resource
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

	// Resource returns a copy of the resolved SDK resource configured in this SDK.
	// The copy preserves the immutability of the SDK-owned resource.
	new(log.LoggerProvider)
}

func (s *SDK) Resource() *sdkresource.Resource { _ = "STUB: not implemented"; return nil }

// Shutdown calls shutdown on all configured providers.
func (s *SDK) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

var noopSDK = SDK{
	loggerProvider: nooplog.LoggerProvider{},
	meterProvider:  noopmetric.MeterProvider{},
	tracerProvider: nooptrace.TracerProvider{},
	resource:       sdkresource.Empty(),
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

// WithLoggerProviderOptions appends LoggerProviderOptions used for constructing
// the LoggerProvider. OpenTelemetryConfiguration takes precedence over these options.
func WithLoggerProviderOptions(opts ...sdklog.LoggerProviderOption) ConfigurationOption {
	_ = "STUB: not implemented"
	return *new(ConfigurationOption)
}

// WithMeterProviderOptions appends metric.Options used for constructing the
// MeterProvider. OpenTelemetryConfiguration takes precedence over these options.
func WithMeterProviderOptions(opts ...sdkmetric.Option) ConfigurationOption {
	_ = "STUB: not implemented"
	return *new(ConfigurationOption)
}

// WithTracerProviderOptions appends TracerProviderOptions used for constructing
// the TracerProvider. OpenTelemetryConfiguration takes precedence over these options.
func WithTracerProviderOptions(opts ...sdktrace.TracerProviderOption) ConfigurationOption {
	_ = "STUB: not implemented"
	return *new(ConfigurationOption)
}

// ParseYAML parses a YAML configuration file into an OpenTelemetryConfiguration.
func ParseYAML(file []byte) (*OpenTelemetryConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// createHeadersConfig combines the two header config fields. Headers take precedence over headersList.
func createHeadersConfig(headers []NameStringValuePair, headersList *string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parsing follows https://github.com/open-telemetry/opentelemetry-configuration/blob/568e5080816d40d75792eb754fc96bde09654159/schema/type_descriptions.yaml#L584.

// Headers take precedence over HeadersList, so this has to be after HeadersList is processed
