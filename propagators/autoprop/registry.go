// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package autoprop // import "go.opentelemetry.io/contrib/propagators/autoprop"

import (
	"errors"
	"sync"

	"go.opentelemetry.io/otel/propagation"

	"go.opentelemetry.io/contrib/propagators/aws/xray"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/contrib/propagators/jaeger"
	"go.opentelemetry.io/contrib/propagators/ot"
)

// none is the special "propagator" name that means no propagator shall be
// configured.
const none = "none"

// propagators is the registry of TextMapPropagators registered with this
// package. It includes all the OpenTelemetry defaults at startup.
var propagators = &registry{
	names: map[string]propagation.TextMapPropagator{
		// W3C Trace Context.
		"tracecontext": propagation.TraceContext{},
		// W3C Baggage.
		"baggage": propagation.Baggage{},
		// B3 single-header format.
		"b3": b3.New(),
		// B3 multi-header format.
		"b3multi": b3.New(b3.WithInjectEncoding(b3.B3MultipleHeader)),
		// Jaeger.
		"jaeger": jaeger.Jaeger{},
		// AWS X-Ray.
		"xray": xray.Propagator{},
		// OpenTracing Trace.
		"ottrace": ot.OT{},

		// No-op TextMapPropagator.
		none: propagation.NewCompositeTextMapPropagator(),
	},
}

// registry maintains a map of propagator names to TextMapPropagator
// implementations that is safe for concurrent use by multiple goroutines
// without additional locking or coordination.
type registry struct {
	mu    sync.Mutex
	names map[string]propagation.TextMapPropagator
}

// load returns the value stored in the registry index for a key, or nil if no
// value is present. The ok result indicates whether value was found in the
// index.
func (r *registry) load(key string) (p propagation.TextMapPropagator, ok bool) {
	_ = "STUB: not implemented"
	return *new(propagation.TextMapPropagator), false
}

var errDupReg = errors.New("duplicate registration")

// store sets the value for a key if is not already in the registry. errDupReg
// is returned if the registry already contains key.
func (r *registry) store(key string, value propagation.TextMapPropagator) error {
	_ = "STUB: not implemented"
	return nil
}

// drop removes key from the registry if it exists, otherwise nothing.
func (r *registry) drop(key string) { _ = "STUB: not implemented"; return }

// RegisterTextMapPropagator sets the TextMapPropagator p to be used when the
// OTEL_PROPAGATORS environment variable contains the propagator name. This
// will panic if name has already been registered or is a default
// (tracecontext, baggage, b3, b3multi, jaeger, xray, or ottrace).
func RegisterTextMapPropagator(name string, p propagation.TextMapPropagator) {
	_ = "STUB: not implemented"
	return
}

// envRegistry.store will return errDupReg if name is already
// registered. Panic here so the user is made aware of the duplicate
// registration, which could be done by malicious code trying to
// intercept cross-cutting concerns.
//
// Panic for all other errors as well. At this point there should not
// be any other errors returned from the store operation. If there
// are, alert the developer that adding them as soon as possible that
// they need to be handled here.

// TextMapPropagator returns a TextMapPropagator composed from the
// passed names of registered TextMapPropagators. Each name must match an
// already registered TextMapPropagator (see the RegisterTextMapPropagator
// function for more information) or a default (tracecontext, baggage, b3,
// b3multi, jaeger, xray, or ottrace).
//
// If "none" is included in the arguments, or no names are provided, the
// returned TextMapPropagator will be a no-operation implementation.
//
// An error is returned for any un-registered names. The remaining, known,
// names will be used to compose a TextMapPropagator that is returned with the
// error.
func TextMapPropagator(names ...string) (propagation.TextMapPropagator, error) {
	_ = "STUB: not implemented"
	return *new(propagation.TextMapPropagator), nil
}

// If "none" is passed in combination with any other propagator,
// the result still needs to be a no-op propagator. Therefore,
// short-circuit here.

// Do not return a composite of a single propagator.
