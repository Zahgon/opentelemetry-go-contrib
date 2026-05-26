// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package testtls provides runtime-generated TLS materials for tests.
package testtls // import "go.opentelemetry.io/contrib/otelconf/internal/testtls"

import (
	"crypto/rsa"
	"crypto/x509"
)

// Material contains generated CA, server, and client cert file paths.
type Material struct {
	CACertPath     string
	CAKeyPath      string
	ServerCertPath string
	ServerKeyPath  string
	ClientCertPath string
	ClientKeyPath  string
	BadCertPath    string
}

// TB is the minimal testing surface needed by this helper.
type TB interface {
	Helper()
	TempDir() string
	Fatalf(format string, args ...any)
}

// Write generates mTLS assets under t.TempDir so tests do not depend on expiring fixtures.
func Write(t TB) Material { _ = "STUB: not implemented"; return *new(Material) }

func mustRSAKey(t TB) *rsa.PrivateKey { _ = "STUB: not implemented"; return nil }

func mustCertificate(t TB, template, parent *x509.Certificate, publicKey *rsa.PublicKey, signer *rsa.PrivateKey) []byte {
	_ = "STUB: not implemented"
	return nil
}

func writeCert(t TB, path string, der []byte) { _ = "STUB: not implemented"; return }

func writeKey(t TB, path string, key *rsa.PrivateKey) { _ = "STUB: not implemented"; return }

func writeBadCert(t TB, path string) { _ = "STUB: not implemented"; return }
