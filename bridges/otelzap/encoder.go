// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelzap // import "go.opentelemetry.io/contrib/bridges/otelzap"

import (
	"time"

	"go.opentelemetry.io/otel/log"
	"go.uber.org/zap/zapcore"
)

var (
	_ zapcore.ObjectEncoder = (*objectEncoder)(nil)
	_ zapcore.ArrayEncoder  = (*arrayEncoder)(nil)
)

type namespace struct {
	name  string
	attrs []log.KeyValue
	next  *namespace
}

// objectEncoder implements zapcore.ObjectEncoder.
// It encodes given fields to OTel key-values.
type objectEncoder struct {
	// root is a pointer to the default namespace
	root *namespace
	// cur is a pointer to the namespace we're currently writing to.
	cur *namespace
}

func newObjectEncoder(n int) *objectEncoder { _ = "STUB: not implemented"; return nil }

// It iterates to the end of the linked list and appends namespace data.
// Run this function before accessing complete result.
func (m *objectEncoder) calculate(o *namespace) { _ = "STUB: not implemented"; return }

func (m *objectEncoder) AddArray(key string, v zapcore.ArrayMarshaler) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *objectEncoder) AddObject(k string, v zapcore.ObjectMarshaler) error {
	_ = "STUB: not implemented"
	// Similar to console_encoder which uses capacity of 2:
	// https://github.com/uber-go/zap/blob/bd0cf0447951b77aa98dcfc1ac19e6f58d3ee64f/zapcore/console_encoder.go#L33.
	return nil
}

func (m *objectEncoder) AddBinary(k string, v []byte) { _ = "STUB: not implemented"; return }

func (m *objectEncoder) AddByteString(k string, v []byte) { _ = "STUB: not implemented"; return }

func (m *objectEncoder) AddBool(k string, v bool) { _ = "STUB: not implemented"; return }

func (m *objectEncoder) AddDuration(k string, v time.Duration) { _ = "STUB: not implemented"; return }

func (m *objectEncoder) AddComplex128(k string, v complex128) { _ = "STUB: not implemented"; return }

func (m *objectEncoder) AddFloat64(k string, v float64) { _ = "STUB: not implemented"; return }

func (m *objectEncoder) AddInt64(k string, v int64) { _ = "STUB: not implemented"; return }

func (m *objectEncoder) AddInt(k string, v int) { _ = "STUB: not implemented"; return }

func (m *objectEncoder) AddString(k, v string) { _ = "STUB: not implemented"; return }

func (m *objectEncoder) AddUint64(k string, v uint64) { _ = "STUB: not implemented"; return }

func (m *objectEncoder) AddReflected(k string, v any) error { _ = "STUB: not implemented"; return nil }

// OpenNamespace opens an isolated namespace where all subsequent fields will
// be added.
func (m *objectEncoder) OpenNamespace(k string) { _ = "STUB: not implemented"; return }

func (m *objectEncoder) AddComplex64(k string, v complex64) { _ = "STUB: not implemented"; return }

func (m *objectEncoder) AddTime(k string, v time.Time) { _ = "STUB: not implemented"; return }

func (m *objectEncoder) AddFloat32(k string, v float32) { _ = "STUB: not implemented"; return }

func (m *objectEncoder) AddInt32(k string, v int32) { _ = "STUB: not implemented"; return }

func (m *objectEncoder) AddInt16(k string, v int16) { _ = "STUB: not implemented"; return }

func (m *objectEncoder) AddInt8(k string, v int8) { _ = "STUB: not implemented"; return }

func (m *objectEncoder) AddUint(k string, v uint) { _ = "STUB: not implemented"; return }

func (m *objectEncoder) AddUint32(k string, v uint32) { _ = "STUB: not implemented"; return }

func (m *objectEncoder) AddUint16(k string, v uint16) { _ = "STUB: not implemented"; return }

func (m *objectEncoder) AddUint8(k string, v uint8) { _ = "STUB: not implemented"; return }

func (m *objectEncoder) AddUintptr(k string, v uintptr) { _ = "STUB: not implemented"; return }

func assignUintValue(v uint64) log.Value { _ = "STUB: not implemented"; return *new(log.Value) }

// arrayEncoder implements [zapcore.ArrayEncoder].
type arrayEncoder struct {
	elems []log.Value
}

func newArrayEncoder() *arrayEncoder { _ = "STUB: not implemented"; return nil }

// Similar to console_encoder which uses capacity of 2:
// https://github.com/uber-go/zap/blob/bd0cf0447951b77aa98dcfc1ac19e6f58d3ee64f/zapcore/console_encoder.go#L33.

func (a *arrayEncoder) AppendArray(v zapcore.ArrayMarshaler) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *arrayEncoder) AppendObject(v zapcore.ObjectMarshaler) error {
	_ = "STUB: not implemented"
	// Similar to console_encoder which uses capacity of 2:
	// https://github.com/uber-go/zap/blob/bd0cf0447951b77aa98dcfc1ac19e6f58d3ee64f/zapcore/console_encoder.go#L33.
	return nil
}

func (a *arrayEncoder) AppendReflected(v any) error { _ = "STUB: not implemented"; return nil }

func (a *arrayEncoder) AppendByteString(v []byte) { _ = "STUB: not implemented"; return }

func (a *arrayEncoder) AppendBool(v bool) { _ = "STUB: not implemented"; return }

func (a *arrayEncoder) AppendFloat64(v float64) { _ = "STUB: not implemented"; return }

func (a *arrayEncoder) AppendFloat32(v float32) { _ = "STUB: not implemented"; return }

func (a *arrayEncoder) AppendInt(v int) { _ = "STUB: not implemented"; return }

func (a *arrayEncoder) AppendInt64(v int64) { _ = "STUB: not implemented"; return }

func (a *arrayEncoder) AppendString(v string) { _ = "STUB: not implemented"; return }

func (a *arrayEncoder) AppendComplex128(v complex128) { _ = "STUB: not implemented"; return }

func (a *arrayEncoder) AppendUint64(v uint64) { _ = "STUB: not implemented"; return }

func (a *arrayEncoder) AppendComplex64(v complex64)    { _ = "STUB: not implemented"; return }
func (a *arrayEncoder) AppendDuration(v time.Duration) { _ = "STUB: not implemented"; return }
func (a *arrayEncoder) AppendInt32(v int32)            { _ = "STUB: not implemented"; return }
func (a *arrayEncoder) AppendInt16(v int16)            { _ = "STUB: not implemented"; return }
func (a *arrayEncoder) AppendInt8(v int8)              { _ = "STUB: not implemented"; return }
func (a *arrayEncoder) AppendTime(v time.Time)         { _ = "STUB: not implemented"; return }
func (a *arrayEncoder) AppendUint(v uint)              { _ = "STUB: not implemented"; return }
func (a *arrayEncoder) AppendUint32(v uint32)          { _ = "STUB: not implemented"; return }
func (a *arrayEncoder) AppendUint16(v uint16)          { _ = "STUB: not implemented"; return }
func (a *arrayEncoder) AppendUint8(v uint8)            { _ = "STUB: not implemented"; return }
func (a *arrayEncoder) AppendUintptr(v uintptr)        { _ = "STUB: not implemented"; return }
