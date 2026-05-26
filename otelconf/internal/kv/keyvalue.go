// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package kv contains function to translate name value pairs into
// attribute.KeyValue.
package kv // import "go.opentelemetry.io/contrib/otelconf/internal/kv"

import (
	"go.opentelemetry.io/otel/attribute"
)

func FromNameValue(k string, v any) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
