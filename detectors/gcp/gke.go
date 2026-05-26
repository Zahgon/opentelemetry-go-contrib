// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package gcp // import "go.opentelemetry.io/contrib/detectors/gcp"

import (
	"context"

	"go.opentelemetry.io/otel/sdk/resource"
)

// GKE collects resource information of GKE computing instances.
//
// Deprecated: Use gcp.NewDetector() instead, which does NOT detect container, pod, and namespace attributes.
// Set those using name using the OTEL_RESOURCE_ATTRIBUTES env var instead.
type GKE struct{}

// compile time assertion that GKE implements the resource.Detector interface.
var _ resource.Detector = (*GKE)(nil)

// Detect detects associated resources when running in GKE environment.
func (*GKE) Detect(ctx context.Context) (*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
