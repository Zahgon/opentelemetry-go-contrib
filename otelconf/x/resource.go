// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package x // import "go.opentelemetry.io/contrib/otelconf/x"

import (
	"go.opentelemetry.io/otel/sdk/resource"
)

func resourceOpts(detectors []ExperimentalResourceDetector) []resource.Option {
	_ = "STUB: not implemented"
	return nil
}

func newResource(r *Resource) (*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
