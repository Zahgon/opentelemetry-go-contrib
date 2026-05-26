// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package gcp // import "go.opentelemetry.io/contrib/detectors/gcp"

import (
	"context"

	"go.opentelemetry.io/otel/sdk/resource"
)

const serviceNamespace = "cloud-run-managed"

// The minimal list of metadata.Client methods we use. Use an interface so we
// can replace it with a fake implementation in the unit test.
type metadataClient interface {
	ProjectID() (string, error)
	Get(string) (string, error)
	InstanceID() (string, error)
}

// CloudRun collects resource information of Cloud Run instance.
//
// Deprecated: Use gcp.NewDetector() instead. Note that it sets faas.* resource attributes instead of service.* attributes.
type CloudRun struct {
	mc     metadataClient
	onGCE  func() bool
	getenv func(string) string
}

// compile time assertion that CloudRun implements the resource.Detector
// interface.
var _ resource.Detector = (*CloudRun)(nil)

// NewCloudRun creates a CloudRun detector.
//
// Deprecated: Use gcp.NewDetector() instead. Note that it sets faas.* resource attributes instead of service.* attributes.
func NewCloudRun() *CloudRun { _ = "STUB: not implemented"; return nil }

func (c *CloudRun) cloudRegion() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Region from the metadata server is in the format /projects/123/regions/r.
// https://cloud.google.com/run/docs/reference/container-contract#metadata-server

// Detect detects associated resources when running on Cloud Run hosts.
// NOTE: the service.namespace attribute is currently hardcoded to be
// "cloud-run-managed". This may change in the future, please do not rely on
// this behavior yet.
func (c *CloudRun) Detect(context.Context) (*resource.Resource, error) {
	_ = "STUB: not implemented"
	// .OnGCE is actually testing whether the metadata server is available.
	// Metadata server is supported on Cloud Run.
	return nil, nil
}

// Part of Cloud Run container runtime contract.
// See https://cloud.google.com/run/docs/reference/container-contract
