// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package foo is an example sub-package.
package foo // import "go.opentelemetry.io/contrib/examples/namedtracer/foo"

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
)

var lemonsKey = attribute.Key("ex.com/lemons")

// SubOperation is an example to demonstrate the use of named tracer.
// It creates a named tracer with its package path.
func SubOperation(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Using global provider. Alternative is to have application provide a getter
	// for its component to get the instance of the provider.
	return nil
}
