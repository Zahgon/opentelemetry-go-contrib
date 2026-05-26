// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package gcp // import "go.opentelemetry.io/contrib/detectors/gcp"

import (
	"context"

	"github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
)

// NewDetector returns a resource detector which detects resource attributes on:
// * Google Compute Engine (GCE).
// * Google Kubernetes Engine (GKE).
// * Google App Engine (GAE).
// * Cloud Run.
// * Cloud Functions.
func NewDetector() resource.Detector { _ = "STUB: not implemented"; return *new(resource.Detector) }

type detector struct {
	detector gcpDetector
}

// Detect detects associated resources when running on GCE, GKE, GAE,
// Cloud Run, and Cloud functions.
func (d *detector) Detect(context.Context) (*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We don't support this platform yet, so just return with what we have

// resourceBuilder simplifies constructing resources using GCP detection
// library functions.
type resourceBuilder struct {
	errs  []error
	attrs []attribute.KeyValue
}

func (r *resourceBuilder) add(key attribute.Key, detect func() (string, error)) {
	_ = "STUB: not implemented"
	return
}

func (r *resourceBuilder) addInt(key attribute.Key, detect func() (string, error)) {
	_ = "STUB: not implemented"
	return
}

// zoneAndRegion functions are expected to return zone, region, err.
func (r *resourceBuilder) addZoneAndRegion(detect func() (string, string, error)) {
	_ = "STUB: not implemented"
	return
}

func (r *resourceBuilder) addZoneOrRegion(detect func() (string, gcp.LocationType, error)) {
	_ = "STUB: not implemented"
	return
}

func (r *resourceBuilder) build() (*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
