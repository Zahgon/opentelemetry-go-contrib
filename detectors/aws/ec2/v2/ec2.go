// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package ec2 provides a resource detector for EC2 instances using aws-sdk-go-v2.
package ec2 // import "go.opentelemetry.io/contrib/detectors/aws/ec2/v2"

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
)

var errClient = errors.New("EC2 Client Error")

// client implements methods to capture EC2 environment metadata information.
type client interface {
	GetInstanceIdentityDocument(ctx context.Context, params *imds.GetInstanceIdentityDocumentInput, optFns ...func(*imds.Options)) (*imds.GetInstanceIdentityDocumentOutput, error)
	GetMetadata(ctx context.Context, params *imds.GetMetadataInput, optFns ...func(*imds.Options)) (*imds.GetMetadataOutput, error)
}

// resource detector collects resource information from EC2 environment.
type resourceDetector struct {
	c client
}

// compile time assertion that imds.Client implements client.
var _ client = (*imds.Client)(nil)

// compile time assertion that resourceDetector implements the resource.Detector interface.
var _ resource.Detector = (*resourceDetector)(nil)

// NewResourceDetector returns a resource detector that will detect AWS EC2 resources.
func NewResourceDetector() resource.Detector {
	_ = "STUB: not implemented"
	return *new(resource.Detector)
}

func (detector *resourceDetector) getClient() client {
	_ = "STUB: not implemented"

	// Detect detects associated resources when running in AWS environment.
	return *new(client)
}

func (detector *resourceDetector) Detect(ctx context.Context) (*resource.Resource, error) {
	_ = "STUB: not implemented"
	// Return nil if not able to establish valid client
	return nil, nil
}

// Available method removed in aws-sdk-go-v2, return empty resource if client returns error

func newClient() client { _ = "STUB: not implemented"; return *new(client) }

type metadata struct {
	client     client
	errs       []error
	attributes []attribute.KeyValue
}

func (m *metadata) add(ctx context.Context, k attribute.Key, n string) {
	_ = "STUB: not implemented"
	return
}

func (m *metadata) recordError(path string, err error) { _ = "STUB: not implemented"; return }
