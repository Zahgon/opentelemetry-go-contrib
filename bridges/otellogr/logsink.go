// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package otellogr provides a [LogSink], a [logr.LogSink] implementation that
// can be used to bridge between the [logr] API and [OpenTelemetry].
//
// # Record Conversion
//
// The logr records are converted to OpenTelemetry [log.Record] in the following
// way:
//
//   - Message is set as the Body using a [log.StringValue].
//   - Level is transformed and set as the Severity. The SeverityText is not
//     set.
//   - KeyAndValues are transformed and set as Attributes.
//   - Error is always set as the record error with the severity
//     [log.SeverityError].
//   - The [context.Context] value in KeyAndValues is propagated to OpenTelemetry
//     log record. All non-nested [context.Context] values are ignored and not
//     added as attributes. If there are multiple [context.Context] the last one
//     is used.
//
// The V-level is transformed by using the [WithLevelSeverity] option. If option is
// not provided then V-level is transformed in the following way:
//
//   - logr.Info and logr.V(0) are transformed to [log.SeverityInfo].
//   - logr.V(1) is transformed to [log.SeverityDebug].
//   - logr.V(2) and higher are transformed to [log.SeverityTrace].
//
// KeysAndValues values are transformed based on their type. The following types are
// supported:
//
//   - [bool] are transformed to [log.BoolValue].
//   - [string] are transformed to [log.StringValue].
//   - [int], [int8], [int16], [int32], [int64] are transformed to
//     [log.Int64Value].
//   - [uint], [uint8], [uint16], [uint32], [uint64], [uintptr] are transformed
//     to [log.Int64Value] or [log.StringValue] if the value is too large.
//   - [float32], [float64] are transformed to [log.Float64Value].
//   - [time.Duration] are transformed to [log.Int64Value] with the nanoseconds.
//   - [complex64], [complex128] are transformed to [log.MapValue] with the keys
//     "r" and "i" for the real and imaginary parts. The values are
//     [log.Float64Value].
//   - [time.Time] are transformed to [log.Int64Value] with the nanoseconds.
//   - [[]byte] are transformed to [log.BytesValue].
//   - [error] are transformed to [log.StringValue] with the error message.
//   - [nil] are transformed to an empty [log.Value].
//   - [struct] are transformed to [log.StringValue] with the struct fields.
//   - [slice], [array] are transformed to [log.SliceValue] with the elements.
//   - [map] are transformed to [log.MapValue] with the key-value pairs.
//   - [pointer], [interface] are transformed to the dereferenced value.
//
// [OpenTelemetry]: https://opentelemetry.io/docs/concepts/signals/logs/
package otellogr // import "go.opentelemetry.io/contrib/bridges/otellogr"

import (
	"context"

	"github.com/go-logr/logr"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/log"
)

type config struct {
	provider   log.LoggerProvider
	version    string
	schemaURL  string
	attributes []attribute.KeyValue

	levelSeverity func(int) log.Severity
}

func newConfig(options []Option) config { _ = "STUB: not implemented"; return *new(config) }

// Option configures a [LogSink].
type Option interface {
	apply(config) config
}

type optFunc func(config) config

func (f optFunc) apply(c config) config {
	_ = "STUB: not implemented"

	// WithVersion returns an [Option] that configures the version of the
	// [log.Logger] used by a [LogSink]. The version should be the version of the
	// package that is being logged.
	return *new(config)
}

func WithVersion(version string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSchemaURL returns an [Option] that configures the semantic convention
// schema URL of the [log.Logger] used by a [LogSink]. The schemaURL should be
// the schema URL for the semantic conventions used in log records.
func WithSchemaURL(schemaURL string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAttributes returns an [Option] that configures the instrumentation scope
// attributes of the [log.Logger] used by a [LogSink].
func WithAttributes(attributes ...attribute.KeyValue) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithLoggerProvider returns an [Option] that configures [log.LoggerProvider]
// used by a [LogSink] to create its [log.Logger].
//
// By default if this Option is not provided, the LogSink will use the global
// LoggerProvider.
func WithLoggerProvider(provider log.LoggerProvider) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithLevelSeverity returns an [Option] that configures the function used to
// convert logr levels to OpenTelemetry log severities.
//
// By default if this Option is not provided, the LogSink will use a default
// conversion function that transforms in the following way:
//
//   - logr.Info and logr.V(0) are transformed to [log.SeverityInfo].
//   - logr.V(1) is transformed to [log.SeverityDebug].
//   - logr.V(2) and higher are transformed to [log.SeverityTrace].
func WithLevelSeverity(f func(int) log.Severity) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// NewLogSink returns a new [LogSink] to be used as a [logr.LogSink].
//
// If [WithLoggerProvider] is not provided, the returned [LogSink] will use the
// global LoggerProvider.
func NewLogSink(name string, options ...Option) *LogSink { _ = "STUB: not implemented"; return nil }

// LogSink is a [logr.LogSink] that sends all logging records it receives to
// OpenTelemetry. See package documentation for how conversions are made.
type LogSink struct {
	// Ensure forward compatibility by explicitly making this not comparable.
	noCmp [0]func() //nolint:unused  // This is indeed used.

	name          string
	provider      log.LoggerProvider
	logger        log.Logger
	levelSeverity func(int) log.Severity
	opts          []log.LoggerOption
	attr          []log.KeyValue
	ctx           context.Context
}

// Compile-time check *Handler implements logr.LogSink.
var _ logr.LogSink = (*LogSink)(nil)

// Enabled tests whether this LogSink is enabled at the specified V-level.
// For example, commandline flags might be used to set the logging
// verbosity and disable some info logs.
func (l *LogSink) Enabled(level int) bool { _ = "STUB: not implemented"; return false }

// Error logs an error, with the given message and key/value pairs.
func (l *LogSink) Error(err error, msg string, keysAndValues ...any) {
	_ = "STUB: not implemented"
	return
}

// Info logs a non-error message with the given key/value pairs.
func (l *LogSink) Info(level int, msg string, keysAndValues ...any) {
	_ = "STUB: not implemented"
	return
}

// Init receives optional information about the logr library this
// implementation does not use it.
func (*LogSink) Init(logr.RuntimeInfo) {
	_ = "STUB: not implemented"
	// We don't need to do anything here.
	// CallDepth is used to calculate the caller's PC.
	// PC is dropped as part of the conversion to the OpenTelemetry log.Record.
	return
}

// WithName returns a new LogSink with the specified name appended.
func (l LogSink) WithName(name string) logr.LogSink {
	_ = "STUB: not implemented"
	return *new(logr.LogSink)
}

// WithValues returns a new LogSink with additional key/value pairs.
func (l LogSink) WithValues(keysAndValues ...any) logr.LogSink {
	_ = "STUB: not implemented"
	return *new(logr.LogSink)
}

// convertKVs converts a list of key-value pairs to a list of [log.KeyValue].
// The last [context.Context] value is returned as the context.
// If no context is found, the original context is returned.
func convertKVs(ctx context.Context, keysAndValues ...any) (context.Context, []log.KeyValue) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// Ensure an odd number of items here does not corrupt the list.

// Ensure that the key is a string.

// Special case when a field is of context.Context type.
