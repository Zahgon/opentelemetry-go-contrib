// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package consistent // import "go.opentelemetry.io/contrib/samplers/probability/consistent"

// These are IEEE 754 double-width floating point constants used with
// math.Float64bits.
const (
	offsetExponentMask = 0x7ff0000000000000
	offsetExponentBias = 1023
	significandBits    = 52
)

// expFromFloat64 returns floor(log2(x)).
func expFromFloat64(x float64) int { _ = "STUB: not implemented"; return 0 }

// The biased exponent can only be expressed with 11 bits (size (i.e. 64) -
// significant (i.e 52) - sign (i.e. 1)). Meaning the int conversion below
// is guaranteed to be lossless.

// expToFloat64 returns 2^x.
func expToFloat64(x int) float64 {
	_ = "STUB: not implemented"
	// The exponent field is an 11-bit unsigned integer from 0 to 2047, in
	// biased form: an exponent value of 1023 represents the actual zero.
	// Exponents range from -1022 to +1023 because exponents of -1023 (all 0s)
	// and +1024 (all 1s) are reserved for special numbers.
	// Therefore, the uint64 conversion below is guaranteed to be lossless.
	return 0
}

// splitProb returns the two values of log-adjusted-count nearest to p
// Example:
//
//	splitProb(0.375) => (2, 1, 0.5)
//
// indicates to sample with probability (2^-2) 50% of the time
// and (2^-1) 50% of the time.
func splitProb(p float64) (uint8, uint8, float64) {
	_ = "STUB: not implemented"

	// Note: spec.
	return 0, 0, 0
}

// Take the exponent and drop the significand to locate the
// smaller of two powers of two.

// Low is the smaller of two log-adjusted counts, the negative
// of the exponent computed above.

// High is the greater of two log-adjusted counts (i.e., one
// less than low, a smaller adjusted count means a larger
// probability).

// Return these to probability values and use linear
// interpolation to compute the required probability of
// choosing the low-probability Sampler.

//nolint:gosec // Exponent can only range from -1022 to +1023.
