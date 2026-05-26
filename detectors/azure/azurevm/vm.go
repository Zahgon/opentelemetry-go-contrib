// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azurevm // import "go.opentelemetry.io/contrib/detectors/azure/azurevm"

import (
	"context"

	"go.opentelemetry.io/otel/sdk/resource"
)

const defaultAzureVMMetadataEndpoint = "http://169.254.169.254/metadata/instance/compute?api-version=2021-12-13&format=json"

// ResourceDetector collects resource information of Azure VMs.
type ResourceDetector struct {
	endpoint string
}

type vmMetadata struct {
	VMId       *string `json:"vmId"`
	Location   *string `json:"location"`
	ResourceId *string `json:"resourceId"`
	Name       *string `json:"name"`
	VMSize     *string `json:"vmSize"`
	OsType     *string `json:"osType"`
	Version    *string `json:"version"`
}

// New returns a [ResourceDetector] that will detect Azure VM resources.
func New() *ResourceDetector { _ = "STUB: not implemented"; return nil }

// Detect detects associated resources when running on an Azure VM.
func (detector *ResourceDetector) Detect(ctx context.Context) (*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (detector *ResourceDetector) getJSONMetadata(ctx context.Context) ([]byte, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}
