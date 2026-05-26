// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelmongo // import "go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/event"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo/internal/semconv"
)

type spanKey struct {
	ConnectionID string
	RequestID    int64
}

type monitor struct {
	sync.Mutex
	spans   map[spanKey]trace.Span
	cfg     config
	semconv semconv.EventMonitor
}

func (m *monitor) Started(ctx context.Context, evt *event.CommandStartedEvent) {
	_ = "STUB: not implemented"
	return
}

func (m *monitor) Succeeded(_ context.Context, evt *event.CommandSucceededEvent) {
	_ = "STUB: not implemented"
	return
}

func (m *monitor) Failed(_ context.Context, evt *event.CommandFailedEvent) {
	_ = "STUB: not implemented"
	return
}

func (m *monitor) Finished(evt *event.CommandFinishedEvent, err error) {
	_ = "STUB: not implemented"
	return
}

// NewMonitor creates a new mongodb event CommandMonitor.
func NewMonitor(opts ...Option) *event.CommandMonitor { _ = "STUB: not implemented"; return nil }

// extractCollection extracts the collection for the given mongodb command event.
// For CRUD operations, this is the first key/value string pair in the bson
// document where key == "<operation>" (e.g. key == "insert").
// For database meta-level operations, such a key may not exist.
func extractCollection(evt *event.CommandStartedEvent) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
