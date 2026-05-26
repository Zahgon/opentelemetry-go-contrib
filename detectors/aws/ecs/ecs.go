// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package ecs provides a resource detector for AWS ECS instances.
package ecs // import "go.opentelemetry.io/contrib/detectors/aws/ecs"

import (
	"context"
	"errors"
	"regexp"

	ecsmetadata "github.com/brunoscheufler/aws-ecs-metadata-go"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
)

const (
	// TypeStr is AWS ECS type.
	TypeStr           = "ecs"
	metadataV3EnvVar  = "ECS_CONTAINER_METADATA_URI"
	metadataV4EnvVar  = "ECS_CONTAINER_METADATA_URI_V4"
	containerIDLength = 64
	defaultCgroupPath = "/proc/self/cgroup"
)

var (
	empty                                 = resource.Empty()
	errCannotReadContainerName            = errors.New("failed to read hostname")
	errCannotParseTaskArn                 = errors.New("cannot parse region and account ID from the Task's ARN: the ARN does not contain at least 6 segments separated by the ':' character")
	errCannotRetrieveLogsGroupMetadataV4  = errors.New("the ECS Metadata v4 did not return a AwsLogGroup name")
	errCannotRetrieveLogsStreamMetadataV4 = errors.New("the ECS Metadata v4 did not return a AwsLogStream name")
	ecsCgroupPathPattern                  = regexp.MustCompile(`/ecs/[^/]+/[a-f0-9]{64}$`)
)

// Create interface for methods needing to be mocked.
type detectorUtils interface {
	getContainerName() (string, error)
	getContainerID() (string, error)
	getContainerMetadataV4(ctx context.Context) (*ecsmetadata.ContainerMetadataV4, error)
	getTaskMetadataV4(ctx context.Context) (*ecsmetadata.TaskMetadataV4, error)
}

// struct implements detectorUtils interface.
type ecsDetectorUtils struct{}

// resource detector collects resource information from Elastic Container Service environment.
type resourceDetector struct {
	utils detectorUtils
}

// compile time assertion that ecsDetectorUtils implements detectorUtils interface.
var _ detectorUtils = (*ecsDetectorUtils)(nil)

// compile time assertion that resource detector implements the resource.Detector interface.
var _ resource.Detector = (*resourceDetector)(nil)

// NewResourceDetector returns a resource detector that will detect AWS ECS resources.
func NewResourceDetector() resource.Detector {
	_ = "STUB: not implemented"
	return *new(resource.Detector)
}

// Detect finds associated resources when running on ECS environment.
func (detector *resourceDetector) Detect(ctx context.Context) (*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A valid ARN should have at least 6 parts.

func (*resourceDetector) getBaseArn(arns ...string) string { _ = "STUB: not implemented"; return "" }

func (*resourceDetector) getLogsAttributes(metadata *ecsmetadata.ContainerMetadataV4) ([]attribute.KeyValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html

// a valid arn should have at least 6 parts

// returns metadata v4 for the container.
func (ecsDetectorUtils) getContainerMetadataV4(ctx context.Context) (*ecsmetadata.ContainerMetadataV4, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// returns metadata v4 for the task.
func (ecsDetectorUtils) getTaskMetadataV4(ctx context.Context) (*ecsmetadata.TaskMetadataV4, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// returns docker container ID from default c group path.
func (ecsDetectorUtils) getContainerID() (string, error) {
	_ = "STUB: not implemented"
	return "",

		// Cgroups are used only under Linux.
		nil
}

// Cgroups file not found.
// For example, windows; or when running integration tests outside of a container.

// returns host name reported by the kernel.
func (ecsDetectorUtils) getContainerName() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getCgroupContainerID(fileData []byte) string { _ = "STUB: not implemented"; return "" }
