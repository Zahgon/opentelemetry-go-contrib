// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package envcar // import "go.opentelemetry.io/contrib/propagators/envcar"

// normalize converts s to a valid POSIX environment variable name.
// The conversion rules are:
//   - A–Z, 0–9, and _ are kept as-is.
//   - a–z are uppercased.
//   - All other characters are replaced with _.
//   - If the result would start with a digit, an underscore is prepended.
func normalize(s string) string { _ = "STUB: not implemented"; return "" }

// Pre-allocate the exact output length. If the first byte is a digit,
// the name must be prefixed with '_', so allocate one extra byte.

// Uppercase letters, digits, and underscores are valid as-is.
//nolint:gosec // G115: overflow is not possible.

// Lowercase letters are converted to uppercase.

// All other characters (including non-ASCII runes) become underscores.
