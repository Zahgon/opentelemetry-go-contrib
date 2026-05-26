// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package otelzap provides a bridge between the [go.uber.org/zap] and
// [OpenTelemetry].
//
// # Record Conversion
//
// The [zapcore.Entry] and [zapcore.Field] are converted to OpenTelemetry [log.Record] in the following
// way:
//
//   - Time is set as the Timestamp.
//   - Message is set as the Body using a [log.StringValue].
//   - Level is transformed and set as the Severity. The SeverityText is also
//     set.
//   - Fields are transformed and set as the Attributes.
//   - Fields of type [error] are attached to the emitted record as an error via
//     [log.Record.SetErr].
//   - Field value of type [context.Context] is used as context when emitting log records.
//   - For named loggers, LoggerName is used to access [log.Logger] from [log.LoggerProvider]
//
// The Level is transformed to the OpenTelemetry Severity types in the following way.
//
//   - [zapcore.DebugLevel] is transformed to [log.SeverityDebug]
//   - [zapcore.InfoLevel] is transformed to [log.SeverityInfo]
//   - [zapcore.WarnLevel] is transformed to [log.SeverityWarn]
//   - [zapcore.ErrorLevel] is transformed to [log.SeverityError]
//   - [zapcore.DPanicLevel] is transformed to [log.SeverityFatal1]
//   - [zapcore.PanicLevel] is transformed to [log.SeverityFatal2]
//   - [zapcore.FatalLevel] is transformed to [log.SeverityFatal3]
//
// Fields are transformed based on their type into log attributes, or
// into a string value encoded using [fmt.Sprintf] if there is no matching type.
//
// [OpenTelemetry]: https://opentelemetry.io/docs/concepts/signals/logs/
package otelzap // import "go.opentelemetry.io/contrib/bridges/otelzap"

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/log"
	semconv "go.opentelemetry.io/otel/semconv/v1.41.0"
	"go.uber.org/zap/zapcore"
)

var (
	exceptionMessageKey = string(semconv.ExceptionMessageKey)
	exceptionTypeKey    = string(semconv.ExceptionTypeKey)
)

type config struct {
	provider   log.LoggerProvider
	version    string
	schemaURL  string
	attributes []attribute.KeyValue
}

func newConfig(options []Option) config { _ = "STUB: not implemented"; return *new(config) }

// Option configures a [Core].
type Option interface {
	apply(config) config
}

type optFunc func(config) config

func (f optFunc) apply(c config) config {
	_ = "STUB: not implemented"

	// WithVersion returns an [Option] that configures the version of the
	// [log.Logger] used by a [Core]. The version should be the version of the
	// package that is being logged.
	return *new(config)
}

func WithVersion(version string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSchemaURL returns an [Option] that configures the semantic convention
// schema URL of the [log.Logger] used by a [Core]. The schemaURL should be
// the schema URL for the semantic conventions used in log records.
func WithSchemaURL(schemaURL string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAttributes returns an [Option] that configures the instrumentation scope
// attributes of the [log.Logger] used by a [Core].
func WithAttributes(attributes ...attribute.KeyValue) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithLoggerProvider returns an [Option] that configures [log.LoggerProvider]
// used by a [Core] to create its [log.Logger].
//
// By default if this Option is not provided, the Handler will use the global
// LoggerProvider.
func WithLoggerProvider(provider log.LoggerProvider) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Core is a [zapcore.Core] that sends logging records to OpenTelemetry.
type Core struct {
	provider log.LoggerProvider
	logger   log.Logger
	opts     []log.LoggerOption
	attr     []log.KeyValue
	ctx      context.Context
	err      error
}

// Compile-time check *Core implements zapcore.Core.
var _ zapcore.Core = (*Core)(nil)

// NewCore creates a new [zapcore.Core] that can be used with [go.uber.org/zap.New].
// The name should be the package import path that is being logged.
// The name is ignored for named loggers created using [go.uber.org/zap.Logger.Named].
func NewCore(name string, opts ...Option) *Core { _ = "STUB: not implemented"; return nil }

// Enabled decides whether a given logging level is enabled when logging a message.
func (o *Core) Enabled(level zapcore.Level) bool { _ = "STUB: not implemented"; return false }

// With adds structured context to the Core.
func (o *Core) With(fields []zapcore.Field) zapcore.Core {
	_ = "STUB: not implemented"
	return *new(zapcore.Core)
}

func (o *Core) clone() *Core { _ = "STUB: not implemented"; return nil }

// Sync flushes buffered logs (if any).
func (*Core) Sync() error {
	_ = "STUB: not implemented"

	// Check determines whether the supplied Entry should be logged.
	// If the entry should be logged, the Core adds itself to the CheckedEntry and returns the result.
	return nil
}

func (o *Core) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	_ = "STUB: not implemented"
	return nil
}

// Write method encodes zap fields to OTel logs and emits them.
func (o *Core) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	_ = "STUB: not implemented"
	return nil
}

func convertField(fields []zapcore.Field) (context.Context, []log.KeyValue, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func hasExceptionAttributes(attrs []log.KeyValue) bool { _ = "STUB: not implemented"; return false }

func convertLevel(level zapcore.Level) log.Severity {
	_ = "STUB: not implemented"
	return *new(log.Severity)
}
