// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package minsev // import "go.opentelemetry.io/contrib/processors/minsev"

import (
	"encoding"
	"encoding/json"
	"fmt"
	"sync/atomic"

	"go.opentelemetry.io/otel/log"
)

// Severity represents a log record severity (also known as log level). Smaller
// numerical values correspond to less severe log records (such as debug
// events), larger numerical values correspond to more severe log records (such
// as errors and critical events).
type Severity int

var (
	// Ensure Severity implements fmt.Stringer.
	_ fmt.Stringer = Severity(0)
	// Ensure Severity implements json.Marshaler.
	_ json.Marshaler = Severity(0)
	// Ensure Severity implements json.Unmarshaler.
	_ json.Unmarshaler = (*Severity)(nil)
	// Ensure Severity implements encoding.TextMarshaler.
	_ encoding.TextMarshaler = Severity(0)
	// Ensure Severity implements encoding.TextUnmarshaler.
	_ encoding.TextUnmarshaler = (*Severity)(nil)
)

// Severity values defined by OpenTelemetry.
const (
	// A fine-grained debugging log record. Typically disabled in default
	// configurations.
	SeverityTrace1 Severity = -8 // TRACE
	SeverityTrace2 Severity = -7 // TRACE2
	SeverityTrace3 Severity = -6 // TRACE3
	SeverityTrace4 Severity = -5 // TRACE4

	// A debugging log record.
	SeverityDebug1 Severity = -4 // DEBUG
	SeverityDebug2 Severity = -3 // DEBUG2
	SeverityDebug3 Severity = -2 // DEBUG3
	SeverityDebug4 Severity = -1 // DEBUG4

	// An informational log record. Indicates that an event happened.
	SeverityInfo1 Severity = 0 // INFO
	SeverityInfo2 Severity = 1 // INFO2
	SeverityInfo3 Severity = 2 // INFO3
	SeverityInfo4 Severity = 3 // INFO4

	// A warning log record. Not an error but is likely more important than an
	// informational event.
	SeverityWarn1 Severity = 4 // WARN
	SeverityWarn2 Severity = 5 // WARN2
	SeverityWarn3 Severity = 6 // WARN3
	SeverityWarn4 Severity = 7 // WARN4

	// An error log record. Something went wrong.
	SeverityError1 Severity = 8  // ERROR
	SeverityError2 Severity = 9  // ERROR2
	SeverityError3 Severity = 10 // ERROR3
	SeverityError4 Severity = 11 // ERROR4

	// A fatal log record such as application or system crash.
	SeverityFatal1 Severity = 12 // FATAL
	SeverityFatal2 Severity = 13 // FATAL2
	SeverityFatal3 Severity = 14 // FATAL3
	SeverityFatal4 Severity = 15 // FATAL4

	// Convenience definitions for the base severity of each level.
	SeverityTrace = SeverityTrace1
	SeverityDebug = SeverityDebug1
	SeverityInfo  = SeverityInfo1
	SeverityWarn  = SeverityWarn1
	SeverityError = SeverityError1
	SeverityFatal = SeverityFatal1
)

// Severity returns the receiver translated to a [log.Severity].
//
// It implements [Severitier].
func (s Severity) Severity() log.Severity {
	_ = "STUB: not implemented"
	// Clamp to the defined range of log.Severity values. This provides a
	// closer approximation for out-of-range values instead of returning
	// log.SeverityUndefined.
	return *new(log.Severity)
}

// The relative ordering and contiguous definition of both sets of
// severities allows a constant offset translation instead of a lookup
// table. Keep this in sync if either definition changes.

// String returns a name for the severity level. If the severity level has a
// name, then that name in uppercase is returned. If the severity level is
// outside named values, then an signed integer is appended to the uppercased
// name.
//
// Examples:
//
//	SeverityWarn1.String() => "WARN"
//	(SeverityInfo1+2).String() => "INFO3"
//	(SeverityFatal4+2).String() => "FATAL+6"
//	(SeverityTrace1-3).String() => "TRACE-3"
func (s Severity) String() string { _ = "STUB: not implemented"; return "" }

// No sign for known fine-grained severity values.

// Exclude zero from positive scale count.

// MarshalJSON implements [encoding/json.Marshaler] by quoting the output of
// [Severity.String].
func (s Severity) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	// AppendQuote is sufficient for JSON-encoding all Severity strings. They
	// don't contain any runes that would produce invalid JSON when escaped.
	return nil, nil
}

// UnmarshalJSON implements [encoding/json.Unmarshaler] It accepts any string
// produced by [Severity.MarshalJSON], ignoring case. It also accepts numeric
// offsets that would result in a different string on output. For example,
// "ERROR-8" will unmarshal as [SeverityInfo].
func (s *Severity) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// AppendText implements [encoding.TextAppender] by calling [Severity.String].
func (s Severity) AppendText(b []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalText implements [encoding.TextMarshaler] by calling
// [Severity.AppendText].
func (s Severity) MarshalText() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// UnmarshalText implements [encoding.TextUnmarshaler]. It accepts any string
		// produced by [Severity.MarshalText], ignoring case. It also accepts numeric
		// offsets that would result in a different string on output. For example,
		// "ERROR-8" will marshal as [SeverityInfo].
		nil
}

func (s *Severity) UnmarshalText(data []byte) error { _ = "STUB: not implemented"; return nil }

// parse parses str into s.
//
// It will return an error if str is not a valid severity string.
//
// The string is expected to be in the format of "NAME[N][+/-OFFSET]", where
// NAME is one of the severity names ("TRACE", "DEBUG", "INFO", "WARN",
// "ERROR", "FATAL"), OFFSET is an optional signed integer offset, and N is an
// optional fine-grained severity level that modifies the base severity name.
//
// Name is parsed in a case-insensitive way. Meaning, "info", "Info",
// "iNfO", etc. are all equivalent to "INFO".
//
// Fine-grained severity levels are expected to be in the range of 1 to 4,
// where 1 is the base severity level, and 2, 3, and 4 are more fine-grained
// levels. However, fine-grained levels greater than 4 are also accepted, and
// they will be treated as an 1-based offset from the base severity level.
//
// For example, "ERROR3" will be parsed as "ERROR" with a fine-grained level of
// 3, which corresponds to [SeverityError3], "FATAL+2" will be parsed as
// "FATAL" with an offset of +2, which corresponds to [SeverityFatal2], and
// "INFO2+1" is parsed as INFO with a fine-grained level of 2 and an offset of
// +1, which corresponds to [SeverityInfo3].
//
// Fine-grained severity levels are based on counting numbers excluding zero.
// If a fine-grained level of 0 is provided it is treaded as equivalent to the
// base severity level.  For example, "INFO0" is equivalent to [SeverityInfo1].
func (s *Severity) parse(str string) (err error) {
	_ = "STUB: not implemented"

	// Handle empty str as a special case and parse it as the default
	// SeverityInfo1.
	//
	// Do not parse this below in the switch statement of the name. That
	// will allow strings like "2", "-1", "2+1", "+3", etc. to be accepted
	// and that adds ambiguity. For example, a user may expect that "2" is
	// parsed as SeverityInfo2 based on an implied "SeverityInfo1" prefix,
	// but they may also expect it be parsed as SeverityInfo3 which has a
	// numeric value of 2. Avoid this ambiguity by treating those inputs
	// as invalid, and only accept the empty string as a special case.
	return nil
}

// Default severity.

// Parse +/- offset suffix, if present.

// Parse fine-grained severity level suffix, if present.
// This supports formats like "ERROR3", "FATAL4", etc.

// Convert 1-based to 0-based.

// A SeverityVar is a [Severity] variable, to allow a [LogProcessor] severity
// to change dynamically. It implements [Severitier] as well as a Set method,
// and it is safe for use by multiple goroutines.
//
// The zero SeverityVar corresponds to [SeverityInfo].
type SeverityVar struct {
	val atomic.Int64
}

var (
	// Ensure Severity implements fmt.Stringer.
	_ fmt.Stringer = (*SeverityVar)(nil)
	// Ensure Severity implements encoding.TextMarshaler.
	_ encoding.TextMarshaler = (*SeverityVar)(nil)
	// Ensure Severity implements encoding.TextUnmarshaler.
	_ encoding.TextUnmarshaler = (*SeverityVar)(nil)
)

// Severity returns v's severity.
func (v *SeverityVar) Severity() log.Severity { _ = "STUB: not implemented"; return *new(log.Severity) }

// Set sets v's Severity to l.
func (v *SeverityVar) Set(l Severity) { _ = "STUB: not implemented"; return }

// String returns a string representation of the SeverityVar.
func (v *SeverityVar) String() string { _ = "STUB: not implemented"; return "" }

// AppendText implements [encoding.TextAppender]
// by calling [Severity.AppendText].
func (v *SeverityVar) AppendText(b []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalText implements [encoding.TextMarshaler]
// by calling [SeverityVar.AppendText].
func (v *SeverityVar) MarshalText() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// UnmarshalText implements [encoding.TextUnmarshaler]
		// by calling [Severity.UnmarshalText].
		nil
}

func (v *SeverityVar) UnmarshalText(data []byte) error { _ = "STUB: not implemented"; return nil }

// A Severitier provides a [log.Severity] value.
type Severitier interface {
	Severity() log.Severity
}
