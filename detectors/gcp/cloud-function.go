// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package gcp provides a resource detector for GCP Cloud Function.
package gcp // import "go.opentelemetry.io/contrib/detectors/gcp"

import (
	"context"

	"go.opentelemetry.io/otel/sdk/resource"
)

const (
	gcpFunctionNameKey = "K_SERVICE"
)

// NewCloudFunction will return a GCP Cloud Function resource detector.
//
// Deprecated: Use gcp.NewDetector() instead, which sets the same resource attributes.
func NewCloudFunction() resource.Detector {
	_ = "STUB: not implemented"
	return *new(resource.Detector)
}

// cloudFunction collects resource information of GCP Cloud Function.
type cloudFunction struct {
	cloudRun *CloudRun
}

// Detect detects associated resources when running in GCP Cloud Function.
func (f *cloudFunction) Detect(context.Context) (*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*cloudFunction) googleCloudFunctionName() (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
