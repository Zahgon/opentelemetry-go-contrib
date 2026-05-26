// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package otellogrus provides a [Hook], a [logrus.Hook] implementation that
// can be used to bridge between the [github.com/sirupsen/logrus] API and
// [OpenTelemetry].
//
// # Record Conversion
//
// The [logrus.Entry] records are converted to OpenTelemetry [log.Record] in
// the following way:
//
//   - Time is set as the Timestamp.
//   - Message is set as the Body using a [log.StringValue].
//   - Level is transformed and set as the Severity. The SeverityText is also set.
//   - Fields are transformed and set as the attributes.
//   - A field with key [logrus.ErrorKey] and an [error] value is set using
//     [log.Record.SetErr].
//
// The Level is transformed to the OpenTelemetry
// Severity types. For example:
//
//   - [logrus.DebugLevel] is transformed to [log.SeverityDebug]
//   - [logrus.InfoLevel] is transformed to [log.SeverityInfo]
//   - [logrus.WarnLevel] is transformed to [log.SeverityWarn]
//   - [logrus.ErrorLevel] is transformed to [log.SeverityError]
//   - [logrus.FatalLevel] is transformed to [log.SeverityFatal]
//   - [logrus.PanicLevel] is transformed to [log.SeverityFatal4]
//
// Field values are transformed based on their type into log attributes, or
// into a string value encoded using [fmt.Sprintf] if there is no matching type.
//
// [OpenTelemetry]: https://opentelemetry.io/docs/concepts/signals/logs/
package otellogrus // import "go.opentelemetry.io/contrib/bridges/otellogrus"

import (
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/log"
)

type config struct {
	provider   log.LoggerProvider
	version    string
	schemaURL  string
	attributes []attribute.KeyValue

	levels []logrus.Level
}

func newConfig(options []Option) config { _ = "STUB: not implemented"; return *new(config) }

func (c config) logger(name string) log.Logger { _ = "STUB: not implemented"; return *new(log.Logger) }

// Option configures a [Hook].
type Option interface {
	apply(config) config
}

type optFunc func(config) config

func (f optFunc) apply(c config) config {
	_ = "STUB: not implemented"

	// WithVersion returns an [Option] that configures the version of the
	// [log.Logger] used by a [Hook]. The version should be the version of the
	// package that is being logged.
	return *new(config)
}

func WithVersion(version string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSchemaURL returns an [Option] that configures the semantic convention
// schema URL of the [log.Logger] used by a [Hook]. The schemaURL should be
// the schema URL for the semantic conventions used in log records.
func WithSchemaURL(schemaURL string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAttributes returns an [Option] that configures the instrumentation scope
// attributes of the [log.Logger] used by a [Hook].
func WithAttributes(attributes ...attribute.KeyValue) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithLoggerProvider returns an [Option] that configures [log.LoggerProvider]
// used by a [Hook].
//
// By default if this Option is not provided, the Hook will use the global
// LoggerProvider.
func WithLoggerProvider(provider log.LoggerProvider) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithLevels returns an [Option] that configures the log levels that will fire
// the configured [Hook].
//
// By default if this Option is not provided, the Hook will fire for all levels.
// LoggerProvider.
func WithLevels(l []logrus.Level) Option { _ = "STUB: not implemented"; return *new(Option) }

// NewHook returns a new [Hook] to be used as a [logrus.Hook].
//
// If [WithLoggerProvider] is not provided, the returned Hook will use the
// global LoggerProvider.
func NewHook(name string, options ...Option) *Hook { _ = "STUB: not implemented"; return nil }

// Hook is a [logrus.Hook] that sends all logging records it receives to
// OpenTelemetry. See package documentation for how conversions are made.
type Hook struct {
	logger log.Logger
	levels []logrus.Level
}

// Levels returns the list of log levels we want to be sent to OpenTelemetry.
func (h *Hook) Levels() []logrus.Level {
	_ = "STUB: not implemented"

	// Fire handles the passed record, and sends it to OpenTelemetry.
	return nil
}

func (h *Hook) Fire(entry *logrus.Entry) error { _ = "STUB: not implemented"; return nil }

func (*Hook) convertEntry(e *logrus.Entry) log.Record {
	_ = "STUB: not implemented"
	return *new(log.Record)
}

func convertFields(fields logrus.Fields) ([]log.KeyValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertSeverity(level logrus.Level) log.Severity {
	_ = "STUB: not implemented"
	return *new(log.Severity)
}

// PanicLevel is not supported by OpenTelemetry, use Fatal4 as the highest severity.

// If the level is not recognized, use SeverityUndefined as the lowest severity.
// we should never reach this point as logrus only uses the above levels.
