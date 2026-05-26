// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package gcp // import "go.opentelemetry.io/contrib/detectors/gcp"

// Version is the current release version of the GCP resource detector.
func Version() string {
	_ = "STUB: not implemented"

	// This string is updated by the pre_release.sh script during release
	return ""
}

// SemVersion is the semantic version to be supplied to tracer/meter creation.
//
// Deprecated: Use [Version] instead.
func SemVersion() string { _ = "STUB: not implemented"; return "" }
