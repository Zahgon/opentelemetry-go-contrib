// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package semconv provides semantic convention types and functionality.
package semconv // import "go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo/internal/semconv"

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/event"
	"go.opentelemetry.io/otel/attribute"
)

// Constants for environment variable keys and versions.
const (
	semconvOptIn    = "OTEL_SEMCONV_STABILITY_OPT_IN"
	semconvOptInDup = "database/dup"
)

// EventMonitor is responsible for monitoring events with a specified semantic
// version.
type EventMonitor struct {
	version string
}

// NewEventMonitor creates an EventMonitor with the version set based on the
// OTEL_SEMCONV_STABILITY_OPT_IN environment variable.
func NewEventMonitor() EventMonitor { _ = "STUB: not implemented"; return *new(EventMonitor) }

// AttributeOptions represents options for tracing attributes.
type AttributeOptions struct {
	collectionName           string
	commandAttributeDisabled bool
}

// AttributeOption is a function type that modifies AttributeOptions.
type AttributeOption func(*AttributeOptions)

// WithCollectionName is a functional option to set the collection name in
// AttributeOptions.
func WithCollectionName(collName string) AttributeOption {
	_ = "STUB: not implemented"
	return *new(AttributeOption)
}

// WithCommandAttributeDisabled is a functional option to enable or disable
// command attributes.
func WithCommandAttributeDisabled(disabled bool) AttributeOption {
	_ = "STUB: not implemented"
	return *new(AttributeOption)
}

// hasOptIn returns true if the comma-separated version string contains the
// exact optIn value.
func hasOptIn(version, optIn string) bool { _ = "STUB: not implemented"; return false }

// CommandStartedTraceAttrs generates trace attributes for a CommandStartedEvent
// based on the EventMonitor version.
func (m EventMonitor) CommandStartedTraceAttrs(
	evt *event.CommandStartedEvent,
	opts ...AttributeOption,
) []attribute.KeyValue {
	_ = "STUB: not implemented"
	// Dup implies both v1.26.0 and v1.21.0
	return nil
}

// peerInfo extracts the hostname and port from a CommandStartedEvent.
func peerInfo(evt *event.CommandStartedEvent) (hostname string, port int) {
	_ = "STUB: not implemented"
	return "", 0
}

// Default MongoDB port

// If there's an error (likely because there's no port), assume default port
// and use ConnectionID as hostname

// sanitizeCommand converts a BSON command to a sanitized JSON string.
// TODO: Sanitize values where possible.
// TODO: Limit maximum size.
func sanitizeCommand(command bson.Raw) string { _ = "STUB: not implemented"; return "" }

// commandStartedTraceAttrs generates trace attributes for the latest semantic
// version.
func commandStartedTraceAttrs(evt *event.CommandStartedEvent, setters ...AttributeOption) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

// commandStartedTraceAttrsV1210 generates trace attributes for semantic version
// 1.21.0.
func commandStartedTraceAttrsV1210(evt *event.CommandStartedEvent, setters ...AttributeOption) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}
