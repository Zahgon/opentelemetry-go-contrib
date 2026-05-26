// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package gcp // import "go.opentelemetry.io/contrib/detectors/gcp"

import (
	"context"

	"go.opentelemetry.io/otel/sdk/resource"
)

// GCE collects resource information of GCE computing instances.
//
// Deprecated: Use gcp.NewDetector() instead, which sets the same resource attributes on GCE.
type GCE struct{}

// compile time assertion that GCE implements the resource.Detector interface.
var _ resource.Detector = (*GCE)(nil)

// Detect detects associated resources when running on GCE hosts.
func (*GCE) Detect(ctx context.Context) (*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// hasProblem checks if the err is not nil or for missing resources.
func hasProblem(err error) bool { _ = "STUB: not implemented"; return false }
