// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package eks provides a resource detector for AWS EKS.
package eks // import "go.opentelemetry.io/contrib/detectors/aws/eks"

import (
	"context"
	"regexp"

	"go.opentelemetry.io/otel/sdk/resource"
	"k8s.io/client-go/kubernetes"
)

const (
	k8sTokenPath      = "/var/run/secrets/kubernetes.io/serviceaccount/token" //nolint:gosec // False positive G101: Potential hardcoded credentials. The detector only check if the token exists.
	k8sCertPath       = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
	authConfigmapNS   = "kube-system"
	authConfigmapName = "aws-auth"
	cwConfigmapNS     = "amazon-cloudwatch"
	cwConfigmapName   = "cluster-info"
	defaultCgroupPath = "/proc/self/cgroup"
	containerIDLength = 64
)

// detectorUtils is used for testing the resourceDetector by abstracting functions that rely on external systems.
type detectorUtils interface {
	fileExists(filename string) bool
	getConfigMap(ctx context.Context, namespace, name string) (map[string]string, error)
	getContainerID() (string, error)
}

// This struct will implement the detectorUtils interface.
type eksDetectorUtils struct {
	clientset *kubernetes.Clientset
}

// resourceDetector for detecting resources running on Amazon EKS.
type resourceDetector struct {
	utils detectorUtils
	err   error
}

// Compile time assertion that resourceDetector implements the resource.Detector interface.
var _ resource.Detector = (*resourceDetector)(nil)

// Compile time assertion that eksDetectorUtils implements the detectorUtils interface.
var _ detectorUtils = (*eksDetectorUtils)(nil)

// is this going to stop working with 1.20 when Docker is deprecated?
var containerIDRegex = regexp.MustCompile(`^.*/docker/(.+)$`)

// NewResourceDetector returns a resource detector that will detect AWS EKS resources.
func NewResourceDetector() resource.Detector {
	_ = "STUB: not implemented"
	return *new(resource.Detector)
}

// Detect returns a Resource describing the Amazon EKS environment being run in.
func (detector *resourceDetector) Detect(ctx context.Context) (*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Return empty resource object if not running in EKS

// Create variable to hold resource attributes

// Get clusterName and append to attributes

// Get containerID and append to attributes

// Return new resource object with clusterName and containerID as attributes

// isEKS checks if the current environment is running in EKS.
func isEKS(ctx context.Context, utils detectorUtils) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Make HTTP GET request

// newK8sDetectorUtils creates the Kubernetes clientset.
func newK8sDetectorUtils() (*eksDetectorUtils, error) {
	_ = "STUB: not implemented"
	// Get cluster configuration
	return nil, nil
}

// Create clientset using generated configuration

// isK8s checks if the current environment is running in a Kubernetes environment.
func isK8s(utils detectorUtils) bool { _ = "STUB: not implemented"; return false }

// fileExists checks if a file with a given filename exists.
func (eksDetectorUtils) fileExists(filename string) bool { _ = "STUB: not implemented"; return false }

// getConfigMap retrieves the configuration map from the k8s API.
func (eksUtils eksDetectorUtils) getConfigMap(ctx context.Context, namespace, name string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getClusterName retrieves the clusterName resource attribute.
func getClusterName(ctx context.Context, utils detectorUtils) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// getContainerID returns the containerID if currently running within a container.
func (eksDetectorUtils) getContainerID() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Retrieve containerID from file
