// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package provider contains various providers
// used to replace variables in configuration files.
package provider // import "go.opentelemetry.io/contrib/otelconf/internal/provider"

import (
	"regexp"
)

const validationPattern = `^[a-zA-Z_][a-zA-Z0-9_]*$`

var (
	validationRegexp        = regexp.MustCompile(validationPattern)
	doubleDollarSignsRegexp = regexp.MustCompile(`\$\$([^{$])`)
	envVarRegexp            = regexp.MustCompile(`([$]*)\{([a-zA-Z_][a-zA-Z0-9_]*-?[^}]*)\}`)
)

func ReplaceEnvVars(input []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	// start by replacing all $$ that are not followed by a {
	return nil, nil
}

// check if we have an odd number of $, which indicates that
// env var replacement should be done

// need to expand any default value env var to support the case $${STRING_VALUE:-${STRING_VALUE}}

// expand the default value

func replaceEnvVar(in string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type defaultValue struct {
	data  string
	valid bool
}

func parseEnvVar(in string) (string, defaultValue) {
	_ = "STUB: not implemented"
	return "", *new(defaultValue)
}

func checkRawConfType(val []byte) error { _ = "STUB: not implemented"; return nil }
