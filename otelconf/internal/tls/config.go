// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package tls provides functionality to translate configuration options into tls.Config.
package tls // import "go.opentelemetry.io/contrib/otelconf/internal/tls"

import (
	"crypto/tls"
)

// CreateConfig creates a tls.Config from certificate files.
func CreateConfig(caCertFile, clientCertFile, clientKeyFile *string) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
