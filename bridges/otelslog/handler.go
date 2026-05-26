// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package otelslog provides [Handler], an [slog.Handler] implementation, that
// can be used to bridge between the [log/slog] API and [OpenTelemetry].
//
// # Record Conversion
//
// The [slog.Record] are converted to OpenTelemetry [log.Record] in the following
// way:
//
//   - Time is set as the Timestamp.
//   - Message is set as the Body using a [log.StringValue].
//   - Level is transformed and set as the Severity. The SeverityText is also
//     set.
//   - PC is dropped.
//   - Attr are transformed and set as the Attributes.
//
// The Level is transformed by using the static offset to the OpenTelemetry
// Severity types. For example:
//
//   - [slog.LevelDebug] is transformed to [log.SeverityDebug]
//   - [slog.LevelInfo] is transformed to [log.SeverityInfo]
//   - [slog.LevelWarn] is transformed to [log.SeverityWarn]
//   - [slog.LevelError] is transformed to [log.SeverityError]
//
// Attribute values are transformed based on their [slog.Kind]:
//
//   - [slog.KindAny] values are transformed based on their type or
//     into a string value encoded using [fmt.Sprintf] if there is no matching type.
//   - [slog.KindBool] are transformed to [log.BoolValue] directly.
//   - [slog.KindDuration] are transformed to [log.Int64Value] as nanoseconds.
//   - [slog.KindFloat64] are transformed to [log.Float64Value] directly.
//   - [slog.KindInt64] are transformed to [log.Int64Value] directly.
//   - [slog.KindString] are transformed to [log.StringValue] directly.
//   - [slog.KindTime] are transformed to [log.Int64Value] as nanoseconds since
//     the Unix epoch.
//   - [slog.KindUint64] are transformed to [log.Int64Value] using int64
//     conversion.
//   - [slog.KindGroup] are transformed to [log.MapValue] using appropriate
//     transforms for each group value.
//   - [slog.KindLogValuer] the value is resolved and then transformed.
//
// [OpenTelemetry]: https://opentelemetry.io/docs/concepts/signals/logs/
package otelslog // import "go.opentelemetry.io/contrib/bridges/otelslog"

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/log"
)

// NewLogger returns a new [slog.Logger] backed by a new [Handler]. See
// [NewHandler] for details on how the backing Handler is created.
func NewLogger(name string, options ...Option) *slog.Logger { _ = "STUB: not implemented"; return nil }

type config struct {
	provider   log.LoggerProvider
	version    string
	schemaURL  string
	attributes []attribute.KeyValue
	source     bool
}

func newConfig(options []Option) config { _ = "STUB: not implemented"; return *new(config) }

func (c config) logger(name string) log.Logger { _ = "STUB: not implemented"; return *new(log.Logger) }

// Option configures a [Handler].
type Option interface {
	apply(config) config
}

type optFunc func(config) config

func (f optFunc) apply(c config) config {
	_ = "STUB: not implemented"

	// WithVersion returns an [Option] that configures the version of the
	// [log.Logger] used by a [Handler]. The version should be the version of the
	// package that is being logged.
	return *new(config)
}

func WithVersion(version string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSchemaURL returns an [Option] that configures the semantic convention
// schema URL of the [log.Logger] used by a [Handler]. The schemaURL should be
// the schema URL for the semantic conventions used in log records.
func WithSchemaURL(schemaURL string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAttributes returns an [Option] that configures the instrumentation scope
// attributes of the [log.Logger] used by a [Handler].
func WithAttributes(attributes ...attribute.KeyValue) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithLoggerProvider returns an [Option] that configures [log.LoggerProvider]
// used by a [Handler] to create its [log.Logger].
//
// By default if this Option is not provided, the Handler will use the global
// LoggerProvider.
func WithLoggerProvider(provider log.LoggerProvider) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSource returns an [Option] that configures the [Handler] to include
// the source location of the log record in log attributes.
func WithSource(source bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// Handler is an [slog.Handler] that sends all logging records it receives to
// OpenTelemetry. See package documentation for how conversions are made.
type Handler struct {
	// Ensure forward compatibility by explicitly making this not comparable.
	noCmp [0]func() //nolint:unused  // This is indeed used.

	attrs  *kvBuffer
	group  *group
	logger log.Logger

	source bool
}

// Compile-time check *Handler implements slog.Handler.
var _ slog.Handler = (*Handler)(nil)

// NewHandler returns a new [Handler] to be used as an [slog.Handler].
//
// If [WithLoggerProvider] is not provided, the returned Handler will use the
// global LoggerProvider.
//
// The provided name needs to uniquely identify the code being logged. This is
// most commonly the package name of the code. If name is empty, the
// [log.Logger] implementation may override this value with a default.
func NewHandler(name string, options ...Option) *Handler { _ = "STUB: not implemented"; return nil }

// Handle handles the passed record.
func (h *Handler) Handle(ctx context.Context, record slog.Record) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *Handler) convertRecord(r slog.Record) log.Record {
	_ = "STUB: not implemented"
	return *new(log.Record)
}

// A Handler should not output groups if there are no attributes.

// A Handler should not output groups if there are no attributes.

// Enabled returns true if the Handler is enabled to log for the provided
// context and Level. Otherwise, false is returned if it is not enabled.
func (h *Handler) Enabled(ctx context.Context, l slog.Level) bool {
	_ = "STUB: not implemented"
	return false
}

// WithAttrs returns a new [slog.Handler] based on h that will log using the
// passed attrs.
func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

// WithGroup returns a new [slog.Handler] based on h that will log all messages
// and attributes within a group of the provided name.
func (h *Handler) WithGroup(name string) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

// group represents a group received from slog.
type group struct {
	// name is the name of the group.
	name string
	// attrs are the attributes associated with the group.
	attrs *kvBuffer
	// next points to the next group that holds this group.
	//
	// Groups are represented as map value types in OpenTelemetry. This means
	// that for an slog group hierarchy like the following ...
	//
	//   WithGroup("G").WithGroup("H").WithGroup("I")
	//
	// the corresponding OpenTelemetry log value types will have the following
	// hierarchy ...
	//
	//   KeyValue{
	//     Key: "G",
	//     Value: []KeyValue{{
	//       Key: "H",
	//       Value: []KeyValue{{
	//         Key: "I",
	//         Value: []KeyValue{},
	//       }},
	//     }},
	//   }
	//
	// When attributes are recorded (i.e. Info("msg", "key", "value") or
	// WithAttrs("key", "value")) they need to be added to the "leaf" group. In
	// the above example, that would be group "I":
	//
	//   KeyValue{
	//     Key: "G",
	//     Value: []KeyValue{{
	//       Key: "H",
	//       Value: []KeyValue{{
	//         Key: "I",
	//         Value: []KeyValue{
	//           String("key", "value"),
	//         },
	//       }},
	//     }},
	//   }
	//
	// Therefore, groups are structured as a linked-list with the "leaf" node
	// being the head of the list. Following the above example, the group data
	// representation would be ...
	//
	//   *group{"I", next: *group{"H", next: *group{"G"}}}
	next *group
}

// Err returns the error to use from g's linked-list (including g itself). If
// no error is found, nil is returned.
func (g *group) Err() error { _ = "STUB: not implemented"; return nil }

// NextNonEmpty returns the next group within g's linked-list that has
// attributes (including g itself). If no group is found, nil is returned.
func (g *group) NextNonEmpty() *group { _ = "STUB: not implemented"; return nil }

// KeyValue returns group g containing kvs as a [log.KeyValue]. The value of
// the returned KeyValue will be of type [log.KindMap].
//
// The passed kvs are rendered in the returned value, but are not added to the
// group.
//
// This does not check g. It is the callers responsibility to ensure g is
// non-empty or kvs is non-empty so as to return a valid group representation
// (according to slog).
func (g *group) KeyValue(kvs ...log.KeyValue) log.KeyValue {
	_ = "STUB: not implemented"
	// Assumes checking of group g already performed (i.e. non-empty).
	return *new(log.KeyValue)
}

// A Handler should not output groups if there are no attributes.

// Clone returns a copy of g.
func (g *group) Clone() *group { _ = "STUB: not implemented"; return nil }

// AddAttrs add attrs to g.
func (g *group) AddAttrs(attrs []slog.Attr) { _ = "STUB: not implemented"; return }

type kvBuffer struct {
	data []log.KeyValue
	err  error
}

func newKVBuffer(n int) *kvBuffer { _ = "STUB: not implemented"; return nil }

// Len returns the number of [log.KeyValue] held by b.
func (b *kvBuffer) Len() int { _ = "STUB: not implemented"; return 0 }

// Clone returns a copy of b.
func (b *kvBuffer) Clone() *kvBuffer { _ = "STUB: not implemented"; return nil }

// KeyValues returns kvs appended to the [log.KeyValue] held by b.
func (b *kvBuffer) KeyValues(kvs ...log.KeyValue) []log.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

// AddAttrs adds attrs to b.
func (b *kvBuffer) AddAttrs(attrs []slog.Attr) { _ = "STUB: not implemented"; return }

// AddAttr adds attr to b and returns true.
//
// This is designed to be passed to the AddAttributes method of an
// [slog.Record].
//
// If attr is a group with an empty key, its values will be flattened.
//
// If attr is empty, it will be dropped.
func (b *kvBuffer) AddAttr(attr slog.Attr) bool { _ = "STUB: not implemented"; return false }

// A Handler should inline the Attrs of a group with an empty key.

// A Handler should ignore an empty Attr.

func convert(v slog.Value) log.Value { _ = "STUB: not implemented"; return *new(log.Value) }

// Try to handle this as gracefully as possible.
//
// Don't panic here. The goal here is to have developers find this
// first if a new slog.Kind is added. A test on the new kind will find
// this malformed attribute as well as a panic. However, it is
// preferable to have user's open issue asking why their attributes
// have a "unhandled: " prefix than say that their code is panicking.
