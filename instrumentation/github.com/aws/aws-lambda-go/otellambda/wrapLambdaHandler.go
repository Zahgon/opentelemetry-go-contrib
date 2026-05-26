// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otellambda // import "go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"

import (
	"context"
	"reflect"
)

// wrappedHandlerFunction is a struct which only holds an instrumentor and is
// able to instrument invocations of the user's lambda handler function.
type wrappedHandlerFunction struct {
	instrumentor instrumentor
}

func errorHandler(e error) func(context.Context, any) (any, error) {
	_ = "STUB: not implemented"
	return nil
}

// Ensure handler takes 0-2 values, with context
// as its first value if two arguments exist.
func validateArguments(handler reflect.Type) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Ensure handler returns 0-2 values, with an error
// as its first value if any exist.
func validateReturns(handler reflect.Type) error { _ = "STUB: not implemented"; return nil }

// Wraps and calls customer lambda handler then unpacks response as necessary.
func (whf *wrappedHandlerFunction) wrapperInternals(ctx context.Context, handlerFunc any, eventJSON []byte, event reflect.Value, takesContext bool) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// convert return values into (any, error)

// InstrumentHandler Provides a lambda handler which wraps customer lambda handler with OTel Tracing.
func InstrumentHandler(handlerFunc any, options ...Option) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// note we will always take context to capture lambda context,
// regardless of whether customer takes context

// customer either takes both context and payload or just payload

// lambda SDK normally unmarshalls to customer event type, however
// with the wrapper the SDK unmarshalls to map[string]any
// due to our use of reflection. Therefore we must convert this map
// to customer's desired event, we do so by simply re-marshaling then
// unmarshalling to the desired event type. The remarshalledPayload
// will also be used by users using custom propagators

// Adds OTel span surrounding customer handler call.
func (whf *wrappedHandlerFunction) wrapper(handlerFunc any) func(ctx context.Context, eventJSON []byte, event any, takesContext bool) []reflect.Value {
	_ = "STUB: not implemented"
	return nil
}

// Determine if an any is nil or the
// if the reflect.Value of the event is nil.
func eventExists(event any) bool { _ = "STUB: not implemented"; return false }

// reflect.Value.isNil() can only be called on
// Values of certain Kinds. Unsupported Kinds
// will panic rather than return false
